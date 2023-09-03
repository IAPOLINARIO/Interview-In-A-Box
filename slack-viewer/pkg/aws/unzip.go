package aws

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cheggaaa/pb"
	"github.com/slack-viewer/pkg/config"
)

// sess, err := GetAWSSession("us-east-2")
func createFolderIfNotExists(svc *s3.S3, folder string) error {
	cfg, err := config.GetConfig()

	if err != nil {
		return fmt.Errorf("Missing configuration parameters")
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(cfg.Server.Bucket),
		Key:    aws.String(strings.TrimSuffix(folder, "/") + "/"),
		Body:   bytes.NewReader([]byte{}),
	})
	return err
}

func UnzipS3File(fileName, outputFolder string, numWorkers int) error {
	cfg, err := config.GetConfig()

	if err != nil {
		return fmt.Errorf("Missing configuration parameters")
	}

	sess, err := GetAWSSession()

	if err != nil {
		return err
	}

	svc := s3.New(sess)

	input := &s3.GetObjectInput{
		Bucket: aws.String(cfg.Server.Bucket),
		Key:    aws.String(fileName),
	}

	obj, err := svc.GetObject(input)

	if err != nil {
		return err
	}
	defer obj.Body.Close()

	err = createFolderIfNotExists(svc, outputFolder)
	if err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp("", "s3unzip-*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, obj.Body)
	if err != nil {
		return err
	}

	zipReader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		return err
	}
	defer zipReader.Close()

	progressBar := pb.StartNew(len(zipReader.File))

	fileChan := make(chan *zip.File, len(zipReader.File))
	for _, f := range zipReader.File {
		fileChan <- f
	}
	close(fileChan)

	errChan := make(chan error, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range fileChan {
				progressBar.Increment()

				if f.FileInfo().IsDir() {
					continue
				}

				rc, err := f.Open()
				if err != nil {
					errChan <- err
					return
				}

				buf := new(bytes.Buffer)
				_, err = io.Copy(buf, rc)
				rc.Close()
				if err != nil {
					errChan <- err
					return
				}

				filePath := outputFolder + "/" + f.Name
				_, err = svc.PutObject(&s3.PutObjectInput{
					Bucket: aws.String(cfg.Server.Bucket),
					Key:    aws.String(strings.TrimPrefix(filePath, "/")),
					Body:   bytes.NewReader(buf.Bytes()),
				})
				if err != nil {
					errChan <- err
					return
				}
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for e := range errChan {
		if e != nil {
			return e
		}
	}

	progressBar.Finish()

	return nil
}

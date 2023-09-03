package aws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cheggaaa/pb/v3"
	"github.com/slack-viewer/pkg/config"
	"github.com/slack-viewer/pkg/dtos"
	"github.com/slack-viewer/pkg/parser"
	"github.com/slack-viewer/pkg/slack"
)

type DataType int

const (
	Channel = iota
	Group
	Mpim
	Dm
)

func (dt DataType) String() string {
	switch dt {
	case Channel:
		return "Channels"
	case Group:
		return "Groups"
	case Mpim:
		return "Mpims"
	case Dm:
		return "Dms"
	default:
		return "Unknown"
	}
}
func SaveSlackHistoryIntoS3BucketAsJSON(data interface{}, destinationDir string) error {
	cfg, err := config.GetConfig()

	if err != nil {
		return fmt.Errorf("Missing configuration parameters")
	}

	// Initialize an AWS session
	sess, err := GetAWSSession()

	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}

	svc := s3.New(sess)
	input := &s3.PutObjectInput{
		Bucket: aws.String(cfg.Server.Bucket),
		Key:    aws.String(destinationDir),
		Body:   bytes.NewReader(jsonData),
	}

	_, err = svc.PutObject(input)
	if err != nil {
		return fmt.Errorf("failed to write JSON to S3: %w", err)
	}

	return nil
}

func CheckIfFileExists(sourceDir string, fileName string) (bool, error) {
	// Initialize an AWS session
	sess, err := GetAWSSession()

	if err != nil {
		return false, err
	}

	// Create an S3 client
	svc := s3.New(sess)

	var objects []*s3.Object

	fullPath := fmt.Sprintf("%s/%s", sourceDir, fileName)
	fmt.Printf("Checking S3 folder: %s\n", fmt.Sprintf("%s/%s", sourceDir, fileName))
	err = GetObjects(svc, sourceDir, fullPath, &objects)
	if err != nil {
		return false, err
	}

	if len(objects) > 0 {
		return true, nil
	}

	return false, err

}

func ReadS3JSONFile(FullFilePath string) ([]byte, error) {

	cfg, err := config.GetConfig()

	if err != nil {
		return nil, fmt.Errorf("Missing configuration parameters")
	}

	// Initialize an AWS session
	sess, err := GetAWSSession()

	if err != nil {
		return nil, err
	}

	// Create an S3 client
	svc := s3.New(sess)

	// Create an S3 input object for the file
	input := &s3.GetObjectInput{
		Bucket: aws.String(cfg.Server.Bucket),
		Key:    aws.String(FullFilePath),
	}

	// Get the file from S3
	output, err := svc.GetObject(input)
	if err != nil {
		return nil, err
	}

	// Read the contents of the file
	contents, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func FindAndParseSlackMessagesInS3Files(folders []string, mainFolder string, progressBar *pb.ProgressBar) (result []dtos.SlackHistoryGroup, err error) {
	for _, folder := range folders {
		groups, err := ParseS3FilesIntoSlackMessagesGroup(mainFolder, folder)

		if err != nil {
			return result, err
		}

		historyGroup := dtos.SlackHistoryGroup{
			GroupName: folder,
			Groups:    groups,
		}

		result = append(result, historyGroup)

		// Increment the progress bar
		progressBar.Increment()
	}

	return result, err
}

func ParseS3FilesIntoSlackMessagesGroup(mainFolder string, subFolderName string) (result []dtos.SlackMessagesGroup, err error) {

	fullPath := fmt.Sprintf("%s/%s", mainFolder, subFolderName)
	files, err := ListFilesInS3Bucket(fullPath)
	messageCount := 0

	if err != nil {
		return nil, err
	}

	for _, file := range files {

		//fmt.Println("Reading file:", file)
		fileContent, err := ReadS3JSONFile(file)

		if err != nil {
			return result, err
		}

		messages, _ := parser.ParseSlackMessagesFromS3(fileContent)
		messageCount += len(messages)

		messageGroup := dtos.SlackMessagesGroup{
			Context:      subFolderName,
			MessageCount: messageCount,
			Messages:     messages,
		}

		result = append(result, messageGroup)

	}

	if err != nil {
		return nil, err
	}

	return result, err
}
func FindUserDataInS3JSONFile(userName string, folderName string, fileName string) (result dtos.SlackUser, err error) {
	// Read the JSON file from S3

	fullFilePath := fmt.Sprintf("%s/%s", folderName, fileName)
	contents, err := ReadS3JSONFile(fullFilePath)
	if err != nil {
		return dtos.SlackUser{}, err
	}

	var users []dtos.SlackUser

	// Parse the JSON file into a slice of SlackUser objects
	err = UnmarshalData(contents, &users)
	if err != nil {
		return dtos.SlackUser{}, err
	}

	// Find the SlackUser object with the matching username
	result = slack.FilterSlackUsersByUserName(userName, users)

	return result, nil
}

func FindDataInS3JSONFiles(dataType DataType, userId, mainFolder string) (result []dtos.SlackHistoryGroup, err error) {
	var (
		fileName     string
		filterMethod func(string, []byte) ([]string, error)
	)

	switch dataType {
	case Channel:
		fileName = "channels.json"
		filterMethod = func(userId string, contents []byte) ([]string, error) {
			var channels []dtos.SlackChannel
			err := UnmarshalData(contents, &channels)
			if err != nil {
				return nil, err
			}
			filteredChannels := slack.FilterChannelsByUserID(userId, channels)
			groupNames := make([]string, len(filteredChannels))
			for i, ch := range filteredChannels {
				groupNames[i] = ch.Name
			}
			return groupNames, nil
		}
	case Group:
		fileName = "groups.json"
		filterMethod = func(userId string, contents []byte) ([]string, error) {
			var groups []dtos.SlackGroup
			err := UnmarshalData(contents, &groups)
			if err != nil {
				return nil, err
			}
			filteredGroups := slack.FilterGroupsByUserID(userId, groups)
			groupNames := make([]string, len(filteredGroups))
			for i, group := range filteredGroups {
				groupNames[i] = group.Name
			}
			return groupNames, nil
		}
	case Mpim:
		fileName = "mpims.json"
		filterMethod = func(userId string, contents []byte) ([]string, error) {
			var mpims []dtos.MPIM
			err := UnmarshalData(contents, &mpims)
			if err != nil {
				return nil, err
			}
			filteredMpims := slack.FilterMPIMsByUserID(userId, mpims)
			groupNames := make([]string, len(filteredMpims))
			for i, mpim := range filteredMpims {
				groupNames[i] = mpim.Name
			}
			return groupNames, nil
		}
	case Dm:
		fileName = "dms.json"
		filterMethod = func(userId string, contents []byte) ([]string, error) {
			var dms []dtos.SlackDM
			err := UnmarshalData(contents, &dms)
			if err != nil {
				return nil, err
			}
			filteredDms := slack.FilterDmsByUserID(userId, dms)
			groupNames := make([]string, len(filteredDms))
			for i, dm := range filteredDms {
				groupNames[i] = dm.ID
			}
			return groupNames, nil
		}
	}

	fullFilePath := fmt.Sprintf("%s/%s", mainFolder, fileName)

	// Read the JSON file from S3
	contents, err := ReadS3JSONFile(fullFilePath)
	if err != nil {
		return nil, err
	}

	dataFolders, err := filterMethod(userId, contents)
	if err != nil {
		return nil, err
	}

	// Create a progress bar
	progressBar := pb.New(len(dataFolders))

	templateString := fmt.Sprintf(`Processing %s: {{counters . }} {{ bar . "[" "=" ">" "_" "]" }} {{percent . }} {{ rtime . "ETA %s" }}`, dataType.String(), "%s")

	progressBar.SetTemplateString(templateString)

	// Set width for the progress bar
	progressBar.SetWidth(80)

	// Start the progress bar
	progressBar.Start()

	// Update the FindAndParseSlackMessagesInS3Files function to accept the progress bar
	result, err = FindAndParseSlackMessagesInS3Files(dataFolders, mainFolder, progressBar)

	progressBar.Finish()

	return result, err
}

func GetObjects(s3Client *s3.S3, folder string, prefix string, objects *[]*s3.Object) error {
	cfg, err := config.GetConfig()

	if err != nil {
		return fmt.Errorf("Missing configuration parameters")
	}

	var objectsInput = &s3.ListObjectsV2Input{}
	//Fix bug with the AWS SDK
	if prefix != "" {
		fmt.Println("Checking with prefix:", prefix)
		objectsInput = &s3.ListObjectsV2Input{
			Bucket: aws.String(cfg.Server.Bucket),
			Prefix: aws.String(prefix),
		}
	} else {
		fmt.Println("Checking without prefix")
		objectsInput = &s3.ListObjectsV2Input{
			Bucket: aws.String(cfg.Server.Bucket),
		}
	}

	result, err := s3Client.ListObjectsV2(objectsInput)
	if err != nil {
		return fmt.Errorf("failed to list objects: %v", err)
	}

	fmt.Printf("%d objects found on %s\n", len(result.Contents), folder)

	for _, object := range result.Contents {
		key := *object.Key
		if key == folder {
			continue
		}
		*objects = append(*objects, object)
		if strings.HasSuffix(key, "/") {
			// Recursively call getObjects for subfolders
			err := GetObjects(s3Client, key, prefix, objects)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteBucketObjects deletes all objects and subfolders in the specified bucket and folder.
// It can only delete 1000 files each time (AWS limitation)
func DeleteBucketObjects(folder string, prefix string) error {

	cfg, err := config.GetConfig()

	if err != nil {
		return fmt.Errorf("Missing configuration parameters")
	}
	sess, err := GetAWSSession()

	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	s3Client := s3.New(sess)

	var objects []*s3.Object
	err = GetObjects(s3Client, folder, prefix, &objects)
	if err != nil {
		return err
	}

	progressBar := pb.StartNew(len(objects))
	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	for _, object := range objects {
		wg.Add(1)
		go func(obj *s3.Object) {
			defer wg.Done()
			deleteObjectInput := &s3.DeleteObjectInput{
				Bucket: aws.String(cfg.Server.Bucket),
				Key:    obj.Key,
			}
			_, err := s3Client.DeleteObject(deleteObjectInput)
			if err != nil {
				errChan <- fmt.Errorf("failed to delete object %s: %v", *obj.Key, err)
			} else {
				progressBar.Increment()
			}
		}(object)
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

func ListFilesInS3Bucket(prefix string) (result []string, err error) {
	cfg, err := config.GetConfig()

	if err != nil {
		return nil, fmt.Errorf("Missing configuration parameters")
	}

	sess, err := GetAWSSession()

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.Server.Bucket),
		Prefix: &prefix,
	}

	err = svc.ListObjectsV2Pages(input, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, object := range page.Contents {
			result = append(result, *object.Key)
		}
		return true
	})

	if err != nil {
		return nil, err
	}

	return result, err
}

func UnmarshalData(byteValues []byte, data interface{}) error {
	err := json.Unmarshal(byteValues, data)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
		return err
	}

	return nil
}

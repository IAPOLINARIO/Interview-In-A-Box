package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/slack-viewer/pkg/config"
)

func GetAWSSession() (*session.Session, error) {

	cfg, err := config.GetConfig()

	if err != nil {
		return nil, fmt.Errorf("Error: Missing AWS configuration")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Server.Region),
		Credentials: credentials.NewStaticCredentials(cfg.Server.AccessKey, cfg.Server.SecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return sess, err
}

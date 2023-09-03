package aws

import (
	"fmt"

	"github.com/slack-viewer/pkg/dtos"
)

const (
	CacheFolderName   = "cache"
	FullReportName    = "full.json"
	DMsCacheFile      = "dms.json"
	ChannelsCacheFile = "channels.json"
	GroupsCacheFile   = "groups.json"
	MPIMsCacheFile    = "mpims.json"
	UserCacheFile     = "userdata.json"
)

func CheckIfDatatIsCachedInS3(userName, sourceDir, cacheFileName string) (*dtos.SlackHistory, error) {
	cachePath := fmt.Sprintf("%s/%s/%s", sourceDir, CacheFolderName, userName)

	// Initialize SlackHistory struct
	slackHistory := dtos.SlackHistory{Month: sourceDir}

	cacheFileExists, err := CheckIfFileExists(cachePath, cacheFileName)

	if err != nil {
		return nil, err
	}

	if cacheFileExists {
		fmt.Println("Reading from cache...")
		fullFilePath := fmt.Sprintf("%s/%s", cachePath, cacheFileName)
		cachedResult, err := ReadS3JSONFile(fullFilePath)

		if err != nil {
			return nil, err
		}

		// Parse the JSON file into a SlackHistory object
		err = UnmarshalData(cachedResult, &slackHistory)

		if err != nil {
			return nil, err
		}

		return &slackHistory, nil
	}

	fmt.Printf("Cache not found\n")
	return nil, err
}

func CacheResultsInS3(data *dtos.SlackHistory, userName, sourceDir string) error {

	fmt.Printf("Caching results for %s...\n", userName)

	//Saves the result as JSON file for caching
	outputDir := fmt.Sprintf("%s/cache/%s/full.json", sourceDir, userName)
	err := SaveSlackHistoryIntoS3BucketAsJSON(data, outputDir)

	if err != nil {
		return err
	}

	outputDir = fmt.Sprintf("%s/cache/%s/dms.json", sourceDir, userName)
	err = SaveSlackHistoryIntoS3BucketAsJSON(data.DMs, outputDir)

	if err != nil {
		return err
	}

	outputDir = fmt.Sprintf("%s/cache/%s/groups.json", sourceDir, userName)
	err = SaveSlackHistoryIntoS3BucketAsJSON(data.Groups, outputDir)

	if err != nil {
		return err
	}

	outputDir = fmt.Sprintf("%s/cache/%s/channels.json", sourceDir, userName)
	err = SaveSlackHistoryIntoS3BucketAsJSON(data.Channels, outputDir)

	if err != nil {
		return err
	}

	outputDir = fmt.Sprintf("%s/cache/%s/mpims.json", sourceDir, userName)
	err = SaveSlackHistoryIntoS3BucketAsJSON(data.Mpims, outputDir)

	if err != nil {
		return err
	}

	outputDir = fmt.Sprintf("%s/cache/%s/userdata.json", sourceDir, userName)
	err = SaveSlackHistoryIntoS3BucketAsJSON(data.User, outputDir)

	if err != nil {
		return err
	}

	return nil
}

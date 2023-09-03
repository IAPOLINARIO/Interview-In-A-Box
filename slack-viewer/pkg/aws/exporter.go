package aws

import (
	"fmt"
	"time"

	"github.com/slack-viewer/pkg/dtos"
	"github.com/slack-viewer/pkg/slack"
)

func BuildSlackHistoryFromS3Bucket(userName, sourceDir string) (*dtos.SlackHistory, error) {

	// Start the timer to track how long it takes to generate the report
	start := time.Now()

	// Initialize SlackHistory struct
	slackHistory := dtos.SlackHistory{
		Month:    sourceDir,
		User:     dtos.SlackUser{},
		DMs:      []dtos.SlackHistoryGroup{},
		Mpims:    []dtos.SlackHistoryGroup{},
		Groups:   []dtos.SlackHistoryGroup{},
		Channels: []dtos.SlackHistoryGroup{},
	}

	userData, err := FindUserDataInS3JSONFile(userName, sourceDir, "users.json")

	if err != nil {
		return &slackHistory, fmt.Errorf("Error: Failed to connect to S3 bucket and retrieve users.json data %v\n", err)
	}

	// Check if user data is not empty
	if userData.Name == "" {
		return &slackHistory, fmt.Errorf("User data not found for %s", userName)
	}

	slackHistory.User = userData

	// Get DMs data from JSON files
	dms, err := FindDataInS3JSONFiles(Dm, userData.ID, sourceDir)
	if err != nil {
		return &slackHistory, err
	}
	slackHistory.DMs = dms

	// Get MPIMs data from JSON files
	mpims, err := FindDataInS3JSONFiles(Mpim, userData.ID, sourceDir)
	if err != nil {
		return &slackHistory, err
	}
	slackHistory.Mpims = mpims

	// Get groups data from JSON files
	groups, err := FindDataInS3JSONFiles(Group, userData.ID, sourceDir)
	if err != nil {
		return &slackHistory, err
	}

	filteredGroups := slack.FilterMessagesContext(groups, userData)
	slackHistory.Groups = filteredGroups

	// Get channels data from JSON files
	channels, err := FindDataInS3JSONFiles(Channel, userData.ID, sourceDir)
	if err != nil {
		return &slackHistory, err
	}
	filteredChannels := slack.FilterMessagesContext(channels, userData)
	slackHistory.Channels = filteredChannels

	//Cache results in S3 Bucket
	err = CacheResultsInS3(&slackHistory, userName, sourceDir)

	if err != nil {
		return &slackHistory, err
	}

	// Calculate how long it took to generate the report, cache it and print
	elapsed := time.Since(start)
	fmt.Printf("\nSlack history took %s to complete.\n\n", elapsed)

	return &slackHistory, nil
}

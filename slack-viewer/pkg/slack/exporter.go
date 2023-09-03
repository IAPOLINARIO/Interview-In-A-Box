package slack

import (
	"fmt"
	"time"

	"github.com/slack-viewer/pkg/dtos"
)

// BuildSlackHistoryStats generates a `SlackHistoryStats` object based on the given `SlackHistory` object.
// This function calculates various statistics about the user's Slack history, including the number of DMs,
// MPIMs, groups, and channels, as well as the total number of messages and the first and last message dates.
//
// Parameters:
//   - slackHistory: A `SlackHistory` object containing the user's Slack history data.
//
// Returns:
//   - A pointer to a `SlackHistoryStats` object containing the calculated statistics.
//   - An error, if any.
func BuildSlackHistoryStats(slackHistory *dtos.SlackHistory) (*dtos.SlackHistoryStats, error) {
	stats := &dtos.SlackHistoryStats{
		UserDisplayName: slackHistory.User.Name,
		TotalDMs:        len(slackHistory.DMs),
		TotalMpims:      len(slackHistory.Mpims),
		TotalGroups:     len(slackHistory.Groups),
		TotalChannels:   len(slackHistory.Channels),
	}

	return stats, nil
}

// BuildSlackHistoryFromJSONFiles builds a Slack history object from JSON files
// containing user data, DMs, MPIMs, groups, and channels data. It takes in the
// username, source folder path, and output directory path as parameters, and
// returns a pointer to the built Slack history object and an error (if any).
func BuildSlackHistoryFromJSONFiles(userName string, sourceFolder string) (*dtos.SlackHistory, error) {
	// Find user data in JSON file
	userData, err := FindUserDataInJSONFile(userName, sourceFolder)
	if err != nil {
		return nil, err
	}

	// Check if user data is not empty
	if userData.Name == "" {
		return nil, fmt.Errorf("User data not found for %s", userName)
	}

	// Initialize SlackHistory struct
	slackHistory := dtos.SlackHistory{
		User:     userData,
		DMs:      []dtos.SlackHistoryGroup{},
		Mpims:    []dtos.SlackHistoryGroup{},
		Groups:   []dtos.SlackHistoryGroup{},
		Channels: []dtos.SlackHistoryGroup{},
	}

	// Start the timer to track how long it takes to generate the report
	start := time.Now()

	// Get DMs data from JSON files
	dms, err := getDMsDataFromJSONFiles(userData, sourceFolder)
	if err != nil {
		return nil, err
	}
	slackHistory.DMs = dms

	// Get MPIMs data from JSON files
	mpims, err := getMpimsDataFromJSONFiles(userData, sourceFolder)
	if err != nil {
		return nil, err
	}
	slackHistory.Mpims = mpims

	// Get groups data from JSON files
	groups, err := getGroupsDataFromJSONFiles(userData, sourceFolder)
	if err != nil {
		return nil, err
	}
	slackHistory.Groups = groups

	// Get channels data from JSON files
	channels, err := getChannelsDataFromJSONFiles(userData, sourceFolder)
	if err != nil {
		return nil, err
	}
	slackHistory.Channels = channels

	// Calculate how long it took to generate the report and print it
	elapsed := time.Since(start)
	fmt.Printf("\nSlack history took %s to complete.\n\n", elapsed)

	return &slackHistory, nil
}

func getChannelsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserChannelsDataInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getGroupsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserGroupsDataInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getMpimsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserMPIMsInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getDMsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserDMsInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil

}

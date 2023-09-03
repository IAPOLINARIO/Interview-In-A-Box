package slack

import (
	"fmt"
	"io/ioutil"

	"github.com/slack-viewer/pkg/dtos"
	"github.com/slack-viewer/pkg/parser"
)

func FindUserDataInJSONFile(userName string, path string) (result dtos.SlackUser, err error) {

	paths := []string{path}
	userFilePath, err := FindFileInListOfPaths(defaultUserFileName, paths)

	if err != nil {
		return dtos.SlackUser{}, err
	}

	if len(userFilePath) > 0 {
		users, err := parser.ParseUsersFromJson(userFilePath[0]) //UserData will always be the same, so no need to run through all the results

		if err != nil {
			return dtos.SlackUser{}, err
		}

		result = FilterSlackUsersByUserName(userName, users)
	}

	return result, nil
}

func FindUserMPIMsInJSONFiles(user dtos.SlackUser, mpimsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{mpimsPath}

	mpimFilePath, err := FindFileInListOfPaths(defaultMPIMFileName, paths)

	if err != nil {
		panic(err)
	}

	mpims, err := parser.ParseMPImsFromJson(mpimFilePath[0]) //Only the first element matters in this case

	if err != nil {
		panic(err)
	}

	filteredMPIMs := FilterMPIMsByUserID(user.ID, mpims)

	mpimFolders := []string{}

	for _, mpim := range filteredMPIMs {
		mpimFolders = append(mpimFolders, mpim.Name)
	}

	result = FindAnParseSlackMessagesInJSONFiles(mpimFolders, mpimsPath)

	return result, err
}

func FindUserDMsInJSONFiles(user dtos.SlackUser, dmsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{dmsPath}
	dmFilePath, err := FindFileInListOfPaths(defaultDMFileName, paths)

	if err != nil {
		panic(err)
	}

	dms, err := parser.ParseDMsFromJson(dmFilePath[0]) //Only the first element matters in this case

	if err != nil {
		panic(err)
	}

	filteredDms := FilterDmsByUserID(user.ID, dms)

	dmFolders := []string{}

	for _, dm := range filteredDms {
		dmFolders = append(dmFolders, dm.ID)
	}

	result = FindAnParseSlackMessagesInJSONFiles(dmFolders, dmsPath)

	return result, err
}

func FindUserGroupsDataInJSONFiles(user dtos.SlackUser, groupsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{groupsPath}
	groupFilePath, err := FindFileInListOfPaths(defaultGroupsFileName, paths)

	if err != nil {
		panic(err)
	}

	groups, err := parser.ParseGroupsFromJson(groupFilePath[0]) //Only the first element matters in this case

	if err != nil {
		panic(err)
	}

	filteredGroups := FilterGroupsByUserID(user.ID, groups)

	groupsFolders := []string{}

	for _, group := range filteredGroups {
		groupsFolders = append(groupsFolders, group.Name)
	}

	allMessages := FindAnParseSlackMessagesInJSONFiles(groupsFolders, groupsPath)

	result = FilterMessagesContext(allMessages, user)

	return result, err
}

func ParseJsonFilesIntoSlackMessagesGroup(jsonPath string, context string) []dtos.SlackMessagesGroup {
	messagesGroup := []dtos.SlackMessagesGroup{}

	//result := [][]SlackMessage{}
	files, err := ioutil.ReadDir(jsonPath)
	messageCount := 0
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%d files found at %s \n", len(files), fullJSONPath)
	for _, file := range files {
		messages, _ := parser.ParseSlackMessagesFromJson(fmt.Sprintf("%s\\%s", jsonPath, file.Name()))
		messageCount += len(messages)

		messageGroup := dtos.SlackMessagesGroup{
			Context:      context,
			MessageCount: messageCount,
			Messages:     messages,
		}

		messagesGroup = append(messagesGroup, messageGroup)

	}

	if err != nil {
		panic(err)
	}

	return messagesGroup
}

func FindUserChannelsDataInJSONFiles(user dtos.SlackUser, channelsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{channelsPath}
	channelsFilePath, err := FindFileInListOfPaths(defaultChannelsFileName, paths)

	if err != nil {
		panic(err)
	}

	channels, err := parser.ParseChannelsFromJson(channelsFilePath[0]) //Only the first element matters in this case

	if err != nil {
		panic(err)
	}

	filteredGroups := FilterChannelsByUserID(user.ID, channels)

	channelsFolders := []string{}

	for _, channel := range filteredGroups {
		channelsFolders = append(channelsFolders, channel.Name)
	}

	allMessages := FindAnParseSlackMessagesInJSONFiles(channelsFolders, channelsPath)

	result = FilterMessagesContext(allMessages, user)

	return result, err
}

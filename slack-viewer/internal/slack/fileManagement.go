package slack

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/slack-viewer/internal/dtos"
	"github.com/slack-viewer/internal/jsonparser"
)

const (
	defaultUserFileName     = "users.json"
	defaultDMFileName       = "dms.json"
	defaultMPIMFileName     = "mpims.json"
	defaultGroupsFileName   = "groups.json"
	defaultChannelsFileName = "channels.json"
)

func findFileInListOfPaths(fileName string, paths []string) (result []string, err error) {

	for _, path := range paths {
		items, err := ioutil.ReadDir(path)

		if err != nil {
			return result, err
		}

		for _, item := range items {
			if !item.IsDir() {
				if item.Name() == fileName {
					result = append(result, fmt.Sprintf("%s\\%s", path, item.Name()))

					if err != nil {
						panic(err)
					}

					return result, err
				}
			}
		}
	}

	return result, err
}

func FindUserDataInJSONFile(userName string, path string) (dtos.SlackUser, error) {

	paths := []string{path}
	userFilePath, err := findFileInListOfPaths(defaultUserFileName, paths)

	if err != nil {
		panic(err)
	}

	users, err := jsonparser.ParseUsersFromJson(userFilePath[0]) //UserData will always be the same, so no need to run through all the results

	if err != nil {
		panic(err)
	}

	result := FilterSlackUsersByUserName(userName, users)

	return result, nil
}

func FilterSlackUsersByUserName(userName string, users []dtos.SlackUser) (result dtos.SlackUser) {

	for _, user := range users {
		if user.Name == userName {
			return user
		}
	}

	return result
}

func FindUserMPIMsInJSONFiles(user dtos.SlackUser, mpimsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{mpimsPath}

	mpimFilePath, err := findFileInListOfPaths(defaultMPIMFileName, paths)

	if err != nil {
		panic(err)
	}

	mpims, err := jsonparser.ParseMPImsFromJson(mpimFilePath[0]) //Only the first element matters in this case

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
	dmFilePath, err := findFileInListOfPaths(defaultDMFileName, paths)

	if err != nil {
		panic(err)
	}

	dms, err := jsonparser.ParseDMsFromJson(dmFilePath[0]) //Only the first element matters in this case

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
	groupFilePath, err := findFileInListOfPaths(defaultGroupsFileName, paths)

	if err != nil {
		panic(err)
	}

	groups, err := jsonparser.ParseGroupsFromJson(groupFilePath[0]) //Only the first element matters in this case

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

func FindUserChannelsDataInJSONFiles(user dtos.SlackUser, channelsPath string) (result []dtos.SlackHistoryGroup, err error) {
	paths := []string{channelsPath}
	channelsFilePath, err := findFileInListOfPaths(defaultChannelsFileName, paths)

	if err != nil {
		panic(err)
	}

	channels, err := jsonparser.ParseChannelsFromJson(channelsFilePath[0]) //Only the first element matters in this case

	if err != nil {
		panic(err)
	}

	filteredGroups := filterChannelsByUserID(user.ID, channels)

	channelsFolders := []string{}

	for _, channel := range filteredGroups {
		channelsFolders = append(channelsFolders, channel.Name)
	}

	allMessages := FindAnParseSlackMessagesInJSONFiles(channelsFolders, channelsPath)

	result = FilterMessagesContext(allMessages, user)

	return result, err
}

func FilterMessagesContext(slackHistory []dtos.SlackHistoryGroup, user dtos.SlackUser) []dtos.SlackHistoryGroup {

	defaultContextMessageCount := 10

	result := []dtos.SlackHistoryGroup{}

	for _, hg := range slackHistory {
		historyGroup := dtos.SlackHistoryGroup{}
		filteredMessagesGroups := []dtos.SlackMessagesGroup{}

		for _, g := range hg.Groups {
			filteredMessages := []dtos.SlackMessage{}
			startMessagesIndex := 0
			endMessageIndex := len(g.Messages) - 1
			alreadyFiltered := false

			for mIndex, m := range g.Messages {
				// Check if user is mentioned in the message or if the user said something
				if !alreadyFiltered || (mIndex > endMessageIndex) {
					if m.User == user.ID || strings.Contains(m.Text, user.Name) {
						alreadyFiltered = true

						// Include previous 10 messages as context if user is active in the chat
						if mIndex-defaultContextMessageCount > startMessagesIndex {
							startMessagesIndex = mIndex - defaultContextMessageCount
						}

						// Include next 10 messages as context if user is active in the chat
						if mIndex+defaultContextMessageCount < endMessageIndex {
							endMessageIndex = mIndex + defaultContextMessageCount
						}

						// Append filtered messages to the list
						for i := startMessagesIndex; i <= endMessageIndex; i++ {
							filteredMessages = append(filteredMessages, g.Messages[i])
						}
					}
				}
			}

			// Append filtered messages to the list of filtered message groups
			if len(filteredMessages) > 0 {
				messageGroup := dtos.SlackMessagesGroup{
					Context:      g.Context,
					MessageCount: len(filteredMessages),
					Messages:     filteredMessages,
				}
				filteredMessagesGroups = append(filteredMessagesGroups, messageGroup)
			}
		}

		// Append filtered message groups to the list of history groups
		if len(filteredMessagesGroups) > 0 {
			historyGroup.GroupName = hg.GroupName
			historyGroup.Groups = filteredMessagesGroups
			result = append(result, historyGroup)
		}
	}

	return result
}

func FindAnParseSlackMessagesInJSONFiles(folders []string, basePath string) (result []dtos.SlackHistoryGroup) {

	for _, folder := range folders {

		filePath := fmt.Sprintf("%s\\%s", basePath, folder)

		historyGroup := dtos.SlackHistoryGroup{
			GroupName: folder,
			Groups:    ParseJsonFilesIntoSlackMessagesGroup(filePath, folder),
		}

		result = append(result, historyGroup)
	}

	return result
}

func FilterMPIMsByUserID(userID string, mpims []dtos.MPIM) (result []dtos.MPIM) {
	for _, mpim := range mpims {

		for _, member := range mpim.Members {
			if member == userID {
				result = append(result, mpim)
				break
			}
		}
	}

	return result
}

func FilterDmsByUserID(userID string, dms []dtos.SlackDM) (result []dtos.SlackDM) {
	for _, dm := range dms {

		for _, member := range dm.Members {
			if member == userID {
				result = append(result, dm)
				break
			}
		}
	}

	return result
}

func FilterGroupsByUserID(userID string, groups []dtos.SlackGroup) (result []dtos.SlackGroup) {
	for _, group := range groups {

		for _, member := range group.Members {
			if member == userID {
				result = append(result, group)
				break
			}
		}
	}

	return result
}

func filterChannelsByUserID(userID string, channels []dtos.SlackChannel) (result []dtos.SlackChannel) {
	for _, channel := range channels {

		for _, member := range channel.Members {
			if member == userID {
				result = append(result, channel)
				break
			}
		}
	}

	return result
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
		messages, _ := jsonparser.ParseSlackMessagesFromJson(fmt.Sprintf("%s\\%s", jsonPath, file.Name()))
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

func GenerateSlackHistoryReport(slackHistory dtos.SlackHistory, fileoutput string) {
	f, err := os.Create(fileoutput)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	//HEADER
	header := fmt.Sprintf("Username: %s\nDMS: %d\nChannels: %d\nGroups: %d\nMPIMs: %d\n", slackHistory.User.Name, len(slackHistory.DMs), len(slackHistory.Channels), len(slackHistory.Groups), len(slackHistory.Mpims))

	_, err = f.WriteString(header)

	if err != nil {
		log.Fatal(err)
	}

	//DMS Section
	writeSectionReport(slackHistory.DMs, slackHistory.User, "DMS", fileoutput)

	//MPIMs Section
	writeSectionReport(slackHistory.Mpims, slackHistory.User, "MPIMs", fileoutput)

	//Groups Section
	writeSectionReport(slackHistory.Groups, slackHistory.User, "Groups", fileoutput)

	//Channels Section
	writeSectionReport(slackHistory.Groups, slackHistory.User, "Channels", fileoutput)
}

func writeSectionReport(group []dtos.SlackHistoryGroup, user dtos.SlackUser, sectionName string, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	sectionHeader := fmt.Sprintf("%s %s %s\n\n", strings.Repeat("=", 100), sectionName, strings.Repeat("=", 100))
	_, err = f.WriteString(sectionHeader)

	if err != nil {
		log.Fatal(err)
	}

	WriteMessagesReport(group, fileName)

	sectionFooter := fmt.Sprintf("\n%s %s %s\n\n", strings.Repeat("*", 100), fmt.Sprintf("END OF %s", sectionName), strings.Repeat("*", 100))
	_, err = f.WriteString(sectionFooter)

	if err != nil {
		log.Fatal(err)
	}

}

func WriteMessagesReport(slackGroups []dtos.SlackHistoryGroup, fileName string) {

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, item := range slackGroups {

		GroupSubHeader := fmt.Sprintf("%s %s %s\n", strings.Repeat("-", 100), item.GroupName, strings.Repeat("-", 100))
		_, err = f.WriteString(GroupSubHeader)

		if err != nil {
			log.Fatal(err)
		}

		for _, group := range item.Groups {
			for _, message := range group.Messages {
				parsedTs := ConvertTimeStamp(message.Ts)
				messageText := fmt.Sprintf("[%s] %s: %s \n", parsedTs, message.UserProfile.DisplayName, message.Text)

				_, err2 := f.WriteString(messageText)

				if err2 != nil {
					log.Fatal(err2)
				}

			}
		}
	}
}

func ConvertTimeStamp(ts string) string {
	unixTime, err := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	if err != nil {
		panic(err)
	}
	seconds, err := time.ParseDuration(ts + "s")
	if err != nil {
		panic(err)

	}
	t := unixTime.Add(seconds)

	return t.Format("2006-01-02 15:04:05")

}

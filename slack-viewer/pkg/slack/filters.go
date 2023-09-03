package slack

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/slack-viewer/pkg/dtos"
)

func FindFileInListOfPaths(fileName string, paths []string) (result []string, err error) {

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

func FilterSlackUsersByUserName(userName string, users []dtos.SlackUser) (result dtos.SlackUser) {

	for _, user := range users {
		if user.Name == userName {
			return user
		}
	}

	return result
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
					if m.User == user.ID || strings.Contains(m.Text, user.Name) || strings.Contains(g.Context, user.Name) || strings.Contains(hg.GroupName, user.Name) {
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

func FilterChannelsByUserID(userID string, channels []dtos.SlackChannel) (result []dtos.SlackChannel) {
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

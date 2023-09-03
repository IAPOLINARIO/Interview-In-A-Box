package slack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-viewer/pkg/dtos"
)

func GenerateSlackHistoryReport(slackHistory dtos.SlackHistory, fileoutput string) error {
	f, err := os.Create(fileoutput)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	//HEADER
	header := fmt.Sprintf("Username: %s\nDMS: %d\nChannels: %d\nGroups: %d\nMPIMs: %d\n", slackHistory.User.Name, len(slackHistory.DMs), len(slackHistory.Channels), len(slackHistory.Groups), len(slackHistory.Mpims))

	_, err = f.WriteString(header)

	if err != nil {
		return err
	}

	//DMS Section
	WriteSectionReport(slackHistory.DMs, slackHistory.User, "DMS", fileoutput)

	//MPIMs Section
	WriteSectionReport(slackHistory.Mpims, slackHistory.User, "MPIMs", fileoutput)

	//Groups Section
	WriteSectionReport(slackHistory.Groups, slackHistory.User, "Groups", fileoutput)

	//Channels Section
	WriteSectionReport(slackHistory.Groups, slackHistory.User, "Channels", fileoutput)

	return nil
}

func WriteSectionReport(group []dtos.SlackHistoryGroup, user dtos.SlackUser, sectionName string, fileName string) {
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

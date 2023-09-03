package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/slack-viewer/internal/dtos"
	"github.com/slack-viewer/internal/jsonparser"
)

func PersistSlackMessages(jSONPath string, channelNameFromMessages string, mongoManager MongoManager) (totalRecords int) {

	messages, _ := jsonparser.ParseSlackMessagesFromJson(jSONPath)

	slackMessages := []interface{}{}

	for i := range messages {
		slackMessages = append(slackMessages, messages[i])
	}

	mongoManager.saveRecordsToMongoDB(slackMessages, fmt.Sprintf("%s_%s", channelNameFromMessages, mongoManager.MessagesColletionName))

	totalRecords = len(messages)

	return totalRecords
}

func PersistSlackGroups(jSONPath string, mongoManager MongoManager) (totalRecords int) {

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jSONPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var records []dtos.SlackGroup

	err = json.Unmarshal(byteValues, &records)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	groupsRecords := []interface{}{}

	for i := range records {
		groupsRecords = append(groupsRecords, records[i])
	}

	mongoManager.saveRecordsToMongoDB(groupsRecords, mongoManager.GroupsColletionName)

	totalRecords = len(records)

	return totalRecords
}

func PersistSlackUser(jSONPath string, mongoManager MongoManager) (totalRecords int) {

	records, err := jsonparser.ParseSlackMessagesFromJson(jSONPath)

	if err != nil {
		panic(err)
	}

	userRecords := []interface{}{}

	for i := range records {
		userRecords = append(userRecords, records[i])
	}

	mongoManager.saveRecordsToMongoDB(userRecords, mongoManager.UsersColletionName)

	totalRecords = len(records)

	return totalRecords
}

func PersistSlackChannels(jSONPath string, mongoManager MongoManager) (totalRecords int) {

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jSONPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var channels []dtos.SlackChannel

	err = json.Unmarshal(byteValues, &channels)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	channelRecords := []interface{}{}

	for i := range channels {
		channelRecords = append(channelRecords, channels[i])
	}

	mongoManager.saveRecordsToMongoDB(channelRecords, mongoManager.ChannelsColletionName)

	totalRecords = len(channels)

	return totalRecords
}

func PersistSlackDMsJson(jSONPath string, mongoManager MongoManager) (totalRecords int) {

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jSONPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var DMs []dtos.SlackDM

	err = json.Unmarshal(byteValues, &DMs)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	dmsRecords := []interface{}{}

	for i := range DMs {
		dmsRecords = append(dmsRecords, DMs[i])
	}

	mongoManager.saveRecordsToMongoDB(dmsRecords, mongoManager.DmsColletionName)

	totalRecords = len(DMs)

	return totalRecords
}

func PersistSlackMPIMsJson(jSONPath string, mongoManager MongoManager) (totalRecords int) {

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jSONPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var MPIMs []dtos.MPIM

	err = json.Unmarshal(byteValues, &MPIMs)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	mpimsRecords := []interface{}{}

	for i := range MPIMs {
		mpimsRecords = append(mpimsRecords, MPIMs[i])
	}

	mongoManager.saveRecordsToMongoDB(mpimsRecords, mongoManager.DmsColletionName)

	totalRecords = len(MPIMs)

	return totalRecords
}

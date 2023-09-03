package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/slack-viewer/pkg/dtos"
)

func ParseChannelsFromJson(jsonPath string) ([]dtos.SlackChannel, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var channels []dtos.SlackChannel

	err = json.Unmarshal(byteValues, &channels)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return channels, err
}

func ParseGroupsFromJson(jsonPath string) ([]dtos.SlackGroup, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var groups []dtos.SlackGroup

	err = json.Unmarshal(byteValues, &groups)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return groups, err
}

func ParseMPImsFromJson(jsonPath string) ([]dtos.MPIM, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var mpims []dtos.MPIM

	err = json.Unmarshal(byteValues, &mpims)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return mpims, err
}

func ParseDMsFromJson(jsonPath string) ([]dtos.SlackDM, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var dms []dtos.SlackDM

	err = json.Unmarshal(byteValues, &dms)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return dms, err
}

func ParseUsersFromJson(jsonPath string) ([]dtos.SlackUser, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}

	var users []dtos.SlackUser

	err = json.Unmarshal(byteValues, &users)
	// Declare an empty slice for the MongoFields docs
	return users, err
}

func ParseSlackMessagesFromJson(jsonPath string) ([]dtos.SlackMessage, error) {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}
	// Declare an empty slice for the MongoFields docs
	var messages []dtos.SlackMessage

	err = json.Unmarshal(byteValues, &messages)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return messages, err
}

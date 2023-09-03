package parser

import (
	"encoding/json"
	"fmt"

	"github.com/slack-viewer/pkg/dtos"
)

func ParseSlackMessagesFromS3(byteValues []byte) ([]dtos.SlackMessage, error) {
	// Declare an empty slice for the MongoFields docs
	var messages []dtos.SlackMessage

	err := json.Unmarshal(byteValues, &messages)

	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}

	return messages, err
}

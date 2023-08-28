package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
)

const logPath = "/usercode/FILESYSTEM/logs"

type Output struct {
	Fields []string
	lines  []string
}

func main() {
	err := readLogFiles(logPath)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func parseKeyValue(line string) map[string]string {
	result := make(map[string]string)

	key := ""
	value := ""
	isValue := false
	quoteStart := false

	for i := 0; i < len(line); i++ {
		currentChar := line[i]

		if currentChar == '=' && !isValue {
			isValue = true
		} else if isValue && currentChar == '"' {
			quoteStart = !quoteStart
		} else if isValue && (currentChar == ' ' && !quoteStart || i == len(line)-1) {
			isValue = false

			if i == len(line)-1 && currentChar != ' ' {
				value += string(currentChar)
			}

			result[key] = value
			key = ""
			value = ""

		} else {

			if isValue {
				value += string(currentChar)
			} else {
				key += string(currentChar)
			}
		}

	}

	return result
}

func readLogFiles(path string) error {
	files, err := os.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			parseLogFile(path, file)
		}
	}

	return nil
}

func parseLogFile(filePath string, file fs.DirEntry) error {

	fields := make(map[string]bool)

	var results []map[string]string

	content, err := os.Open(fmt.Sprintf("%s/%s", filePath, file.Name()))

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		pairs := parseKeyValue(line)
		results = append(results, pairs)

		for key := range pairs {
			fields[key] = true
		}

	}

	sort.Slice(results, func(i, j int) bool {
		id1, err1 := strconv.Atoi(results[i]["id"])
		id2, err2 := strconv.Atoi(results[j]["id"])

		if err1 != nil || err2 != nil {
			return results[i]["id"] < results[j]["id"]
		}

		return id1 < id2
	})

	output := make(map[string]interface{})
	var logsOut []map[string]interface{}

	for _, logMap := range results {
		tempValue := make(map[string]interface{})

		for field := range fields {
			if val, exists := logMap[field]; exists {
				tempValue[field] = val
			} else {
				tempValue[field] = nil
			}
		}

		logsOut = append(logsOut, tempValue)
	}

	fieldsKeys := make([]string, 0, len(fields))

	for k := range fields {
		fieldsKeys = append(fieldsKeys, k)
	}

	output["fields"] = fieldsKeys
	output["lines"] = logsOut

	jsonOutput, err := json.MarshalIndent(output, "", " ")

	if err != nil {
		return err
	}

	//fmt.Print(fields)

	fmt.Println(string(jsonOutput))

	return nil
}

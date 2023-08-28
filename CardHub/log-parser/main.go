package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile("\\[(\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2})\\] \\[(\\w+)\\] \\[(\\w+)\\]: (.+)$")

func main() {
	args := os.Args

	path := args[1]
	logLevel := args[2]
	module := args[3]

	//fmt.Println(path, logLevel)

	readFiles(path, logLevel, module)
}

type LogSummary struct {
	SeveretyCount map[string]int
	ModuleCount   map[string]int
}

func readFiles(dir string, logLevel string, module string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var validFiles []fs.DirEntry

	for _, file := range files {
		if !file.IsDir() {
			validFiles = append(validFiles, file)
		}
	}

	results := make(chan LogSummary, len(validFiles))

	for _, file := range validFiles {
		go parseFile(dir, file, logLevel, module, results)
	}

	// Aggregate results from goroutines
	finalSummary := LogSummary{
		SeveretyCount: make(map[string]int),
		ModuleCount:   make(map[string]int),
	}

	for range validFiles {
		result := <-results
		for k, v := range result.SeveretyCount {
			finalSummary.SeveretyCount[k] += v
		}
		for k, v := range result.ModuleCount {
			finalSummary.ModuleCount[k] += v
		}
	}

	fmt.Println(finalSummary)

}

func parseFile(path string, file fs.DirEntry, logFilter string, moduleFilter string, results chan LogSummary) {

	fmt.Printf("Filename: %s\n", file.Name())

	content, err := os.Open(fmt.Sprintf("%s/%s", path, file.Name()))
	if err != nil {
		fmt.Println("Error reading file:", err)
		results <- LogSummary{} // send empty result
		return
	}
	defer content.Close()

	localSummary := LogSummary{
		SeveretyCount: make(map[string]int),
		ModuleCount:   make(map[string]int),
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		matchValues := pattern.FindStringSubmatch(scanner.Text())

		if matchValues != nil {
			// date := matchValues[1] // This is not being used in the function, so commented it out.
			severityType := matchValues[2]
			moduleType := matchValues[3]
			//logMessage := matchValues[4]

			if strings.Contains(logFilter, severityType) && strings.Contains(moduleFilter, moduleType) {
				localSummary.ModuleCount[moduleType]++
				localSummary.SeveretyCount[severityType]++
				//fmt.Println(severityType, moduleType, logMessage)
			}
		} else {
			localSummary.ModuleCount["malformed"]++
		}
	}

	results <- localSummary
}

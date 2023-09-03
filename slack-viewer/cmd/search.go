package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-viewer/pkg/aws"
	"github.com/slack-viewer/pkg/slack"
	"github.com/spf13/cobra"
)

var searchUserName string
var searchSourceDir string
var searchOutputDir string
var searchEnvironment string

func init() {
	CmdSearchLocal.PersistentFlags().StringVarP(&searchUserName, "username", "u", "", "Username to search for in the Slack history files")
	CmdSearchLocal.PersistentFlags().StringVarP(&searchSourceDir, "source", "s", ".\\", "Source directory to search for Slack history files")
	CmdSearchLocal.PersistentFlags().StringVarP(&searchOutputDir, "output", "o", ".\\", "Output directory for the report")

	CmdSearchRemote.PersistentFlags().StringVarP(&searchUserName, "username", "u", "", "Username to search for in the Slack history files")
	CmdSearchRemote.PersistentFlags().StringVarP(&searchSourceDir, "source", "s", ".\\", "Source directory to search for Slack history files")
	CmdSearchRemote.PersistentFlags().StringVarP(&searchOutputDir, "output", "o", ".\\", "Output directory for the report")

}

var CmdSearchLocal = &cobra.Command{
	Use:   "search-local",
	Short: "Search for Slack history files in a local directory",
	Long:  `Search for Slack history files containing messages from a specific user.`,
	Run: func(cmd *cobra.Command, args []string) {

		slackHistory, err := slack.BuildSlackHistoryFromJSONFiles(searchUserName, searchSourceDir)
		if err != nil {
			fmt.Println("Error generating the report: ", err)
			os.Exit(1)
		}

		slack.GenerateSlackHistoryReport(*slackHistory, searchOutputDir)
	},
}

var CmdSearchRemote = &cobra.Command{
	Use:   "search-remote",
	Short: "Search for Slack history files in a remote directory (S3)",
	Long:  `Search for Slack history files containing messages from a specific user.`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()
		//NO NEED TO CHECK IF THE REPORT IS CACHED OR NOT, BECAUSE IT WAS ALREADY DONE IN THE FUNCTION BEFORE CALLING BuildSlackHistoryFromS3Bucket
		cachedInS3Result, err := aws.CheckIfDatatIsCachedInS3(searchUserName, searchSourceDir, "full.json")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if cachedInS3Result != nil {
			// Calculate how long it took to generate the report, cache it and print
			elapsed := time.Since(start)

			fmt.Printf("\nSlack history took %s to complete.\n\n", elapsed)

			slack.GenerateSlackHistoryReport(*cachedInS3Result, searchOutputDir)
		} else {

			result, err := aws.BuildSlackHistoryFromS3Bucket(searchUserName, searchSourceDir)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			slack.GenerateSlackHistoryReport(*result, searchOutputDir)
		}

	},
}

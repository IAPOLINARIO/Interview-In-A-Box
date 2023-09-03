package cmd

import (
	"fmt"
	"os"

	"github.com/slack-viewer/pkg/aws"
	"github.com/spf13/cobra"
)

var unzipFileName string
var unzipOutputFolder string

var CmdUnzip = &cobra.Command{
	Use:   "unzip",
	Short: "Start a web server for the Slack APIS",
	Long:  `Start a web server and expose the endpoints for the Slack Viewer Tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Unziping file...\n")

		err := aws.UnzipS3File(unzipFileName, unzipOutputFolder, 5)
		//err := aws.ListFilesInS3Bucket(unzipBucketName)

		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	},
}

func init() {
	CmdUnzip.PersistentFlags().StringVarP(&unzipFileName, "filename", "f", "file.zip", "The name of the file to unzip")
	CmdUnzip.PersistentFlags().StringVarP(&unzipOutputFolder, "output", "o", "/ouput", "The name of the output folder to unzip the files")
}

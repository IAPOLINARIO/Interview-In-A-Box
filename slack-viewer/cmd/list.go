package cmd

import (
	"fmt"
	"os"

	"github.com/slack-viewer/pkg/aws"
	"github.com/spf13/cobra"
)

var listPrefix string

var CmdListRemote = &cobra.Command{
	Use:   "list-remote",
	Short: "List content on a remote S3 bucket",
	Long:  `List content on S3 bucket`,
	Run: func(cmd *cobra.Command, args []string) {

		files, err := aws.ListFilesInS3Bucket(listPrefix)

		if err != nil {
			fmt.Printf("Error listing directory: %v \n", err)
			os.Exit(1)
		}

		for _, file := range files {
			fmt.Println(file)
		}

	},
}

func init() {
	CmdListRemote.PersistentFlags().StringVarP(&listPrefix, "prefix", "p", "", "Prefix (folder name) of the filename to search for")
}

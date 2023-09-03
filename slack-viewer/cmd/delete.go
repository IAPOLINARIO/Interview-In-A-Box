package cmd

import (
	"fmt"
	"os"

	"github.com/slack-viewer/pkg/aws"
	"github.com/spf13/cobra"
)

var deleteFolder string

var CmdDeleteRemote = &cobra.Command{
	Use:   "delete-remote",
	Short: "Delete files in a remote directory (S3)",
	Long:  `Delete all the S3 bucket files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := aws.DeleteBucketObjects(deleteFolder, "")

		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

	},
}

func init() {
	CmdDeleteRemote.PersistentFlags().StringVarP(&deleteFolder, "folder", "f", "/", "Name of the folder")
}

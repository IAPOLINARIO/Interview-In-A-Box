package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/slack-viewer/pkg/api"
)

var port int
var CmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start a web server for the Slack APIS",
	Long:  `Start a web server and expose the endpoints for the Slack Viewer Tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting the server...\n")
		// Start the server
		addr := fmt.Sprintf(":%d", port)
		err := api.StartServer(addr)
		if err != nil {
			log.Fatalf("Server Error: %v", err)
		}
	},
}

func init() {
	CmdServer.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port number to listen on")
}

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/slack-viewer/cmd"
	"github.com/spf13/cobra"
)

const Version = "0.1.5"

var rootCmd = &cobra.Command{
	Use:   "slackctl",
	Short: "Search your Slack history files from the command line with ease using our Slack command-line tool.",
	Long:  `Our command-line tool makes searching your Slack history a breeze. Easily find what you need without leaving the terminal. Configurable and simple to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define color styles
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		// Print the welcome message with colors
		fmt.Println("")
		fmt.Printf("%sWelcome to Slackctl!\n", green("[+] "))
		fmt.Printf("%sEnter a command to start exploring the %sslackctl%s tool !\n", yellow("[*] "), green(""), yellow(""))
		fmt.Printf("%sFor help, type '%s%s help%s', or try hacking the system with '%s%s matrix%s'.\n", yellow("[*] "), green(""), yellow(""), green(""), green(""), yellow(""), green(""))
		fmt.Printf("%sFor a list of available commands, use your hacker skills and type '%s--help%s'.\n", yellow("[*] "), yellow(""), green(""))
		fmt.Printf("%sAnd remember, you are the one who can use the Slack searching into high gear!\n", yellow("[*] "), green(""), yellow(""))
		fmt.Println("")

	},
}

func main() {
	rootCmd.AddCommand(cmd.CmdSearchLocal)
	rootCmd.AddCommand(cmd.CmdSearchRemote)
	rootCmd.AddCommand(cmd.CmdServer)
	rootCmd.AddCommand(cmd.CmdUnzip)
	rootCmd.AddCommand(cmd.CmdDeleteRemote)
	rootCmd.AddCommand(cmd.CmdListRemote)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

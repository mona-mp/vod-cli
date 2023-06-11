package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vod-cli",
	Short: "It is the vod cli which you can use to manage your videos",
	Long:  " VOD Client \nThis client helps you manage your videos in  VOD Service, get videos detail or update them.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("help", "h", false, "CLI help")
}

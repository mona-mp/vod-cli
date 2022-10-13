/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// updateDetailCmd represents the updateDetail command
var updateDetailCmd = &cobra.Command{
	Use:   "updateDetail",
	Short: "Update video title and description",
	Long: `This command is used to update the datail of the
	your video. You can pass the values to in by flags as a string.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the title and description value from user

		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")

		// call the function to update the details

		updateDetail(args, title, description)
	},
}

func init() {
	rootCmd.AddCommand(updateDetailCmd)

	// create flags
	updateDetailCmd.PersistentFlags().String("title", "", "Update title for video")
	updateDetailCmd.PersistentFlags().String("description", "", "Update description for video")
}

// send update request function
func updateDetail(args []string, title string, description string) {
	video_id := strings.Join(args, "")

	payload, err := json.Marshal(map[string]interface{}{
		"title":       title,
		"description": description,
	})
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	url := "https://napi.arvancloud.com/vod/2.0/videos/" + video_id
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))

	req.Header.Add("Authorization", readapikey())
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

}

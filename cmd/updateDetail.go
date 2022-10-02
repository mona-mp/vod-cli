/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// updateDetailCmd represents the updateDetail command
var updateDetailCmd = &cobra.Command{
	Use:   "updateDetail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateDetail(args)
	},
}

func init() {
	rootCmd.AddCommand(updateDetailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateDetailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateDetailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func updateDetail(args []string) {
	args2 := strings.Join(args, "")
	payload, err := json.Marshal(map[string]interface{}{
		"title":       "my simple todo",
		"description": "ey khooooda che giri kardim",
	})
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	url := "https://napi.arvancloud.com/vod/2.0/videos/" + args2
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	fmt.Println(req)
	req.Header.Add("Authorization", "Apikey XXXX")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}

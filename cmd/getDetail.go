/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// getDetailCmd represents the getDetail command
var getDetailCmd = &cobra.Command{
	Use:   "getDetail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getDetail(args)
	},
}

func init() {
	rootCmd.AddCommand(getDetailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDetailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDetailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDetail(args []string) {
	args2 := strings.Join(args, "")
	req, err := http.NewRequest(
		http.MethodGet,
		"https://napi.arvancloud.com/vod/2.0/videos/"+args2,
		nil,
	)

	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Authorization", "Apikey XXXX")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	defer res.Body.Close()
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))
}

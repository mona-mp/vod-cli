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
	Short: "Get video detail",
	Long: `Get the video detail using this command.example:
	arvan-vod-cli getDetail <video-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		getDetail(args)
	},
}

// add flags
func init() {
	rootCmd.AddCommand(getDetailCmd)
}

// sen GET request to get video detail function
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

	req.Header.Add("Authorization", readapikey())

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

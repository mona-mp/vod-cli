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

// getDetailCmd represents the getDetail command
var getDetailCmd = &cobra.Command{
	Use:   "getDetail",
	Short: "Get video detail",
	Long: `Get the video detail using this command.example:
	vod-cli getDetail <video-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		getDetail(args)
	},
}

// add flags
func init() {
	rootCmd.AddCommand(getDetailCmd)
}

// send GET request to get video detail function
func getDetail(args []string) {
	args2 := strings.Join(args, "")
	req, err := http.NewRequest(
		http.MethodGet,
		"https://test.com/vod/2.0/videos/"+args2,
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
	res2, err2 := PrettyString(string(responseBytes))
	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println(res2)

}

// function to get the prettyJSON
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `Log in to Arvan API and save login for subsequent use.example:
	arvan-vod-cli login <APIKEY>`,
	Run: func(cmd *cobra.Command, args []string) {
		addapikey(args)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

// get the config address
func getdirectory() string {
	homedir, _ := os.UserHomeDir()
	return homedir + "/.arvan-vod/config"
}

// add API key to the config file
func addapikey(args []string) {

	a2 := args[1]
	Apikey := []byte(a2)
	os.WriteFile(getdirectory(), Apikey, 0644)

}

// read the API key
func readapikey() (apikey string) {
	ApiKey, _ := ioutil.ReadFile(getdirectory())

	return "Apikey " + string(ApiKey)
}

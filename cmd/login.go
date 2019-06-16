/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/izolight/proxmox-api-go/api"
	"github.com/spf13/cobra"
)

var (
	apiUrl string
	username string
	password string
	debug bool
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs in to the Proxmox API",
	Long: `Logs in to Proxmox API and stores the token`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := api.NewClient(apiUrl, username, password)
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			fmt.Println(c)
		}
		users, err := c.GetUsers()
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			fmt.Println(users)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&apiUrl, "apiUrl", "a", "", "Url of the Proxmox API server")
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username for login")
	loginCmd.Flags().StringVarP(&password,"password", "p", "", "Password")
	loginCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	loginCmd.MarkFlagRequired("apiUrl")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

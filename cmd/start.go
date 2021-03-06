/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"agentx/core"
	"agentx/module"
	"log"
	"time"

	"github.com/posteo/go-agentx"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start agent",
	Long:  `Start agent to serve`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := agentx.Dial("tcp", "localhost:705")
		if err != nil {
			log.Fatalf("%s", err)
		}
		client.Timeout = 1 * time.Minute
		client.ReconnectInterval = 1 * time.Second

		session, err := client.Session()
		module.Init()
		core.RegisterModules(session)
		time.Sleep(3 * time.Minute)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

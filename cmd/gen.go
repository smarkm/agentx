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

	"github.com/spf13/cobra"
)

//
var (
	MibPath string
	Mib     string
	DryRun  bool
	Info    bool
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code base on loading mib",
	Long:  `Generate code base on loading mib`,
	Run: func(cmd *cobra.Command, args []string) {
		core.LoadAndGenerateCode(MibPath, Mib, DryRun)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&MibPath, "path", "p", "mibs", "mib path")
	genCmd.Flags().StringVarP(&Mib, "mib", "m", "", "moading mib,if no setting will load all under mib path")
	genCmd.Flags().BoolVarP(&DryRun, "dry-run", "d", false, "dry run will not generate file")
	genCmd.Flags().BoolVarP(&Info, "info-oid", "i", false, "print raw oid informaton from mib")
}

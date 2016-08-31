// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/plugins"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: up,
}

func init() {
	devctlCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func up(cmd *cobra.Command, args []string) {
	config := new(parser.ConfigurationStruct)
	config.ParseFileDefault()

	// create .devctl folder
	os.Mkdir(".devctl", 0700)

	pluginsUsed := plugins.Used(config)

	preInstall(config, pluginsUsed)
	install(config, pluginsUsed)
	postInstall(config, pluginsUsed)
}

func preInstall(config *parser.ConfigurationStruct, pluginsUsed []plugins.Plugin) {
	for _, i := range pluginsUsed {
		i.PreInstall(config)
	}
}

func install(config *parser.ConfigurationStruct, pluginsUsed []plugins.Plugin) {
	for _, i := range pluginsUsed {
		i.Install(config)
	}
}

func postInstall(config *parser.ConfigurationStruct, pluginsUsed []plugins.Plugin) {
	for _, i := range pluginsUsed {
		i.PostInstall(config)
	}
}

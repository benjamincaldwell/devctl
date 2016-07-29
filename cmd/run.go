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
	"sort"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/services"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func init() {
	devctlCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		utilities.ErrorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	scriptName := args[0]

	// TODO: Work your own magic here
	config := new(parser.ConfigurationStruct)
	config.ParseFile("./devctl.yaml")

	servicesUsed := services.ServicesUsed(config)
	scripts := make(map[string]utilities.RunCommand)

	for _, i := range servicesUsed {
		scripts = mapmerge(scripts, i.Scripts(config))
	}

	// merge most important last
	scripts = mapmerge(scripts, config.Scripts)

	if val, ok := scripts[scriptName]; ok {
		post := new(utilities.PostCommand)
		post.RunCommand(val.Command)
		post.Write()
	} else {
		// fuzzy search
		keys := make([]string, len(scripts))
		i := 0
		for k := range scripts {
			keys[i] = k
			i++
		}
		fuzzyFind := fuzzy.RankFind(scriptName, keys)
		sort.Sort(fuzzyFind)

		if len(fuzzyFind) > 0 {
			val := scripts[fuzzyFind[0].Target]
			post := new(utilities.PostCommand)
			post.RunCommand(val.Command)
			post.Write()
		} else {
			printer.Fail("%s script not found", scriptName)
		}
	}
}

func mapmerge(base map[string]utilities.RunCommand, mapsToMerge ...map[string]utilities.RunCommand) map[string]utilities.RunCommand {
	for _, m := range mapsToMerge {
		for key, value := range m {
			base[key] = value
		}
	}
	return base
}
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
	"log"

	"github.com/Eldius/game-manager-go/config"
	"github.com/spf13/cobra"
)

// setupMinecraftCmd represents the setupMinecraft command
var setupMinecraftCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a minecraft server",
	Long:  `Sets up a minecraft server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetAppConfig()
		log.Println(cfg)
	},
}

func init() {
	minecraftCmd.AddCommand(setupMinecraftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupMinecraftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupMinecraftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
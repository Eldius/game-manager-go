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

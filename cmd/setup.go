package cmd

import (
	"log"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/setup"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set's up the runtime environment",
	Long:  `Set's up the runtime environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetAppConfig()
		log.Println("Is environment ready?", setup.ValidateSetup(*cfg))
		setup.Setup(*cfg)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

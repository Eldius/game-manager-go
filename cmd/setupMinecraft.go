package cmd

import (
	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/logger"
	"github.com/Eldius/game-manager-go/provisioning"
	"github.com/Eldius/game-manager-go/scripts"
	"github.com/spf13/cobra"
)

// setupMinecraftCmd represents the setupMinecraft command
var setupMinecraftCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a minecraft server",
	Long:  `Sets up a minecraft server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetAppConfig()
		hostCfg := scripts.NewServerProvisioning("minecraft", hostIP, sshPort, hostIP, remoteUser, args)
		provisioning.Provision(*cfg, hostCfg)
		logger.Debug(hostCfg)
	},
}

var (
	sshKey     string
	sshPort    int
	hostIP     string
	remoteUser string
)

func init() {
	minecraftCmd.AddCommand(setupMinecraftCmd)

	setupMinecraftCmd.Flags().StringVarP(&sshKey, "ssh-key", "k", "~/.ssh/id_rsa", "The SSH key to log into remote server.")
	setupMinecraftCmd.Flags().IntVarP(&sshPort, "ssh-port", "p", 22, "The SSH port to connect in the remote server.")
	setupMinecraftCmd.Flags().StringVarP(&hostIP, "server", "s", "", "The host to configure.")
	setupMinecraftCmd.Flags().StringVarP(&remoteUser, "user", "u", "", "The user to log in the remote machine.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupMinecraftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupMinecraftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

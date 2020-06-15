/*
Package cmd is where all commands will be
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "game-manager-go",
	Short: "A simple tool to manage game servers",
	Long:  `A simple tool to manager game servers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.game-manager/config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Run in verbose mode")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	setupLog(home)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in $HOME/.game-manager directory with name "config.yml" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".game-manager"))
		viper.SetConfigName("config")
		viper.SetConfigType("yml")

		viper.SetDefault("workspace", filepath.Join(home, ".game-manager", "workspace"))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func setupLog(home string) {
	logDir := filepath.Join(home, ".game-manager", "log")
	os.MkdirAll(logDir, os.ModePerm)
	logFile := filepath.Join(logDir, "execution.log")
	fmt.Println("logging to", logFile)
	if f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

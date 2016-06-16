package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "lpn COMMAND [OPTIONS]",
	Short: "libphonenumber CLI tool",
	Long: `A CLI tool for working with the XML, JSON, and protobuf formats of the
libphonenumber library.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Reads in config file and ENV variables.
	cobra.OnInitialize(initConfig)

	// Global configuration setting.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lpnr.yaml)")
}

func initConfig() {
	// Enable ability to specify config file via flag.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	// Config file is named '.lpn.<ext>', loaded from the users home
	// directory, and overridden environment variables.
	viper.SetConfigName(".lpnr")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guardian360/go-lighthouse/cmd/lh/config"
	v2 "github.com/guardian360/go-lighthouse/cmd/lh/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	appName      = "lh"
	shortAppDesc = "A CLI tool for managing Lighthouse resources"
	longAppDesc  = `Lighthouse CLI is a command-line interface tool that allows users to manage and interact with Lighthouse resources efficiently. It provides various commands to perform operations such as creating, updating, deleting, and listing resources.`
)

var (
	rootCmd = &cobra.Command{
		Use:   appName + " [command]",
		Short: shortAppDesc,
		Long:  longAppDesc,
	}

	configFile string
)

func init() {
	rootCmd.AddCommand(v2.RootCmd)

	setupConfig()
}

func setupConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		viper.AddConfigPath(home + "/.config/" + appName)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	_ = viper.ReadInConfig()

	viper.Unmarshal(&config.Config)

	fmt.Printf("config.Config = %+v\n", config.Config)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

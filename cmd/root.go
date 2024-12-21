/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const logo = `
     /\_/\  . o O ( f cURL )
    ( o.o )  
     > ^ <  
      ~~~
█▄▀ █ ▀█▀ ▀█▀ █▄█   █▀ ▀█▀ █▄█ █   █▀▀
█░█ █ ░█░ ░█░ ░█░   ▄█ ░█░ ░█░ █▄▄ ██▄

`

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kitty-style",
	Short: "A user-friendly CLI tool for crafting and managing HTTP requests in the simplest way, with AI-powered query generation.",
	Long: `is user-friendly CLI tool for simplifying data transfers between servers. It is a cURL alternative that with intiutive quering
and readable responses. Say goodbye to complex cURL commands and hello to an elegant, user-friendly alternative.

1. Effortless query construction.
2. AI-powered query generation.
3. Beautifully formatted responses.

Meow Meow!
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logo)
		fmt.Print(cmd.Long)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kitty-style.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".kitty-style" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".kitty-style")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

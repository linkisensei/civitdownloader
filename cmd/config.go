/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/linkisensei/civitdownloader/app/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Sets the program configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		installation_path, _ := cmd.Flags().GetString("path")
		if installation_path != "" {
			config.Config.Set(config.INSTALLATION_PATH, installation_path)
			fmt.Printf("Automatic1111's installation path set to %s", installation_path)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().String("path", "", "Automatic1111's installation path")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/linkisensei/civitdownloader/app"
	"github.com/linkisensei/civitdownloader/app/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "civitdownloader",
	Short: "A simple tool for downloading Civit AI Models directly to Automatic1111",
	Long:  `A simple tool for downloading Civit AI Models directly to Automatic1111`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		redColor := color.New(color.FgRed).Add(color.Bold)
		greenColor := color.New(color.FgGreen).Add(color.Underline)

		fmt.Printf(" +-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+-+\n |C|i|v|i|t| |D|o|w|n|l|o|a|d|e|r|\n +-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+-+\n")

		fmt.Println(config.Config.Get(config.INSTALLATION_PATH))

		if config.Config.Get(config.INSTALLATION_PATH) == "" {
			redColor := color.New(color.FgRed).Add(color.Bold)
			redColor.Println("Automatic1111's installation path missing!")
			fmt.Println("Please set the path by executing this program with the following arguments \"config --path PATH_TO_AUTOMATIC1111_INSTALLATION_FOLDER\"\n")
			return
		}

		for {
			fmt.Printf("\n\nType \"exit\" to exit")
			greenColor.Printf("\nInsert the model URL: ")

			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				continue
			}

			// Cleaning Input
			input = strings.ReplaceAll(input, "\r\n", "")
			input = strings.ReplaceAll(input, "\n", "")

			if input == "exit" {
				break
			}

			if input != "" {
				err := app.DownloadModel(input)
				if err != nil {
					redColor.Printf("Error: %s", err.Error()+"\n")
				}
			}
		}
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

func init() {}

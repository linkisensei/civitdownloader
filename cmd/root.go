package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/linkisensei/civitdownloader/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installation_path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "civitdownloader",
	Short: "A simple tool for downloading Civit AI Models directly to Automatic1111",
	Long:  `A simple tool for downloading Civit AI Models directly to Automatic1111`,

	Run: func(cmd *cobra.Command, args []string) {

		redColor := color.New(color.FgRed).Add(color.Bold)
		greenColor := color.New(color.FgGreen).Add(color.Underline)

		fmt.Printf(" +-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+-+\n |C|i|v|i|t| |D|o|w|n|l|o|a|d|e|r|\n +-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+-+\n")

		fmt.Printf("\n   AUTOMATIC1111'S PATH: " + viper.GetString("path"))
		fmt.Printf("\n   Type \"exit\" to exit\n\n")

		installation_path := viper.GetString("path")

		if installation_path == "" {
			redColor := color.New(color.FgRed).Add(color.Bold)
			redColor.Println("Automatic1111's installation path missing!")
			fmt.Println("Please set the path by executing this program with the following arguments \"config --path PATH_TO_AUTOMATIC1111_INSTALLATION_FOLDER\"\n")
			return
		}

		for {

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
	rootCmd.PersistentFlags().String("path", "", "Automatic1111's installation path")
	if err := viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}

	cobra.OnInitialize(initConfig)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}
}

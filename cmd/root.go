/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/linkisensei/civitdownloader/app"
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
		// model, _ := civit.GetModel(25306)
		// fmt.Printf("MODEL: %s\n", model.Name)

		// fmt.Printf("LATEST VERSION: %s\n", model.GetVersion(0).Name)
		// downloadUrl, _ := model.Versions[0].GetDownloadUrl()
		// fmt.Printf("DOWNLOAD URL: %s\n", downloadUrl)

		// fmt.Printf("63765 VERSION: %s\n", model.GetVersion(63765).Name)
		// downloadUrl2, _ := model.GetVersion(63765).GetDownloadUrl()
		// fmt.Printf("DOWNLOAD URL: %s\n", downloadUrl2)

		// civit.createRequestInfoFromUrl("https://civitai.com/models/25306?modelVersionId=40541")

		app.DownloadModel("https://civitai.com/models/25306?modelVersionId=40541")
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.civitdownloader.yaml)")

}

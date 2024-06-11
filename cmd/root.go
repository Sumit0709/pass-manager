/*
Copyright © 2024 Sumit0709 sumitranjan327@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CommitId string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pass-manager",
	Short: "A Cli tool to manage your passwords safely and securely.",
	Long: `This application allows you to manage all your passwords using a single secret. All your passwords 
are encrypted and stored on your machine. There is no internet connectivity required.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This application is built from commit ID:", CommitId)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pass-manager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

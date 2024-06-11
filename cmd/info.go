/*
Copyright © 2023 Sumit0709 prasad28sumit@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Sumit0709/pass-manager/pkg/user"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints the location of data files",
	Long:  `This command is used to know about the location of the genereted data stored on the machine`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := user.Info()
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

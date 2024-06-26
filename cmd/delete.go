/*
Copyright © 2024 Sumit0709 sumitranjan327@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/Sumit0709/pass-manager/pkg/user"
	"golang.org/x/term"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes one or more user data",
	Long: `This command is for deleting one or more app data. It takes an argument secret and has flags for specifying 
the app data to be deleted`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Your secret: ")
		secret, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		fmt.Println()
		u := user.User{
			App:    cmd.Flag("app").Value.String(),
			Email:  cmd.Flag("email").Value.String(),
			UserId: cmd.Flag("userId").Value.String(),
		}
		var users []user.User
		users, err = user.Delete(u, string(secret), false)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(users) > 1 {
			var choice string
			for _, u := range users {
				fmt.Println(u.Print())
			}
			fmt.Println("More than one such user exists", "Do you want to delete all(Yes/No)?")
			fmt.Scanf("%s", &choice)
			choice = strings.ToUpper(choice)
			if choice == "YES" {
				user.Delete(u, string(secret), true)
			}
		}
		fmt.Println("Deleted!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	deleteCmd.Flags().StringP("app", "a", "", "App name")
	deleteCmd.Flags().StringP("email", "e", "", "Email")
	deleteCmd.Flags().StringP("userId", "u", "", "User ID")
}

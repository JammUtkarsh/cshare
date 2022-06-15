package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var u, pw string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to cshare server",
	Long:  `login to cshare server to enable text sharing.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	//TODO: required later
	//rootCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
	//rootCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
	//err := rootCmd.MarkPersistentFlagRequired("username")
	//if err != nil {
	//	return
	//}

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

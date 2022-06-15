package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "To send clips in a higher encrypted format.",
	Long: `To send clips in a much higher level encryption format than the usual http.
Use case: To share passwords and private ssh keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("secret called")
	},
}

func init() {
	rootCmd.AddCommand(secretCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secretCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secretCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

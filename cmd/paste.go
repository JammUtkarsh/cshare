package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pasteCmd represents the paste command
var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "paste the last copied text.",
	Long:  `paste the last copied text from server sent by other device.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("paste called")
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pasteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pasteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

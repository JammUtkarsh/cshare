package cmd

import (
	"errors"
	"github.com/JammUtkarsh/cshare/httpBin"
	"github.com/spf13/cobra"
)

var clip string
var cpCmd = &cobra.Command{
	Use: "cp",
	//TODO: Improve the documentation.
	Short: "add the text to be sent",
	Long:  `add the text to be sent across the device.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if clip == "" && len(args) < 1 {

			return errors.New("accept(s) 1 argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add http parsing methods
		if clip == "" {
			httpControllers.POSTrequest(args[0])
		} else {
			httpControllers.POSTrequest(clip)
		}
	},
	Example: `cshare cp text
cshare cp "hello world"
cshare cp -c "hello world"
`,
}

func init() {
	rootCmd.AddCommand(cpCmd)
	cpCmd.PersistentFlags().StringVarP(&clip, "clip", "c", "", "A string for cp command")
}

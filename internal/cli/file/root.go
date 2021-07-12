package file

import "github.com/spf13/cobra"

var CmdFile = &cobra.Command{
	Use:   "file [File Operation]",
	Short: "File Operation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CmdFile.AddCommand(cmdModel)
	CmdFile.AddCommand(cmdController)
}

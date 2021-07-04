package db

import (
	"github.com/spf13/cobra"
)

var CmdDb = &cobra.Command{
	Use:   "db [Database Operation]",
	Short: "Database Operation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CmdDb.AddCommand(cmdMigrate, cmdSeed)
}

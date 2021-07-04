package main

import (
	"ginadmin/cli/db"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ginadmin-cli"}
	rootCmd.AddCommand(db.CmdDb)
	rootCmd.Execute()

}

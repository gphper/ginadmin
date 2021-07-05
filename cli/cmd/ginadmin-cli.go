package main

import (
	"ginadmin/cli/db"
	files "ginadmin/cli/file"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ginadmin-cli"}
	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(files.CmdFile)
	rootCmd.Execute()

}

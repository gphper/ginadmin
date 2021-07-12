package main

import (
	"ginadmin/internal/cli/db"
	files "ginadmin/internal/cli/file"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ginadmin-cli"}
	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(files.CmdFile)
	rootCmd.Execute()

}

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 15:11:10
 */
package main

import (
	"github/gphper/ginadmin/internal/cli/db"
	files "github/gphper/ginadmin/internal/cli/file"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ginadmin-cli"}
	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(files.CmdFile)
	rootCmd.Execute()

}

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-06-01 20:15:04
 */
package main

import (
	"github.com/gphper/ginadmin/cmd/cli/db"
	"github.com/gphper/ginadmin/cmd/cli/file"
	"github.com/gphper/ginadmin/cmd/cli/run"
	"github.com/spf13/cobra"
)

var (
	release bool = true
)

// @title GinAdmin Api
// @version 1.0
// @description GinAdmin 示例项目

// @contact.name gphper
// @contact.url https://github.com/gphper/ginadmin

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:20011
// @basepath /api
func main() {

	var rootCmd = &cobra.Command{Use: "ginadmin"}
	rootCmd.AddCommand(run.CmdRun, db.CmdDb, file.CmdFile)
	rootCmd.Execute()

}

/**
 * @Author: GPHPER
 * @Date: 2022-12-14 08:50:00
 * @Github: https://github.com/gphper
 * @LastEditTime: 2022-12-14 11:25:56
 * @FilePath: \ginadmin\cmd\cli\file\root.go
 * @Description:
 */
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
	CmdFile.AddCommand(cmdModel, cmdController)
}

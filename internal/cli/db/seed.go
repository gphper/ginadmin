/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-04 15:11:11
 */
package db

import (
	"fmt"
	"github/gphper/ginadmin/internal/models"
	"github/gphper/ginadmin/pkg/comment"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cmdSeed = &cobra.Command{
	Use:   "seed [-t table]",
	Short: "DB Seed",
	Run:   seedFunc,
}

var tableSeed string

func init() {
	cmdSeed.Flags().StringVarP(&tableSeed, "table", "t", "", "input a table name")
}

func seedFunc(cmd *cobra.Command, args []string) {

	var err error
	if len(args) != 0 && args[0] == "real" {
		err = realExecSeed()
	} else {
		err = execSelfSeed()
	}

	if err != nil {
		fmt.Printf("migrate database fail:%s", err.Error())
	}
}

func realExecSeed() error {
	modelss := models.GetModels()
	if len(table) == 0 {
		for _, v := range modelss {
			tabler := v.(models.GaTabler)
			tabler.FillData()
		}
	} else {
		var tableExit bool = false
		for _, v := range modelss {
			tabler := v.(models.GaTabler)
			if tabler.TableName() == tableSeed {
				tableExit = true
				tabler.FillData()
			}
		}
		if !tableExit {
			fmt.Println("data table information does not exist")
		}
	}
	return nil
}

func execSelfSeed() error {
	var out []byte
	var err error
	rootPath, _ := comment.RootPath()
	cmd := exec.Command("go", "run", rootPath+"\\cmd\\ginadmin-cli\\ginadmin-cli.go", "db", "seed", "real")
	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	_ = out
	return err
}

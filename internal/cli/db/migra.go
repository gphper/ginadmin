package db

import (
	"errors"
	"fmt"
	"ginadmin/internal/models"
	"ginadmin/pkg/comment"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cmdMigrate = &cobra.Command{
	Use:   "migrate [-t table]",
	Short: "DB Migrate",
	Run:   migrateFunc,
}

var table string

func init() {
	cmdMigrate.Flags().StringVarP(&table, "table", "t", "", "input a table name")
}

func migrateFunc(cmd *cobra.Command, args []string) {
	var err error
	if len(args) != 0 && args[0] == "real" {
		err = realExec()
	} else {
		err = execSelf()
	}

	if err != nil {
		fmt.Println("migrate database fail:", err.Error())
	}
}

func realExec() error {
	var err error
	modelss := models.GetModels()
	if len(table) == 0 {
		err = models.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(modelss...)
	} else {
		var tableExit bool = false
		for _, v := range modelss {
			tabler := v.(models.GaTabler)
			if tabler.TableName() == table {
				tableExit = true
				err = models.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(v)
			}
		}
		if !tableExit {
			err = errors.New("data table information does not exist")
		}
	}
	return err
}

func execSelf() error {
	var out []byte
	var err error
	rootPath, _ := comment.RootPath()
	cmd := exec.Command("go", "run", rootPath+"\\cli\\cmd\\ginadmin-cli.go", "db", "migrate", "real")
	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	_ = out
	return err
}

package db

import (
	"fmt"
	"ginadmin/models"

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
	modelss := models.GetModels()
	var err error
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
			fmt.Println("data table information does not exist")
		}
	}

	if err != nil {
		fmt.Println("migrate database fail:", err.Error())
	}
}

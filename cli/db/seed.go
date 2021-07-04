package db

import (
	"fmt"
	"ginadmin/models"

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
	modelss := models.GetModels()
	var err error
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

	if err != nil {
		fmt.Println("migrate database fail:", err.Error())
	}
}

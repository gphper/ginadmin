package db

import (
	"strings"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/internal/redis"

	"github.com/spf13/cobra"
)

var cmdSeed = &cobra.Command{
	Use:   "seed [-t table]",
	Short: "DB Seed",
	Run:   seedFunc,
}

var tableSeed string
var confPath string

func init() {
	cmdSeed.Flags().StringVarP(&confPath, "config path", "c", "", "config path")
	cmdSeed.Flags().StringVarP(&tableSeed, "table", "t", "", "input a table name")
}

func seedFunc(cmd *cobra.Command, args []string) {
	var tableMap map[string]struct{}

	configs.Init(confPath)
	redis.Init()
	models.Init()

	tableMap = make(map[string]struct{})
	if tableSeed != "" {
		tablesSlice := strings.Split(tableSeed, ",")
		for _, v := range tablesSlice {
			tableMap[v] = struct{}{}
		}
	}

	for _, v := range models.GetModels() {

		if tableSeed != "" {
			if _, ok := tableMap[v.(models.GaTabler).TableName()]; !ok {
				continue
			}
		}

		tabler := v.(models.GaTabler)
		db := models.GetDB(tabler)
		tabler.FillData(db)
	}

}

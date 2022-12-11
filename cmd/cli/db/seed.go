package db

import (
	"log"
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
	var err error

	err = configs.Init(configPath)
	if err != nil {
		log.Fatalf("start fail:[Config Init] %s", err.Error())
	}

	err = redis.Init()
	if err != nil {
		log.Fatalf("start fail:[Redis Init] %s", err.Error())
	}

	err = models.Init()
	if err != nil {
		log.Fatalf("start fail:[Mysql Init] %s", err.Error())
	}

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

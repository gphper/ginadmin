package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gphper/ginadmin/configs"
	"github.com/gphper/ginadmin/internal/models"
	"github.com/gphper/ginadmin/internal/redis"

	"github.com/spf13/cobra"
)

var cmdMigrate = &cobra.Command{
	Use:   "migrate [-t table]",
	Short: "DB Migrate",
	Run:   migrateFunc,
}

var tables string
var configPath string

func init() {
	cmdMigrate.Flags().StringVarP(&configPath, "config path", "c", "", "config path")
	cmdMigrate.Flags().StringVarP(&tables, "table", "t", "", "input a table name")
}

func migrateFunc(cmd *cobra.Command, args []string) {

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
	if tables != "" {
		tablesSlice := strings.Split(tables, ",")
		for _, v := range tablesSlice {
			fmt.Println(v)
			tableMap[v] = struct{}{}
		}

	}

	for _, v := range models.GetModels() {
		db := models.GetDB(v.(models.GaTabler))
		if tables != "" {
			if _, ok := tableMap[v.(models.GaTabler).TableName()]; !ok {
				continue
			}
		}

		err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(v)
		if err != nil {
			fmt.Println("migrate database fail:", err.Error())
			os.Exit(0)
		}
	}
}

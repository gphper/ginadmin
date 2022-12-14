// go run .\cmd\ginadmin\ file model -m shop_type
package file

import (
	"fmt"
	"log"

	"github.com/gphper/ginadmin/configs"
	gstrings "github.com/gphper/ginadmin/pkg/utils/strings"
	"github.com/spf13/cobra"
)

var cmdModel = &cobra.Command{
	Use:   "model [-m modelName]",
	Short: "create model",
	Run:   modelFunc,
}

var modelName string

func init() {
	cmdModel.Flags().StringVarP(&modelName, "model", "m", "", "input model name eg: shop_items")
}

func modelFunc(cmd *cobra.Command, args []string) {

	var err error

	err = configs.Init("")
	if err != nil {
		log.Fatalf("start fail:[Config Init] %s", err.Error())
	}

	if len(modelName) == 0 {
		cmd.Help()
		return
	}

	fileName, firstName, secondUpper := gstrings.StrFirstToUpper(modelName)
	// 添加mode文件
	err = writeModel(fileName, firstName)
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		return
	}

	//修改default文件
	modifyDefault(fileName)

	//添加dao文件
	writeDao(fileName, secondUpper)
}

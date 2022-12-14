/**
 * @Author: GPHPER
 * @Date: 2022-12-14 10:07:22
 * @Github: https://github.com/gphper
 * @LastEditTime: 2022-12-14 10:10:49
 * @FilePath: \ginadmin\cmd\cli\file\model_gen_default.go
 * @Description:
 */
package file

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/gphper/ginadmin/configs"
)

func modifyDefault(fileName string) {

	var filePath = configs.RootPath + "internal" + string(filepath.Separator) + "models" + string(filepath.Separator) + "default.go"

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	ast.Walk(&Visitor{
		fset: fset,
		name: fileName,
	}, f)

	//写回到文件中
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0766)
	if err != nil {
		log.Fatalf("open err %s", err.Error())
	}
	err = format.Node(file, fset, f)
	if err != nil {
		log.Fatal(err)
	}
}

type Visitor struct {
	fset *token.FileSet
	name string
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	//判断ast分类
	case *ast.FuncDecl:
		demo := node.(*ast.FuncDecl)
		if demo.Name.Name == "GetModels" {

			returnStm, ok := demo.Body.List[0].(*ast.ReturnStmt)
			if !ok {
				return v
			}

			comp, ok := returnStm.Results[0].(*ast.CompositeLit)
			if !ok {
				return v
			}

			comp.Elts = append(comp.Elts, &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.CompositeLit{
					Type: &ast.Ident{
						Name: v.name,
					},
				},
			})

		}

	}

	return v
}

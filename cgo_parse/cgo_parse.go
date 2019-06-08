package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/loader"
)

func main() {

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/Users/yongfuchen/go/src/go_test/cgo/use_c.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fset, file)
	if err == nil {
		fmt.Println(buf.String())
	}

	// ===============================

	conf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			FakeImportC: true,
		},
	}
	conf.Import("go_test/cgo")
	prog, err := conf.Load()
	if err != nil {
		return
	}

	for _, pkgInfo := range prog.InitialPackages() {
		for _, file := range pkgInfo.Files {

			posn := conf.Fset.Position(file.Pos())

			fmt.Println(posn.Filename)

			fileIngo, err := parser.ParseFile(conf.Fset, posn.Filename, nil, parser.ParseComments)
			if fileIngo != nil {
				var buf bytes.Buffer
				err = format.Node(&buf, fset, fileIngo)
				if err == nil {
					fmt.Println(buf.String())
				}
			}

			var buf bytes.Buffer
			err = format.Node(&buf, fset, file)
			if err == nil {
				fmt.Println(buf.String())
			}
		}
	}
}

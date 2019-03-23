package aspect

import (
	"fmt"
	"github.com/tenntenn/goq"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"testing"
)

const src1 = `package main
import "github.com/go-redis/redis"

type R struct {
	client1 redis.Client
}

`

func run1() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src1, 0)
	if err != nil {
		return err
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}

	files := []*ast.File{f}
	if _, err := config.Check("main", fset, files, info); err != nil {
		return err
	}

	results := goq.New(fset, files, info).Query(&goq.Type{})
	for _, r := range results {
		fmt.Println(r.Object, "at", fset.Position(r.Node.Pos()))
	}

	return nil
}

func TestAstQuery1(t *testing.T) {
	if err := run1(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

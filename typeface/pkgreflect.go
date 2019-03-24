/*
Problem: Go reflection does not support enumerating types, variables and functions of packages.

pkgreflect generates a file named pkgreflect.go in every parsed package directory.
This file contains the following maps of exported names to reflection types/values:

	var Types = map[string]reflect.Type{ ... }
	var Functions = map[string]reflect.Value{ ... }
	var Variables = map[string]reflect.Value{ ... }

Command line usage:

	pkgreflect --help
	pkgreflect [-notypes][-nofuncs][-novars][-unexported][-norecurs][-gofile=filename.go] [DIR_NAME]

If -norecurs is not set, then pkgreflect traverses recursively into sub-directories.
If no DIR_NAME is given, then the current directory is used as root.
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Jakegogo/aspectgo/compiler/util/osadapter"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/loader"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	notypes    bool
	nofuncs    bool
	novars     bool
	noconsts   bool
	unexported bool
	norecurs   bool
	stdout     bool
	gofile     string
	notests    bool
)

func main() {
	flag.BoolVar(&notypes, "notypes", false, "Don't list package types")
	flag.BoolVar(&nofuncs, "nofuncs", false, "Don't list package functions")
	flag.BoolVar(&novars, "novars", false, "Don't list package variables")
	flag.BoolVar(&noconsts, "noconsts", false, "Don't list package consts")
	flag.BoolVar(&unexported, "unexported", false, "Also list unexported names")
	flag.BoolVar(&norecurs, "norecurs", false, "Don't parse sub-directories resursively")
	flag.StringVar(&gofile, "gofile", "type_registry.go", "Name of the generated .go file")
	flag.BoolVar(&stdout, "stdout", false, "Write to stdout.")
	flag.BoolVar(&notests, "notests", false, "Don't list test related code")
	flag.Parse()

	var allPkgs = new([]pkgSet)
	var dir, _ = os.Getwd()
	if len(flag.Args()) > 0 {
		for _, dir := range flag.Args() {
			parseDir(dir, allPkgs)
		}
	} else {
		dir = "."
		parseDir(dir, allPkgs)
	}

	var buf *bytes.Buffer = &bytes.Buffer{}

	printAll(allPkgs, buf)

	if stdout {
		io.Copy(os.Stdout, buf)
	} else {
		filename := filepath.Join(dir, gofile)
		newFileData := buf.Bytes()
		oldFileData, _ := ioutil.ReadFile(filename)
		if !bytes.Equal(newFileData, oldFileData) {
			err := ioutil.WriteFile(filename, newFileData, 0660)
			if err != nil {
				panic(err)
			}
		}
	}
}

type pkgSet struct {
	pkg      *ast.Package
	fSet     *token.FileSet
	dir      *string
	pkgPath  *string
	pkgAlias *string
	skip     bool
}

func parseDir(dir string, allPkgs *[]pkgSet) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Printf("parse dir of %s error: %s", dir, err)
		}
	}()

	dirFile, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()
	info, err := dirFile.Stat()
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		panic("Path is not a directory: " + dir)
	}

	var fileSet = token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, dir, filter, 0)

	if err == nil {
		for _, pkg := range pkgs {
			*allPkgs = append(*allPkgs, pkgSet{pkg: pkg, fSet: fileSet, dir: &dir})
		}
	} else {
		fmt.Printf("parse package %s error: %v", dir, err)
	}

	if !norecurs {
		dirs, err := dirFile.Readdir(-1)
		if err != nil {
			panic(err)
		}
		for _, info := range dirs {
			if info.IsDir() {
				parseDir(filepath.Join(dir, info.Name()), allPkgs)
			}
		}
	}
}

func printAll(allPkgs *[]pkgSet, buf *bytes.Buffer) {
	goPath := os.Getenv("GOPATH")
	goPath = FindAppropriateGoPath("", goPath)
	pwd, _ := os.Getwd()
	sourceDir := filepath.Join(goPath, "src")

	if !strings.HasPrefix(pwd, sourceDir) {
		panic(fmt.Errorf("current dir not in GOPATH"))
		return
	}

	baseDir := strings.Replace(pwd, sourceDir, "", -1)
	baseDir = strings.TrimLeft(baseDir, (string)(filepath.Separator))

	fmt.Fprintln(buf, "// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.\n")
	fmt.Fprintln(buf, "package", strings.Replace(baseDir, ".", "_", -1))
	fmt.Fprintln(buf, "")
	fmt.Fprintln(buf, "import (\n\t\"reflect\"")

	printImport(allPkgs, buf, baseDir)

	fmt.Fprintln(buf, "\n)")
	// Types
	if !notypes {
		fmt.Fprintln(buf, "var Types = map[string]reflect.Type{")
		for _, pkgset := range *allPkgs {
			print(&pkgset, buf, ast.Typ, "\t\"%s\": reflect.TypeOf((*%s)(nil)).Elem(),\n")
		}
		fmt.Fprintln(buf, "}")
		fmt.Fprintln(buf, "")
	}

	// Functions
	if !nofuncs {
		fmt.Fprintln(buf, "var Functions = map[string]reflect.Value{")
		for _, pkgset := range *allPkgs {
			print(&pkgset, buf, ast.Fun, "\t\"%s\": reflect.ValueOf(%s),\n")
		}
		fmt.Fprintln(buf, "}")
		fmt.Fprintln(buf, "")
	}

	if !novars {
		// Addresses of variables
		fmt.Fprintln(buf, "var Variables = map[string]reflect.Value{")
		for _, pkgset := range *allPkgs {
			print(&pkgset, buf, ast.Var, "\t\"%s\": reflect.ValueOf(&%s),\n")
		}
		fmt.Fprintln(buf, "}")
		fmt.Fprintln(buf, "")
	}

	if !noconsts {
		// Addresses of consts
		fmt.Fprintln(buf, "var Consts = map[string]reflect.Value{")
		for _, pkgset := range *allPkgs {
			print(&pkgset, buf, ast.Con, "\t\"%s\": reflect.ValueOf(%s),\n")
		}
		fmt.Fprintln(buf, "}")
		fmt.Fprintln(buf, "")
	}
}

func printImport(allPkgs *[]pkgSet, buf *bytes.Buffer, baseDir string) {
	for i := 0; i < len(*allPkgs); i++ {
		resolveImport(&((*allPkgs)[i]), buf, baseDir)
	}
	slicePkg := *allPkgs
	sort.Slice(slicePkg, func(i, j int) bool {
		p1 := slicePkg[i].pkgPath
		p2 := slicePkg[j].pkgPath
		return *p1 < *p2
	})

	for i := 0; i < len(*allPkgs); i++ {
		pkgset := (*allPkgs)[i]
		if pkgset.skip {
			continue
		}
		fmt.Fprintf(buf, "\n\t%s \"%s\"", *pkgset.pkgAlias, *pkgset.pkgPath)
	}
}

func resolveImport(pkgset *pkgSet, buffer *bytes.Buffer, baseDir string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			emptyPkg := ""
			(*pkgset).skip = true
			(*pkgset).pkgPath = &emptyPkg
			fmt.Printf("resolve import of %s error: %s", baseDir, err)
		}
	}()

	var pkgPath string
	_, prog, _, _ := loadTarget(filepath.Join(baseDir, *pkgset.dir))

	initialPkgs := prog.InitialPackages()
	if len(initialPkgs) != 1 {
		fmt.Printf("unexpected initial packages: %v", initialPkgs)
	}
	pkgInfo := initialPkgs[0]
	if len(pkgInfo.Errors) != 0 {
		fmt.Printf("package %s has errors: %v %s", pkgInfo, pkgInfo.Errors)
	}
	if pkgInfo.Pkg.Name() == "main" {
		pkgset.skip = true
	}

	pkgPath = pkgInfo.Pkg.Path()

	pkgset.pkgPath = &pkgPath
	if strings.Contains(pkgPath, "/") {
		aliasPrefix := strings.Replace(strings.Replace(pkgPath, "/", "_", -1), ".", "_", -1)
		pkgAlias := pkgPath[strings.LastIndex(pkgPath, "/")+1:] + aliasPrefix + "_alias"
		pkgAlias = strings.Replace(pkgAlias, "-", "_", -1)
		pkgset.pkgAlias = &pkgAlias
	} else {
		pkgAlias := pkgPath + "_alias"
		pkgset.pkgAlias = &pkgAlias
	}
}

func print(pkgset *pkgSet, w io.Writer, kind ast.ObjKind, format string) {
	if pkgset.skip {
		return
	}

	pkg := pkgset.pkg
	fileSet := pkgset.fSet

	names := []struct{ pkg, alias string }{}
	for _, f := range pkg.Files {
		posn := fileSet.Position(f.Name.NamePos)
		// ignore _test.go file
		if strings.HasSuffix(posn.Filename, "_test.go") {
			continue
		}

		for name, object := range f.Scope.Objects {
			if object.Kind == kind && (unexported || ast.IsExported(name)) {
				names = append(names, struct{ pkg, alias string }{
					*pkgset.pkgPath + "." + name,
					*pkgset.pkgAlias + "." + name})
			}
		}
	}
	sort.Slice(names, func(i, j int) bool {
		return names[i].alias > names[j].alias
	})
	for _, name := range names {
		fmt.Fprintf(w, format, name.pkg, name.alias)
	}
}

func filter(info os.FileInfo) bool {

	name := info.Name()

	if info.IsDir() {
		return false
	}

	if name == gofile {
		return false
	}

	if filepath.Ext(name) != ".go" {
		return false
	}

	if strings.HasSuffix(name, "_test.go") && notests {
		return false
	}

	return true

}

func FindAppropriateGoPath(s string, oldGOPATH string) string {
	var curPath string
	var err error
	if curPath, err = os.Getwd(); err != nil {
		fmt.Printf("get current path error %s", err)
		return oldGOPATH
	}

	var backupPath = ""
	for _, path := range strings.Split(oldGOPATH, osadapter.GOPATH_SPLITTER) {
		path = strings.TrimRight(path, osadapter.DIR_SPLITTER)
		if (path == curPath || strings.HasPrefix(curPath, path)) && (s == "" || strings.HasPrefix(s, path)) {
			return path
		}
		// get the longest path
		if s == "" || (strings.HasPrefix(s, path)) && (len(path) < len(backupPath) || backupPath == "") {
			backupPath = path
		}
	}
	return backupPath
}

func loadTarget(target string) (*loader.Config, *loader.Program, string, error) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Printf("load target of %s error: %s", target, err)
			panic(err)
		}
	}()

	conf := loader.Config{
		ParserMode: parser.PackageClauseOnly,
		TypeChecker: types.Config{
			IgnoreFuncBodies:         true,
			DisableUnusedImportCheck: true,
		},
		AllowErrors: true,
	}

	var filename string

	if strings.HasSuffix(target, ".go") {
		path := filepath.Dir(target)
		conf.Import(path)
		filename = target
	} else {
		conf.Import(target)
	}

	prog, err := conf.Load()
	if err != nil {
		if strings.Contains(err.Error(), "no initial packages were loaded") {
			fmt.Printf("please check the target: %s ", target)
		}
		return nil, nil, "", err
	}

	return &conf, prog, filename, nil
}
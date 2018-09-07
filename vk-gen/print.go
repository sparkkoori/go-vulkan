package main

import (
	"fmt"
	goast "go/ast"
	"go/printer"
	"go/token"
	"os"
	"path"

	cast "github.com/elliotchance/c2go/ast"
)

func print(goAst goast.Node, cAst cast.Node) {
	// 	src := `
	// package main
	// func main() {
	// 	println("Hello, World!")
	// }
	// `
	//
	// 	fset := token.NewFileSet() // positions are relative to fset
	// 	goAst, err := parser.ParseFile(fset, "", src, 0)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	printGo(goAst)
}

func printGo(goAst goast.Node) {
	fset := token.NewFileSet() // positions are relative to fset
	cfg := printer.Config{Tabwidth: 1}
	cfg.Mode |= printer.RawFormat

	f, err := os.Create(path.Join(pkgdir(), "vk_auto.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := cfg.Fprint(f, fset, goAst); err != nil {
		fmt.Errorf("print: %s", err)
	}
	// goast.Print(fset, goAst)
}

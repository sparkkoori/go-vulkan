package main

import (
	"fmt"
	goast "go/ast"
	"go/printer"
	"go/token"
	"os"
	"path"
)

func print(target Target) {
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
	printGo(target.targetGo)
}

func printGo(nodes []goast.Node) {
	fset := token.NewFileSet() // positions are relative to fset
	cfg := printer.Config{Tabwidth: 1}
	cfg.Mode |= printer.RawFormat

	f, err := os.Create(path.Join(pkgdir(), "vk_auto.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file := &goast.File{}
	file.Name = &goast.Ident{Name: "vk"}
	for _, node := range nodes {
		switch n := node.(type) {
		case goast.Decl:
			file.Decls = append(file.Decls, n)
		default:
			// goast.Print(fset, node)
			panic(fmt.Sprintf("unkown node %#v", node))
		}
	}

	if err := cfg.Fprint(f, fset, file); err != nil {
		fmt.Errorf("print: %s", err)
	}
}

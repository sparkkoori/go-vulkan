package main

import (
	"fmt"
	goast "go/ast"
	"go/format"
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

	f, err := os.Create(path.Join(pkgdir(), "vk.auto.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file := &goast.File{}
	file.Name = &goast.Ident{Name: "vk"}

	file.Decls = []goast.Decl{
		&goast.GenDecl{
			Doc: &goast.CommentGroup{
				List: []*goast.Comment{
					&goast.Comment{
						Text: "//#include \"vulkan/vulkan.h\"",
					},
				},
			},
			Tok: token.IMPORT,
			Specs: []goast.Spec{
				&goast.ImportSpec{
					// Name: &goast.Ident{Name: ""},
					Path: &goast.BasicLit{token.NoPos, token.STRING, "\"C\""},
				},
			},
		},
	}

	for _, node := range nodes {
		if n, ok := node.(goast.Decl); ok {
			file.Decls = append(file.Decls, n)
		} else {
			panic(fmt.Sprintf("Node isn't ast.Decl:\n %#v\n", node))
		}
	}

	if err := format.Node(f, fset, file); err != nil {
		fmt.Errorf("format.Node() return error: %s", err)
	}

	// if err := cfg.Fprint(f, fset, file); err != nil {
	// 	fmt.Errorf("print: %s", err)
	// }
}

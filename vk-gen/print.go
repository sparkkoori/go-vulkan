package main

import (
	"bufio"
	"fmt"
	goast "go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	cast "github.com/elliotchance/c2go/ast"
)

func print(target Target) {
	printC(target.targetC)
	printGo(target.targetGo)
}

func printC(nodes []cast.Node) {
	cf, err := os.Create(path.Join(pkgdir(), "bridges.auto.c"))
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	hf, err := os.Create(path.Join(pkgdir(), "bridges.auto.h"))
	if err != nil {
		panic(err)
	}
	defer hf.Close()

	cw := bufio.NewWriter(cf)
	writec := func(str string) {
		_, err = cw.WriteString(str)
		if err != nil {
			cw.Flush()
			panic(err)
		}
	}

	writeh := func(str string) {
		_, err = hf.WriteString(str)
		if err != nil {
			hf.Sync()
			panic(err)
		}
	}

	writech := func(str string) {
		writec(str)
		writeh(str)
	}

	writec("#include \"bridges.auto.h\"\n")
	writec("\n")

	writeh("#pragma once\n")
	writeh("#include \"vulkan/vulkan.h\"\n")
	writeh("\n")

	for _, node := range nodes {
		decl := node.(*cast.TypedefDecl)
		vars := splitFuncPointerProtoType(decl.Type)
		writech(vars[0])
		writech("call")
		writech(decl.Name)
		writech("(")
		writech(decl.Name)
		writech(" ")
		writech("f")
		for i, arg := range vars[1:] {
			writech(", ")
			writech(arg)
			writech(" ")
			writech("arg")
			writech(strconv.Itoa(i))
		}
		writech(")")
		writeh(";\n")
		writec("\n{\n")
		writec("  return f(")
		for i, _ := range vars[1:] {
			if i != 0 {
				writec(", ")
			}
			writec("arg")
			writec(strconv.Itoa(i))
		}
		writec(");\n")
		writec("};\n")
		writech("\n")
	}

	cw.Flush()
	hf.Sync()
}

func splitFuncPointerProtoType(p string) []string {
	vars := []string{}

	rg := regexp.MustCompile("^(.*)\\s?\\(\\s?\\*\\s?\\)\\s?\\((.*)\\)$")
	subs := rg.FindStringSubmatch(p)

	vars = append(vars, subs[1])
	for _, param := range strings.Split(subs[2], ",") {
		if param == "void" {
			break
		}
		vars = append(vars, param)
	}
	return vars
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
					&goast.Comment{
						Text: "//#include \"bridges.auto.h\"",
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

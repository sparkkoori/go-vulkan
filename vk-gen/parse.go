package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/elliotchance/c2go/ast"
	"github.com/elliotchance/c2go/preprocessor"
)

func parse() Source {
	/*
		Because the ast dump produced by clang lacks of type ast tree in ParmVarDecl
		and FieldDecl.
		So a tick must be used.

		1. Parse source file to get all type strings in ParmVarDecls and FieldDecls.
		2. Append new typedefs to source file for every encoded type strings.
		3. Parse source file agian
		4. Fix all ParmVarDecls and FieldDecls use type strings to node mapping.
	*/

	ppvk := ""
	var unit *ast.TranslationUnitDecl
	typeStrs := make(map[string]bool, 128)
	encodedTypeStrs := make(map[string]ast.Node, 128)
	//step 1
	{
		ppvk = makePreprocessedVulkan()
		unit = parseSourceFileToAst(ppvk)
		for _, node := range unit.ChildNodes {
			switch n := node.(type) {
			case *ast.FunctionDecl:
				for _, cnode := range n.ChildNodes {
					cn, ok := cnode.(*ast.ParmVarDecl)
					if ok {
						typeStrs[cn.Type] = true
					}
				}
			case *ast.RecordDecl:
				for _, cnode := range n.ChildNodes {
					cn, ok := cnode.(*ast.FieldDecl)
					if ok {
						typeStrs[cn.Type] = true
					}
				}
			}
		}
	}

	//step 2
	{
		f, err := os.OpenFile(ppvk, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		for str, _ := range typeStrs {
			ident := encodeCTypeString(str)
			declStr := "typedef " + getCVarDeclString(ident, str) + ";\n"
			f.WriteString(declStr)
		}
		f.Sync()
	}

	//step 3
	{
		unit = parseSourceFileToAst(ppvk)
		for i := len(unit.ChildNodes) - 1; i >= 0; i-- {
			switch n := unit.ChildNodes[i].(type) {
			case *ast.TypedefDecl:
				if strings.HasPrefix(n.Name, "XXX") {
					encodedTypeStrs[n.Name] = n.ChildNodes[0]
					unit.ChildNodes = append(unit.ChildNodes[:i], unit.ChildNodes[i+1:]...)
				}
			}
		}
	}

	//step 4
	{
		for _, node := range unit.ChildNodes {
			switch n := node.(type) {
			case *ast.FunctionDecl:
				for _, cnode := range n.ChildNodes {
					cn, ok := cnode.(*ast.ParmVarDecl)
					if ok {
						et := encodeCTypeString(cn.Type)
						cn.AddChild(encodedTypeStrs[et])
						println(cn.Type, reflect.TypeOf(encodedTypeStrs[et]).String())
					}
				}
			case *ast.RecordDecl:
				for _, cnode := range n.ChildNodes {
					cn, ok := cnode.(*ast.FieldDecl)
					if ok {
						et := encodeCTypeString(cn.Type)
						cn.AddChild(encodedTypeStrs[et])
						println(cn.Type, reflect.TypeOf(encodedTypeStrs[et]).String())
					}
				}
			}
		}
	}

	// Debug:
	// for _, node := range unit.ChildNodes {
	// 	if _, ok := node.(*ast.RecordDecl); ok {
	// 		deepPrint(node, 0)
	// 	}
	// }
	return Source(unit.ChildNodes)
}

func encodeCTypeString(str string) string {
	str = strings.Replace(str, " ", "SPACE", -1)
	str = strings.Replace(str, "*", "STAR", -1)
	str = strings.Replace(str, "(", "LPAREN", -1)
	str = strings.Replace(str, ")", "RPAREN", -1)
	str = strings.Replace(str, "[", "LBRACK", -1)
	str = strings.Replace(str, "]", "RBRACK", -1)
	return "XXX" + str
}

func getVulkanPath() string {
	headerPath := path.Join(pkgdir(), "vulkan", "vulkan.h")
	_, err := os.Stat(headerPath)
	if err != nil {
		log.Fatalf("%s isn't found \n", headerPath)
	}
	return headerPath
}

func cpreprocess(srcPath string) []byte {
	flags := []string{}
	pp, _, _, err := preprocessor.Analyze([]string{srcPath}, flags)
	if err != nil {
		panic(err)
	}
	return pp
}

func makePreprocessedVulkan() string {
	p := getVulkanPath()
	code := cpreprocess(p)
	ppPath := path.Join(tempDir, "vulkan.pp.h")
	err := ioutil.WriteFile(ppPath, code, 0644)
	if err != nil {
		log.Fatalf("writing to %s failed: %v\n", ppPath, err)
	}
	return ppPath
}

func parseSourceFileToAst(srcPath string) *ast.TranslationUnitDecl {
	dump := execClangAstDump(srcPath)
	unit := parseAstDump(dump)
	return unit
}

func findCTypeIdentIndex(typeStr string) int {
	// Using the syntax rules, we know that:
	// to the right of all the "pointer to" derived type tokens
	// to the left of all "array of" derived type tokens
	// to the left of all "function returning" derived type tokens
	// inside all the grouping parentheses

	var baseIdx int
	// inside all the grouping parentheses
	{
		re := regexp.MustCompile("\\(\\s?\\*(.*?)\\)")
		for {
			subidx := re.FindStringSubmatchIndex(typeStr)
			if subidx == nil {
				break
			}
			typeStr = typeStr[subidx[2]:subidx[3]]
			baseIdx = subidx[2]
		}
	}

	var i int
	if i = strings.Index(typeStr, "["); i >= 0 {
		// i = i
	} else if i = strings.LastIndex(typeStr, "*"); i >= 0 {
		i = i + 1
	} else {
		i = len(typeStr)
	}
	return i + baseIdx
}

func getCVarDeclString(varIdent string, typeStr string) string {
	idx := findCTypeIdentIndex(typeStr)
	if idx > len(typeStr) {
		return typeStr[0:idx] + " " + varIdent
	} else {
		return typeStr[0:idx] + " " + varIdent + " " + typeStr[idx:]
	}
}

func execClangAstDump(p string) string {
	ast, err := exec.Command("clang", "-Xclang", "-ast-dump",
		"-fsyntax-only", "-fno-color-diagnostics", p).Output()
	if err != nil {
		// If clang fails it still prints out the AST, so we have to run it
		// again to get the real error.
		errBody, _ := exec.Command("clang", p).CombinedOutput()
		panic("clang failed: " + err.Error() + ":\n\n" + string(errBody))
	}
	return string(ast)
}

func parseAstDump(dump string) *ast.TranslationUnitDecl {
	lines := strings.Split(string(dump), "\n")

	var nodes []ast.Node
	stack := make([]ast.Node, 0)
	for _, line := range lines {
		trimmed := strings.TrimLeft(line, "|\\- `")
		if trimmed == "" {
			continue
		}
		level := (len(line) - len(trimmed)) / 2
		node := ast.Parse(trimmed)

		//find the parent and update the stack
		var parent ast.Node = nil
		for i := len(stack) - 1; i >= 0; i-- {
			if level > i {
				parent = stack[i]
				break
			} else {
				stack = stack[:i]
			}
		}
		stack = append(stack, node)

		if parent == nil {
			nodes = append(nodes, node)
		} else {
			parent.AddChild(node)
		}
	}

	if len(nodes) > 1 {
		panic("the top level node should be the only one")
	}

	unit := nodes[0].(*ast.TranslationUnitDecl)
	return unit
}

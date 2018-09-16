package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/elliotchance/c2go/ast"
	"github.com/elliotchance/c2go/preprocessor"
)

func parse() Source {
	var headerPath string
	{
		headerPath = path.Join(pkgdir(), "vulkan", "vulkan.h")
		_, err := os.Stat(headerPath)
		if err != nil {
			log.Fatalf("%s isn't found \n", headerPath)
		}
	}

	var proprossed []byte
	{
		flags := []string{}
		pp, _, _, err := preprocessor.Analyze([]string{headerPath}, flags)
		if err != nil {
			panic(err)
		}
		proprossed = pp
	}

	var ppPath string
	{
		ppPath = path.Join(tempDir, "vulkan.pp.h")
		err := ioutil.WriteFile(ppPath, proprossed, 0644)
		if err != nil {
			log.Fatalf("writing to %s failed: %v\n", ppPath, err)
		}
	}

	dump := execClangAstDump(ppPath)
	unit := parseAstDump(dump)
	nodes := unit.Children()
	fixTypeMiss(nodes)
	return Source(nodes)
}

func fixTypeMiss(nodes []ast.Node) {
	for _, node := range nodes {
		switch n := node.(type) {
		case *ast.FunctionDecl:
			for _, cnode := range n.ChildNodes {
				cn, ok := cnode.(*ast.ParmVarDecl)
				if ok {
					t := parseTypeString(cn.Type)
					n.AddChild(t)
				}
			}
		case *ast.RecordDecl:
			for _, cnode := range n.ChildNodes {
				cn, ok := cnode.(*ast.FieldDecl)
				if ok {
					t := parseTypeString(cn.Type)
					n.AddChild(t)
				}
			}
		}
	}
}

func findCTypeIdentIndex(typeStr string) int {
	var i int
	if i = strings.Index(typeStr, "["); i >= 0 {
		return i
	} else if i = strings.LastIndex(typeStr, "*"); i >= 0 {
		return i + 1
	} else {
		return len(typeStr)
	}
}

func getVarString(varIdent string, typeStr string) string {
	idx := findCTypeIdentIndex(typeStr)
	if idx > len(typeStr) {
		return typeStr[0:idx] + " " + varIdent
	} else {
		return typeStr[0:idx] + " " + varIdent + " " + typeStr[idx:]
	}
}

func parseTypeString(typeStr string) ast.Node {
	const ident = "XXX"
	declStr := "typedef " + getVarString(ident, typeStr) + ";"
	var tmp string
	{
		tmp = path.Join(tempDir, "typedef-dump.h")
		err := ioutil.WriteFile(tmp, []byte(declStr), 0644)
		if err != nil {
			log.Fatalf("writing to %s failed: %v\n", tmp, err)
		}
	}
	dump := execClangAstDump(tmp)
	unit := parseAstDump(dump)

	var def *ast.TypedefDecl
	for _, child := range unit.Children() {
		_def, ok := child.(*ast.TypedefDecl)
		if ok && _def.Name == ident {
			def = _def
			break
		}
	}
	return def.ChildNodes[0]
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

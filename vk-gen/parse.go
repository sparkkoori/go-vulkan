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
			log.Fatalln("%s isn't found")
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
		// dir, err := ioutil.TempDir("", "vk-gen")
		// if err != nil {
		// 	log.Fatalf("Cannot create temp folder: %v\n", err)
		// }
		// defer os.RemoveAll(dir)
		ppPath = path.Join("", "vulkan.auto.h")
		err := ioutil.WriteFile(ppPath, proprossed, 0644)
		if err != nil {
			log.Fatalf("writing to %s failed: %v\n", ppPath, err)
		}
	}

	ast, err := exec.Command("clang", "-Xclang", "-ast-dump",
		"-fsyntax-only", "-fno-color-diagnostics", ppPath).Output()
	if err != nil {
		// If clang fails it still prints out the AST, so we have to run it
		// again to get the real error.
		errBody, _ := exec.Command("clang", ppPath).CombinedOutput()
		panic("clang failed: " + err.Error() + ":\n\n" + string(errBody))
	}

	return parseClangAstDump(string(ast))
}

func parseClangAstDump(dump string) Source {
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
	src := Source(unit.Children())
	return src
}

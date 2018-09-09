package main

import (
	"fmt"

	cast "github.com/elliotchance/c2go/ast"
)

func generate(src Source) (tgt Target) {
	// mapping := make(map[string]string)

L:
	for _, node := range src {
		switch n := node.(type) {
		// case *cast.TypedefDecl:
		// 	return n.Name
		// case *cast.RecordDecl:
		// 	return n.Name
		// case *cast.EnumDecl:
		// 	return n.Name
		// case *cast.FunctionDecl:
		// 	return n.Name
		default:
			fmt.Printf("Ignore node\n")
			deepPrint(n, 0)
			break L
		}
	}

	return tgt
}

func deepPrint(node cast.Node, level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" - ")
	}
	fmt.Printf("%#v\n", node)

	for _, child := range node.Children() {
		deepPrint(child, level+1)
	}
}

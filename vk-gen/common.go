package main

import (
	"flag"
	"fmt"
	goast "go/ast"
	"reflect"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

type Source []cast.Node

type Target struct {
	targetGo []goast.Node
	targetC  []cast.Node
}

func (t *Target) addGo(n goast.Node) {
	t.targetGo = append(t.targetGo, n)
}

func (t *Target) addC(n cast.Node) {
	t.targetC = append(t.targetC, n)
}

func pkgdir() string {
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	return dir
}

func deepPrint(node cast.Node, level int) {
	if node == nil {
		return
	}
	header := ""
	for i := 0; i < level; i++ {
		header += " - "
	}
	header += reflect.TypeOf(node).String()
	fmt.Println(aurora.Blue(header))
	fmt.Printf("%#v\n", node)

	for _, child := range node.Children() {
		deepPrint(child, level+1)
	}
}

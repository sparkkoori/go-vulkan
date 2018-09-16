package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"

	goast "go/ast"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

var tempDir string

func init() {
	log.SetFlags(0)

	dir := path.Join(".", ".vk-gen-temp")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
	tempDir = dir
}

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

func main() {
	flag.Parse()

	src := parse()
	fmt.Printf("Parsed %d nodes.\n", aurora.Green(len(src)))
	tgt := generate(src)
	fmt.Printf("Generated %d go nodes and %d c nodes.\n", aurora.Green(len(tgt.targetGo)), aurora.Green(len(tgt.targetC)))
	print(tgt)
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

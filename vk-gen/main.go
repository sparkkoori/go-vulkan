package main

import (
	"flag"
	"log"

	goast "go/ast"

	cast "github.com/elliotchance/c2go/ast"
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

func main() {
	flag.Parse()

	src := parse()
	log.Printf("Parsed %d nodes\n", len(src))
	tgt := generate(src)
	log.Printf("Generated %d go nodes and %d c nodes\n", len(tgt.targetGo), len(tgt.targetC))
	print(tgt)
}

func pkgdir() string {
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	return dir
}

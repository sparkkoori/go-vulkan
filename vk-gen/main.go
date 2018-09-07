package main

import (
	"flag"
	goast "go/ast"

	cast "github.com/elliotchance/c2go/ast"
)

func main() {
	flag.Parse()

	header := parse()
	goAst, cAst := generate(header)
	print(goAst, cAst)
}

func print(goAst goast.Node, cAst cast.Node) {
}

func generate(cast.Node) (goast.Node, cast.Node) {
	return nil, nil
}

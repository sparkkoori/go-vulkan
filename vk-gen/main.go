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

func generate(cast.Node) (goast.Node, cast.Node) {
	return nil, nil
}

func pkgdir() string {
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	return dir
}

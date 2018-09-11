package main

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"reflect"
	"strings"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

func generate(src Source) (tgt Target) {
	ws := &workspace{}
	ws.init(src)

	for _, node := range src {
		name := getNodeName(node)
		if name == "" {
			continue
		}
		if !strings.HasPrefix(name, "VK") &&
			!strings.HasPrefix(name, "vk") &&
			!strings.HasPrefix(name, "Vk") {
			continue
		}
		ws.gen(name)
	}
	return ws.target
}

type workspace struct {
	target  Target
	source  map[string]cast.Node
	mapping map[string]string

	invalid bool
}

func (ws *workspace) init(nodes []cast.Node) {
	ws.mapping = map[string]string{
		"char":                   "byte",    //C.char
		"singed char":            "int8",    //C.schar
		"unsigned char":          "uint8",   //C.uchar
		"short":                  "int16",   //C.short
		"unsigned short":         "uint16",  //C.ushort
		"int":                    "int32",   //C.int
		"unsigned int":           "uint32",  //C.uint
		"long":                   "int32",   //C.long
		"unsigned long":          "uint32",  //C.ulong
		"long long int":          "int64",   //C.longlong
		"unsigned long long int": "uint64",  //C.ulonglong
		"float":                  "float32", //	C.float
		"double":                 "float64", //C.double

		"size_t":   "uint",
		"int8_t":   "int8",
		"uint8_t":  "uint8",
		"int16_t":  "int16",
		"uint16_t": "uint16",
		"int32_t":  "int32",
		"uint32_t": "uint32",
		"int64_t":  "int64",
		"uint64_t": "uint64",
	}

	ws.source = make(map[string]cast.Node, 2048)
	for _, node := range nodes {
		name := getNodeName(node)
		if name == "" {
			continue
		}
		if _node, ok := ws.source[name]; ok {
			fmt.Printf("Conflict node (original):\n")
			deepPrint(_node, 0)
			fmt.Printf("Conflict node (new):\n")
			deepPrint(node, 0)
			panic("Name conflict for " + name)
		}
		ws.source[name] = node
	}
}

func (ws *workspace) gen(name string) string {
	if ws.invalid {
		return ""
	}
	if goname, ok := ws.mapping[name]; ok {
		return goname //already generated
	}

	node, ok := ws.source[name]
	if !ok {
		ws.mapping[name] = ""
		return ""
	}

	var goname string
	switch n := node.(type) {
	case *cast.TypedefDecl:
		goname = ws.genTypedefDecl(n, name)
	// case *cast.RecordDecl:
	// 	return "struct " + n.Name
	// case *cast.EnumDecl:
	// 	return "enum " + n.Name
	// case *cast.FunctionDecl:
	// 	return n.Name
	default:
		println("Unkown node for gen():")
		deepPrint(node, 0)
		ws.invalid = true
		return ""
	}

	ws.mapping[name] = goname
	return goname
}

func (ws *workspace) genTypedefDecl(node *cast.TypedefDecl, cn string) string {
	ct := node.Type
	gt := ws.gen(ct)
	if gt == "" {
		gt = "C." + cn
	}
	gn := vkName(cn)

	ws.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: &goast.Ident{token.NoPos, gn, nil},
				Type: &goast.Ident{token.NoPos, gt, nil},
			},
		},
	})
	return gn
}

func getNodeName(node cast.Node) string {
	switch n := node.(type) {
	case *cast.TypedefDecl:
		return n.Name
	case *cast.RecordDecl:
		return "struct " + n.Name
	case *cast.EnumDecl:
		return "enum " + n.Name
	case *cast.FunctionDecl:
		return n.Name
	}
	deepPrint(node, 0)
	panic("Unkown node for getNodeName()")
	return ""
}

func vkName(n string) string {
	n = strings.TrimPrefix(n, "C.")
	n = strings.TrimPrefix(n, "Vk")
	n = strings.TrimPrefix(n, "VK_")
	n = strings.TrimPrefix(n, "vk")
	return n
}

func deepPrint(node cast.Node, level int) {
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

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
		// if name == "enum VkResult" {
		// 	ws.halt("", node)
		// }
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
	case *cast.EnumDecl:
		goname = ws.genEnumDecl(n, name)
	// case *cast.FunctionDecl:
	// 	return n.Name
	default:
		ws.halt("Unkown node for gen()", node)
	}

	ws.mapping[name] = goname
	return goname
}

func (ws *workspace) genTypedefDecl(node *cast.TypedefDecl, cn string) string {
	ct := node.Type
	gt := ws.gen(ct)

	if gt != "" {
		if strings.HasPrefix(ct, "enum ") ||
			strings.HasPrefix(ct, "struct ") {
			return gt //two-one mapping
		}
	}

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

func (ws *workspace) genEnumDecl(node *cast.EnumDecl, cn string) string {
	gn := strings.TrimPrefix(cn, "enum ")
	gn = vkName(gn)
	ws.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: &goast.Ident{token.NoPos, gn, nil},
				Type: &goast.Ident{token.NoPos, "int", nil},
			},
		},
	})

	specs := []goast.Spec{}
	constMapping := map[string]string{}
	for _, child := range node.ChildNodes {
		if ws.invalid {
			continue
		}
		con := child.(*cast.EnumConstantDecl)

		var expr = convExpr(constMapping, con.ChildNodes[0])

		// ws.halt("Unkown enum const", v)
		if expr != nil {
			gcon := vkName(con.Name)
			specs = append(specs, &goast.ValueSpec{
				Names: []*goast.Ident{
					&goast.Ident{Name: gcon},
				},
				Type:   &goast.Ident{Name: gn},
				Values: []goast.Expr{expr},
			})
			constMapping[con.Name] = gcon
		}
	}

	ws.target.addGo(&goast.GenDecl{
		Tok:    token.CONST,
		Lparen: token.Pos(1),
		Specs:  specs,
		Rparen: token.Pos(1),
	})

	return gn
}

func (ws *workspace) halt(str string, node cast.Node) {
	if !ws.invalid {
		ws.invalid = true
		fmt.Println(aurora.Red("Generating Halted:"), str)
		deepPrint(node, 0)
	}
}

func convExpr(refs map[string]string, node cast.Node) goast.Expr {
	switch v := node.(type) {
	case *cast.IntegerLiteral:
		return &goast.BasicLit{Kind: token.INT, Value: v.Value}
	case *cast.DeclRefExpr:
		return &goast.Ident{Name: refs[v.Name]}
	case *cast.ParenExpr:
		return &goast.ParenExpr{
			Lparen: token.Pos(1),
			X:      convExpr(refs, v.ChildNodes[0]),
			Rparen: token.Pos(1),
		}
	case *cast.BinaryOperator:
		return &goast.BinaryExpr{
			X:  convExpr(refs, v.ChildNodes[0]),
			Op: convToken(v.Operator),
			Y:  convExpr(refs, v.ChildNodes[1]),
		}
	case *cast.UnaryOperator:
		return &goast.UnaryExpr{
			X:  convExpr(refs, v.ChildNodes[0]),
			Op: convToken(v.Operator),
		}
	default:
		return nil
	}
}

func convToken(op string) token.Token {
	switch op {
	case token.ADD.String():
		return token.ADD
	case token.SUB.String():
		return token.SUB
	case token.MUL.String():
		return token.MUL
	case token.QUO.String():
		return token.QUO
	case token.REM.String():
		return token.REM

	case token.AND.String():
		return token.AND
	case token.OR.String():
		return token.OR
	case token.XOR.String():
		return token.XOR
	case token.SHL.String():
		return token.SHL
	case token.SHR.String():
		return token.SHR
	case token.AND_NOT.String():
		return token.AND_NOT

	default:
		return token.ILLEGAL
	}
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
	panic("Unkown node name for " + reflect.TypeOf(node).String())
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

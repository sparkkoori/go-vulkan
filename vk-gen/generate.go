package main

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

type halting struct {
	msg  string
	node cast.Node
}

func halt(msg string, node cast.Node) {
	panic(&halting{msg, node})
}

func generate(src Source) (tgt Target) {
	ws := &workspace{}

	defer func() {
		if r := recover(); r != nil {
			if h, ok := r.(*halting); ok {
				fmt.Println(aurora.Red("Generating Halted:"), h.msg)
				deepPrint(h.node, 0)
				tgt = ws.target
			} else {
				panic(r)
			}
		}
	}()

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

//mapping is c-to-go id mapping.
type mapping struct {
	types  map[string]string
	consts map[string]string
	funcs  map[string]string

	typeGenFn  func(string)
	constGenFn func(string)
	funcGenFn  func(string)
}

func (m *mapping) init() {
	m.types = map[string]string{
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

		"void *": "unsafe.Pointer",

		"size_t":   "uint",
		"int8_t":   "int8",
		"uint8_t":  "uint8",
		"int16_t":  "int16",
		"uint16_t": "uint16",
		"int32_t":  "int32",
		"uint32_t": "uint32",
		"int64_t":  "int64",
		"uint64_t": "uint64",

		"VkBool32": "bool",

		"char *": "string",
	}

	m.consts = map[string]string{}
	m.funcs = map[string]string{}
}

func (m *mapping) get(k string) (string, bool) {
	v, ok := m.types[k]
	if ok {
		return v, ok
	}
	v, ok = m.consts[k]
	if ok {
		return v, ok
	}
	v, ok = m.funcs[k]
	if ok {
		return v, ok
	}
	return "", false
}

func (m *mapping) setType(cv, gv string) {
	if _, ok := m.types[cv]; ok {
		panic("Type mapping conflicts for " + cv)
	}
	m.types[cv] = gv
}

func (m *mapping) getType(t string) string {
	gt, ok := m.types[t]
	if !ok {
		m.typeGenFn(t)
		gt, _ = m.types[t]
		return gt
	}
	return gt
}

func (m *mapping) mustGetType(t string) string {
	gt := m.getType(t)
	if gt == "" {
		panic("No mapping for " + t)
	}
	return gt
}

func (m *mapping) setConst(cv, gv string) {
	if _, ok := m.consts[cv]; ok {
		halt("Type mapping conflicts for "+cv, nil)
	}
	m.consts[cv] = gv
}

func (m *mapping) getConst(c string) string {
	gc, ok := m.consts[c]
	if !ok {
		m.constGenFn(c)
		gc, _ = m.consts[c]
		return gc
	}
	return gc
}

func (m *mapping) setFunc(cv, gv string) {
	if _, ok := m.funcs[cv]; ok {
		halt("Type mapping conflicts for "+cv, nil)
	}
	m.funcs[cv] = gv
}

func (m *mapping) getFunc(f string) string {
	gf, ok := m.consts[f]
	if !ok {
		m.funcGenFn(f)
		gf, _ = m.consts[f]
		return gf
	}
	return gf
}

func convComplexType(m *mapping, cn cast.Node) goast.Expr {
	switch n := cn.(type) {
	case *cast.Typedef:
		if n.Type == "void" {
			return nil
		}
		return &goast.Ident{Name: m.mustGetType(n.Type)}
	case *cast.PointerType:
		if v, ok := n.ChildNodes[0].(*cast.Typedef); ok {
			if v.Type == "void" {
				return &goast.Ident{Name: m.mustGetType("void *")}
			} else if v.Type == "char" {
				return &goast.Ident{Name: m.mustGetType("char *")}
			}
		}
		return &goast.StarExpr{
			Star: token.Pos(1),
			X:    convComplexType(m, n.ChildNodes[0]),
		}
	case *cast.ConstantArrayType:
		return &goast.ArrayType{
			Lbrack: token.Pos(1),
			Len:    &goast.BasicLit{Kind: token.INT, Value: strconv.Itoa(n.Size)},
			Elt:    convComplexType(m, n.ChildNodes[0]),
		}
	case *cast.FunctionProtoType:
		params := []*goast.Field{}
		for _, pn := range n.ChildNodes[1:] {
			params = append(params, &goast.Field{
				Type: convComplexType(m, pn),
			})
		}
		results := []*goast.Field{}
		{
			result := convComplexType(m, n.ChildNodes[0])
			if result != nil {
				results = append(results, &goast.Field{
					Type: result,
				})
			}
		}
		return &goast.FuncType{
			Func:    token.Pos(1),
			Params:  &goast.FieldList{List: params},
			Results: &goast.FieldList{List: results},
		}
	}
	halt("Unkown c complex type node", cn)
	return nil
}

func parseComplexCType(t string, baseNode cast.Node) cast.Node {
	t = strings.TrimSpace(t)
	if t == "" {
		return baseNode
	}

	//ignore cv qualifier
	for {
		t = strings.TrimSpace(t)
		if strings.HasPrefix(t, "const") {
			t = strings.TrimPrefix(t, "const")
			continue
		} else if strings.HasPrefix(t, "violate") {
			t = strings.TrimPrefix(t, "violate")
			continue
		}
		break
	}

	submatch := func(reg string) []string {
		re := regexp.MustCompile(reg)
		subs := re.FindStringSubmatch(t)
		return subs
	}

	match := func(reg string) string {
		re := regexp.MustCompile(reg)
		str := re.FindString(t)
		return str
	}

	//do not support anonymouse struct, union, enum
	if anony := match("^(struct\\s|enum\\s|union\\s)\\{.*\\}"); anony != "" {
		halt("Can't parse anonymouse struct, union, enum", nil)
	}

	if id := match("^(struct\\s|enum\\s)?[0-9A-Za-z_]*"); id != "" {
		n := &cast.Typedef{Type: id}
		// baseNode must be nil
		t = strings.TrimPrefix(t, id)
		return parseComplexCType(t, n)
	}

	if id := match("^\\*"); id != "" {
		n := &cast.PointerType{}
		n.AddChild(baseNode)
		t = strings.TrimPrefix(t, id)
		return parseComplexCType(t, n)
	}

	if arr := submatch("\\[(\\d*)\\]$"); arr != nil {
		size, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		n := &cast.ConstantArrayType{Size: size}
		n.AddChild(baseNode)
		t = strings.TrimPrefix(t, arr[0])
		return parseComplexCType(t, n)
	}

	if fn := submatch("\\((.*)\\)\\((.*)\\)$"); fn != nil {
		params := strings.Split(fn[2], ",")
		n := &cast.FunctionProtoType{}
		n.AddChild(baseNode)
		for _, param := range params {
			n.AddChild(parseComplexCType(param, nil))
		}
		t = fn[1]
		return parseComplexCType(t, n)
	}

	halt("Unkown complex c type "+t, nil)
	return nil
}

type workspace struct {
	target  Target
	source  map[string]cast.Node
	mapping mapping
}

func (ws *workspace) init(nodes []cast.Node) {
	ws.mapping.init()
	ws.mapping.typeGenFn = ws.gen
	ws.mapping.constGenFn = ws.gen
	ws.mapping.funcGenFn = ws.gen

	ws.source = make(map[string]cast.Node, 2048)
	for _, node := range nodes {
		name := getNodeName(node)
		if name == "" {
			continue
		}
		if _, ok := ws.source[name]; ok {
			halt("Nodes conflicts", node)
		}
		ws.source[name] = node
	}
}

func (ws *workspace) transType(t string) goast.Expr {
	cn := parseComplexCType(t, nil)
	gn := convComplexType(&ws.mapping, cn)
	return gn
}

func (ws *workspace) gen(name string) {
	_, ok := ws.mapping.get(name)
	if ok {
		return //already generated
	}

	node, ok := ws.source[name]
	if !ok {
		return
	}

	switch n := node.(type) {
	case *cast.TypedefDecl:
		ws.genTypedefDecl(n)
	// case *cast.RecordDecl:
	// 	return "struct " + n.Name
	case *cast.RecordDecl:
		ws.genRecordDecl(n)
	case *cast.EnumDecl:
		ws.genEnumDecl(n)
	// case *cast.FunctionDecl:
	// 	return n.Name
	default:
		halt("gen()", node)
	}
	// println(name)
}

func (ws *workspace) genTypedefDecl(node *cast.TypedefDecl) {
	cname := node.Name
	goname := toGoName(cname)

	//Skip, if any struct or enum with same name has been generated.
	if _, ok := ws.source["struct "+cname]; ok {
		goname = ws.mapping.mustGetType("struct " + cname)
		ws.mapping.setType(cname, goname)
		return
	} else if _, ok := ws.source["enum "+cname]; ok {
		goname = ws.mapping.mustGetType("enum " + cname)
		ws.mapping.setType(cname, goname)
		return
	}

	var t goast.Expr
	//Is this type is an opaque type? Use C.xxx or yyy?
	re := regexp.MustCompile("^struct\\s[a-z|A-Z]*_T \\*$")
	if re.MatchString(node.Type) {
		t = &goast.Ident{token.NoPos, "C." + cname, nil}
	} else {
		t = ws.transType(node.Type)
	}

	ws.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: &goast.Ident{token.NoPos, goname, nil},
				Type: t,
			},
		},
	})
	ws.mapping.setType(cname, goname)
}

func (ws *workspace) genEnumDecl(node *cast.EnumDecl) {
	cname := "enum " + node.Name
	goname := toGoName(cname)

	//generate enum type
	//use xxxFlags if the name is not xxxFlagBits.
	if strings.Contains(cname, "FlagBits") {
		fcname := strings.Replace(cname, "FlagBits", "Flags", 1)
		fcname = strings.TrimPrefix(fcname, "enum ")
		goname = ws.mapping.mustGetType(fcname)
	} else {
		ws.target.addGo(&goast.GenDecl{
			Tok: token.TYPE,
			Specs: []goast.Spec{
				&goast.TypeSpec{
					Name: &goast.Ident{token.NoPos, goname, nil},
					Type: &goast.Ident{token.NoPos, "int", nil},
				},
			},
		})
	}
	ws.mapping.setType(cname, goname)

	//generate consts
	specs := []goast.Spec{}
	for _, child := range node.ChildNodes {
		con := child.(*cast.EnumConstantDecl)
		var expr = convConst(&ws.mapping, con.ChildNodes[0])

		if expr != nil {
			gcon := toGoName(con.Name)
			specs = append(specs, &goast.ValueSpec{
				Names: []*goast.Ident{
					&goast.Ident{Name: gcon},
				},
				Type:   &goast.Ident{Name: goname},
				Values: []goast.Expr{expr},
			})
			ws.mapping.setConst(con.Name, gcon)
		}
	}

	ws.target.addGo(&goast.GenDecl{
		Tok:    token.CONST,
		Lparen: token.Pos(1),
		Specs:  specs,
		Rparen: token.Pos(1),
	})
}

func (ws *workspace) genRecordDecl(node *cast.RecordDecl) {
	cname := "struct " + node.Name
	goname := toGoName(cname)

	ws.mapping.setType(cname, goname)

	fields := []*goast.Field{}
	for i := 0; i < len(node.ChildNodes); i++ {
		fnode := node.ChildNodes[i]
		fdecl, ok := fnode.(*cast.FieldDecl)
		if !ok {
			halt("Expected *cast.FieldDecl type", fnode)
		}

		fcn := fdecl.Name
		fct := fdecl.Type
		fgn := toGoName(fcn)
		fgt := ws.transType(fct)

		fields = append(fields, &goast.Field{
			Names: []*goast.Ident{&goast.Ident{Name: fgn}},
			Type:  fgt,
		})
	}

	ws.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: &goast.Ident{token.NoPos, goname, nil},
				Type: &goast.StructType{
					Fields: &goast.FieldList{
						List: fields,
					},
				},
			},
		},
	})

}

func convConst(m *mapping, node cast.Node) goast.Expr {
	switch v := node.(type) {
	case *cast.IntegerLiteral:
		return &goast.BasicLit{Kind: token.INT, Value: v.Value}
	case *cast.DeclRefExpr:
		return &goast.Ident{Name: m.getConst(v.Name)}
	case *cast.ParenExpr:
		return &goast.ParenExpr{
			Lparen: token.Pos(1),
			X:      convConst(m, v.ChildNodes[0]),
			Rparen: token.Pos(1),
		}
	case *cast.BinaryOperator:
		return &goast.BinaryExpr{
			X:  convConst(m, v.ChildNodes[0]),
			Op: parseOperator(v.Operator),
			Y:  convConst(m, v.ChildNodes[1]),
		}
	case *cast.UnaryOperator:
		return &goast.UnaryExpr{
			X:  convConst(m, v.ChildNodes[0]),
			Op: parseOperator(v.Operator),
		}
	default:
		halt("convConst()", node)
	}
	return nil
}

func parseOperator(op string) token.Token {
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

func toGoName(n string) string {
	n = strings.TrimPrefix(n, "enum ")
	n = strings.TrimPrefix(n, "struct ")

	n = strings.TrimPrefix(n, "Vk")
	n = strings.TrimPrefix(n, "VK_")
	n = strings.TrimPrefix(n, "vk")

	re := regexp.MustCompile("^(p+)[A-Z]+")
	if subs := re.FindStringSubmatch(n); subs != nil {
		n = strings.TrimPrefix(n, subs[1])
	}

	re = regexp.MustCompile("^s[A-Z]+")
	if re.MatchString(n) {
		n = strings.TrimPrefix(n, "s")
	}

	n = strings.Title(n)

	return n
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

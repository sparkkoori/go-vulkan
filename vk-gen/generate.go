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
	"github.com/iancoleman/strcase"
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

		"enum VkStructureType":      "",
		"struct VkBaseInStructure":  "",
		"struct VkBaseOutStructure": "",
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
		} else if _, ok := n.ChildNodes[0].(*cast.FunctionProtoType); ok {
			return &goast.Ident{Name: m.mustGetType("void *")}
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
		return &goast.Ident{Name: "interface{}"}
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
		goname = ws.mapping.getType("struct " + cname)
		ws.mapping.setType(cname, goname)
		return
	} else if _, ok := ws.source["enum "+cname]; ok {
		goname = ws.mapping.getType("enum " + cname)
		ws.mapping.setType(cname, goname)
		return
	}

	match := func(str, r string) bool {
		re := regexp.MustCompile(r)
		return re.MatchString(str)
	}

	var t goast.Expr
	//Is this type is an opaque type? Use C.xxx or yyy?
	if match(node.Type, "^struct\\s[a-z|A-Z]*_T \\*$") { // vulkan handle
		t = &goast.Ident{Name: "C." + cname}
	} else if match(node.Type, "\\(\\*") { // c function pointer
		t = &goast.Ident{Name: "struct { raw C." + cname + " }"}
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

	if match(node.Type, "\\(\\*") {
		ws.genBridgeCallMethod(goname, node)
	}
}

func (ws *workspace) genBridgeCallMethod(goname string, node *cast.TypedefDecl) {
	pt := parseComplexCType(node.Type, nil).(*cast.PointerType)
	fpt := pt.ChildNodes[0].(*cast.FunctionProtoType)

	var paramFields []*goast.Field
	var resultFields []*goast.Field
	var args []goast.Expr
	for i, f := range fpt.ChildNodes {
		gf := convComplexType(&ws.mapping, f)
		if i == 0 {
			if gf != nil {
				resultFields = append(resultFields, &goast.Field{Type: gf})
			}
		} else {
			arg := "arg" + strconv.Itoa(i-1)
			args = append(args, &goast.Ident{Name: arg})
			paramFields = append(paramFields, &goast.Field{
				Names: []*goast.Ident{
					&goast.Ident{Name: arg},
				},
				Type: gf,
			})
		}
	}

	//print.go will generate bridge call directly form these typedefs.
	ws.target.addC(node)

	warpReturnStmtIf := func(expr goast.Expr, cond bool) goast.Stmt {
		rs := &goast.ReturnStmt{
			Return: token.Pos(1),
		}

		if cond {
			rs.Results = append(rs.Results, expr)
			return rs
		} else {
			return &goast.ExprStmt{
				X: expr,
			}
		}
	}

	ws.target.addGo(&goast.FuncDecl{
		Recv: &goast.FieldList{
			Opening: token.Pos(1),
			List: []*goast.Field{
				&goast.Field{
					Names: []*goast.Ident{
						&goast.Ident{Name: "p"},
					},
					Type: &goast.Ident{Name: goname},
				},
			},
			Closing: token.Pos(1),
		},
		Name: &goast.Ident{Name: "Call"},
		Type: &goast.FuncType{
			Func: token.Pos(1),
			Params: &goast.FieldList{
				Opening: token.Pos(1),
				List:    paramFields,
				Closing: token.Pos(1),
			},
			Results: &goast.FieldList{
				Opening: token.Pos(1),
				List:    resultFields,
				Closing: token.Pos(1),
			},
		},
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List: []goast.Stmt{
				warpReturnStmtIf(&goast.CallExpr{
					Fun:    &goast.Ident{Name: "C.call" + node.Name},
					Lparen: token.Pos(1),
					Args: append([]goast.Expr{
						&goast.CallExpr{
							Fun:    &goast.Ident{Name: "C." + node.Name},
							Lparen: token.Pos(1),
							Args: []goast.Expr{
								&goast.Ident{Name: "p.raw"},
							},
							Rparen: token.Pos(1),
						},
					}, args...),
					Rparen: token.Pos(1),
				}, len(resultFields) > 0),
			},
			Rbrace: token.Pos(1),
		},
	})
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
	var isStructure bool
	var xxxCount *cast.FieldDecl
	var pendAdding *goast.Field
	for i := 0; i < len(node.ChildNodes); i++ {
		fnode := node.ChildNodes[i]
		fdecl, ok := fnode.(*cast.FieldDecl)
		if !ok {
			halt("Expected *cast.FieldDecl type", fnode)
		}

		fcn := fdecl.Name
		if fcn == "sType" {
			isStructure = true
			continue //ignore sType
		}
		fct := fdecl.Type
		fgn := toGoName(fcn)
		fgt := ws.transType(fct)

		if fcn == "pNext" {
			fgt = &goast.Ident{Name: "Structure"}
		}

		if strings.HasSuffix(fcn, "Count") && fct == "uint32_t" {
			xxxCount = fdecl
			pendAdding = &goast.Field{
				Names: []*goast.Ident{&goast.Ident{Name: fgn}},
				Type:  fgt,
			}
		} else if xxxCount != nil &&
			strings.HasSuffix(fct, "*") {
			pendAdding = nil

			pointer := fgt.(*goast.StarExpr)
			slice := &goast.ArrayType{
				Elt: pointer.X,
			}

			fields = append(fields, &goast.Field{
				Names: []*goast.Ident{&goast.Ident{Name: fgn}},
				Type:  slice,
			})
		} else {
			if xxxCount != nil {
				xxxCount = nil
				if pendAdding != nil {
					fields = append(fields, pendAdding)
					pendAdding = nil
				}
			}
			fields = append(fields, &goast.Field{
				Names: []*goast.Ident{&goast.Ident{Name: fgn}},
				Type:  fgt,
			})
		}
	}
	if pendAdding != nil {
		fields = append(fields, pendAdding)
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

	if isStructure {
		ws.genStructureMethods(cname, goname)
	}
}

func (ws *workspace) genStructureMethods(cname, goname string) {
	/*VkStructureType value is obtained by taking the name of the structure,
	  stripping the leading Vk, prefixing each capital letter with _,
	  converting the entire resulting string to upper case,
	  and prefixing it with VK_STRUCTURE_TYPE_.*/
	sType := strings.TrimPrefix(cname, "struct ")
	sType = strings.TrimPrefix(sType, "Vk")
	sType = strcase.ToScreamingSnake(sType)
	sType = "VK_STRUCTURE_TYPE_" + sType

	//sType()
	ws.target.addGo(&goast.FuncDecl{
		Recv: &goast.FieldList{
			Opening: token.Pos(1),
			List: []*goast.Field{
				&goast.Field{
					Names: []*goast.Ident{
						&goast.Ident{Name: "s"},
					},
					Type: &goast.StarExpr{
						X: &goast.Ident{Name: goname},
					},
				},
			},
			Closing: token.Pos(1),
		},
		Name: &goast.Ident{Name: "sType"},
		Type: &goast.FuncType{
			Func: token.Pos(1),
			Params: &goast.FieldList{
				Opening: token.Pos(1),
				Closing: token.Pos(1),
			},
			Results: &goast.FieldList{
				List: []*goast.Field{
					&goast.Field{
						Type: &goast.Ident{Name: "C.VkStructureType"},
					},
				},
			},
		},
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List: []goast.Stmt{
				&goast.ReturnStmt{
					Return: token.Pos(1),
					Results: []goast.Expr{
						&goast.Ident{Name: "C." + sType},
					},
				},
			},
			Rbrace: token.Pos(1),
		},
	})

	//GetNext()
	ws.target.addGo(&goast.FuncDecl{
		Recv: &goast.FieldList{
			Opening: token.Pos(1),
			List: []*goast.Field{
				&goast.Field{
					Names: []*goast.Ident{
						&goast.Ident{Name: "s"},
					},
					Type: &goast.StarExpr{
						X: &goast.Ident{Name: goname},
					},
				},
			},
			Closing: token.Pos(1),
		},
		Name: &goast.Ident{Name: "GetNext"},
		Type: &goast.FuncType{
			Func: token.Pos(1),
			Params: &goast.FieldList{
				Opening: token.Pos(1),
				Closing: token.Pos(1),
			},
			Results: &goast.FieldList{
				List: []*goast.Field{
					&goast.Field{
						Type: &goast.Ident{Name: "Structure"},
					},
				},
			},
		},
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List: []goast.Stmt{
				&goast.ReturnStmt{
					Return: token.Pos(1),
					Results: []goast.Expr{
						&goast.Ident{Name: "s.Next"},
					},
				},
			},
			Rbrace: token.Pos(1),
		},
	})

	//SetNext()
	ws.target.addGo(&goast.FuncDecl{
		Recv: &goast.FieldList{
			Opening: token.Pos(1),
			List: []*goast.Field{
				&goast.Field{
					Names: []*goast.Ident{
						&goast.Ident{Name: "s"},
					},
					Type: &goast.StarExpr{
						X: &goast.Ident{Name: goname},
					},
				},
			},
			Closing: token.Pos(1),
		},
		Name: &goast.Ident{Name: "SetNext"},
		Type: &goast.FuncType{
			Func: token.Pos(1),
			Params: &goast.FieldList{
				Opening: token.Pos(1),
				List: []*goast.Field{
					&goast.Field{
						Names: []*goast.Ident{
							&goast.Ident{Name: "n"},
						},
						Type: &goast.Ident{Name: "Structure"},
					},
				},
				Closing: token.Pos(1),
			},
		},
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List: []goast.Stmt{
				&goast.AssignStmt{
					Lhs: []goast.Expr{
						&goast.Ident{Name: "s.Next"},
					},
					Tok: token.ASSIGN,
					Rhs: []goast.Expr{
						&goast.Ident{Name: "n"},
					},
				},
			},
			Rbrace: token.Pos(1),
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

	re = regexp.MustCompile("^pfn[A-Z]+")
	if re.MatchString(n) {
		n = strings.TrimPrefix(n, "pfn")
	}

	if strings.HasPrefix(n, "PFN_vk") {
		n = strings.Replace(n, "PFN_vk", "PFN", 1)
	}

	n = strings.Title(n)

	return n
}

package main

import (
	"fmt"
	goast "go/ast"
	"go/token"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

type typeInfo struct {
	gotype goast.Expr
	ctype  cast.Node
	c2go   func(govar, cvar goast.Expr) goast.Stmt
	go2c   func(govar, cvar goast.Expr) goast.Stmt
	isRef  bool
}

func (ti *typeInfo) cgoTypeExpr() goast.Expr {
	t := cgoType(ti.ctype)
	return &goast.Ident{Name: t}
}

func (ti *typeInfo) sizeofCgoTypeExpr() goast.Expr {
	st := sizeofCgoType(ti.ctype)
	return &goast.Ident{Name: st}
}

func saalloc(node cast.Node) *goast.CallExpr {
	fun := parenExpr(starExpr(ident(cgoType(node))))
	arg := callExpr(ident("_sa.alloc"), ident(sizeofCgoType(node)))
	call := callExpr(fun, arg)
	return call
}

func cgoTypeExpr(node cast.Node) goast.Expr {
	t := cgoType(node)
	return &goast.Ident{Name: t}
}

func sizeofCgoTypeExpr(node cast.Node) goast.Expr {
	st := sizeofCgoType(node)
	return &goast.Ident{Name: st}
}

func cgoType(node cast.Node) string {
	switch n := node.(type) {
	case *cast.BuiltinType:
		return "C." + cgobuiltin(n.Type)
	case *cast.PointerType:
		if n.Type == "void *" || n.Type == "void const *" {
			return "unsafe.Pointer"
		} else {
			t := cgoType(n.ChildNodes[0])
			if t == "" {
				return "*[0]byte"
			} else {
				return "*" + t
			}
		}
	case *cast.ConstantArrayType:
	case *cast.TypedefType:
		return "C." + n.Type
	case *cast.ParenType:
	default:
		halt("Unkown type for cgoType()", node)
	}
	return ""
}

func sizeofCgoType(node cast.Node) string {
	switch n := node.(type) {
	case *cast.BuiltinType:
		return "C.sizeof_" + cgobuiltin(n.Type)
	case *cast.PointerType:
		return "unsafe.Sizeof(*[0]byte)"
	case *cast.ConstantArrayType:
	case *cast.TypedefType:
		return "C.sizeof_" + n.Type
	case *cast.ParenType:
	default:
		halt("Unkown type for sizeofCgoType()", node)
	}
	return ""
}

func cgobuiltin(k string) string {
	var m = map[string]string{
		"char":                   "char",
		"singed char":            "schar",
		"unsigned char":          "uchar",
		"short":                  "short",
		"unsigned short":         "ushort",
		"int":                    "int",
		"unsigned int":           "uint",
		"long":                   "long",
		"unsigned long":          "ulong",
		"long long int":          "longlong",
		"unsigned long long int": "ulonglong",
		"float":                  "float",
		"double":                 "double",
	}
	return m[k]
}

type generator struct {
	target Target
	nodes  map[string]cast.Node //name node map
	types  map[string]*typeInfo
	consts map[string]bool
	funcs  map[string]bool
}

func (g *generator) init() {
	g.nodes = make(map[string]cast.Node, 2048)
	g.types = make(map[string]*typeInfo, 512)
	g.consts = make(map[string]bool, 128)
	g.funcs = make(map[string]bool, 128)
}

func (g *generator) genType(node cast.Node) *typeInfo {
	// deepPrint(node, 0)
	switch n := node.(type) {
	case *cast.BuiltinType:
		return g.genBuiltinType(n)
	case *cast.PointerType:
		return g.genPointerType(n)
	case *cast.ConstantArrayType:
		return g.genConstantArrayType(n)
	case *cast.TypedefType:
		return g.genTypedefType(n)
	case *cast.RecordType:
		return g.genRecordType(n)
	case *cast.EnumType:
		return g.genEnumType(n)
	case *cast.FunctionProtoType:
		return g.genFunctionProtoType(n)
	case *cast.QualType:
		return g.genQualType(n)
	case *cast.IncompleteArrayType:
		return nil
	default:
		return nil
	}
}

func (g *generator) genRecordType(n *cast.RecordType) *typeInfo {
	return nil
}

func (g *generator) genEnumType(n *cast.EnumType) *typeInfo {
	return nil
}

func (g *generator) genFunctionProtoType(n *cast.FunctionProtoType) *typeInfo {
	return nil
}

func (g *generator) genQualType(n *cast.QualType) *typeInfo {
	return nil
}

func (g *generator) genPointerType(n *cast.PointerType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}

	if n.Type == "void *" || n.Type == "void const *" {
		return &typeInfo{
			ctype:  n,
			gotype: &goast.Ident{Name: "unsafe.Pointer"},
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			},
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			},
		}
	}

	o := n.ChildNodes[0]
	info := g.genType(o)
	if info == nil {
		return &typeInfo{
			ctype:  n,
			gotype: starExpr(ident("[0]byte")),
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			},
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			},
		}
	} else {
		return &typeInfo{
			ctype:  n,
			gotype: starExpr(info.gotype),
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(govar, callExpr(ident("new"), info.gotype))
				conv := info.c2go(starExpr(govar), starExpr(cvar))
				return blockStmt(alloc, conv)
			},
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(cvar, saalloc(o))
				conv := info.go2c(starExpr(govar), starExpr(cvar))
				return blockStmt(alloc, conv)
			},
			isRef: true,
		}
	}
}

func (g *generator) genConstantArrayType(n *cast.ConstantArrayType) *typeInfo {
	deepPrint(n, 0)
	return nil
}

func (g *generator) genTypedefType(n *cast.TypedefType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}

	o := n.ChildNodes[1]
	oinfo := g.genType(o)
	checkTypeInfo(oinfo, o)

	gotype := ident(n.Type)
	cgotype := cgoTypeExpr(n)

	g.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: gotype,
				Type: oinfo.gotype,
			},
		},
	})

	info := &typeInfo{
		ctype:  n,
		gotype: gotype,
		c2go: func(govar, cvar goast.Expr) goast.Stmt {
			_temp := ident("_temp")
			def := varDeclStmt(oinfo.gotype, _temp)
			convc := oinfo.c2go(_temp, callExpr(parenExpr(oinfo.cgoTypeExpr()), cvar))
			convbase := assignStmt1n1(govar, callExpr(gotype, _temp))
			return blockStmt(def, convc, convbase)
		},

		go2c: func(govar, cvar goast.Expr) goast.Stmt {
			_temp := ident("_temp")
			def := varDeclStmt(oinfo.cgoTypeExpr(), _temp)
			convc := oinfo.go2c(callExpr(parenExpr(oinfo.gotype), govar), _temp)
			convbase := assignStmt1n1(cvar, callExpr(cgotype, _temp))
			return blockStmt(def, convc, convbase)
		},
	}

	g.types[n.Type] = info
	return info
}

func (g *generator) genBuiltinType(n *cast.BuiltinType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}

	var m = map[string]string{
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
		"void":                   "",
	}
	gotypestr := m[n.Type]
	if gotypestr == "" {
		return nil
	}

	gotype := ident(gotypestr)
	cgotype := cgoTypeExpr(n)

	return &typeInfo{
		ctype:  n,
		gotype: gotype,
		c2go: func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(gotype, cvar))
		},
		go2c: func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(cgotype, govar))
		},
	}
}

func (g *generator) genConst() {
	// TODO:genConst
}

func (g *generator) genFunc(fn *cast.FunctionDecl) {
	if g.funcs[fn.Name] {
		return
	}

	var params, results []*goast.Field
	var hasRefParam bool

	cargdefs := []goast.Stmt{}
	cargs := []goast.Expr{}
	go2cs := []goast.Stmt{}
	c2gos := []goast.Stmt{}
	//Params
	for _, param := range fn.ChildNodes[1:] {
		pvd, ok := param.(*cast.ParmVarDecl)
		if !ok {
			break
		}
		node := pvd.ChildNodes[0]
		info := g.genType(node)
		if info == nil {
			break //for void
		}

		goarg := ident(pvd.Name)
		carg := ident("c_" + goarg.Name)
		gotype := info.gotype
		ctype := info.cgoTypeExpr()

		cargdefs = append(cargdefs, varDeclStmt(ctype, carg))
		cargs = append(cargs, carg)
		params = append(params, field(gotype, goarg))
		go2cs = append(go2cs, info.go2c(goarg, carg))
		if info.isRef {
			hasRefParam = true
			oinfo := g.genType(node.(*cast.PointerType).ChildNodes[0])
			c2gos = append(c2gos, oinfo.c2go(starExpr(goarg), starExpr(carg)))
		}
	}

	//Call and Results
	var ccall goast.Stmt
	var retstmt goast.Stmt
	{
		rs := fn.ChildNodes[0]
		info := g.genType(rs)
		arg := ident("ret")
		carg := ident("c_ret")
		if info != nil {
			results = append(results, field(info.gotype, arg))
			ccall = assignStmt1n1D(carg, callExpr(ident("C."+fn.Name), cargs...))
			varDeclStmt(info.cgoTypeExpr(), carg)
			c2gos = append(c2gos, info.c2go(arg, carg))
			retstmt = &goast.ReturnStmt{}
		} else {
			ccall = exprStmt(callExpr(ident("C."+fn.Name), cargs...))
		}
	}

	//stackAllocator
	sastmts := []goast.Stmt{}
	if hasRefParam {
		sa := ident("_sa")
		sastmts = append(sastmts,
			assignStmt1n1D(sa, callExpr(ident("pool.take"))))
		sastmts = append(sastmts, &goast.DeferStmt{
			Defer: token.Pos(1),
			Call:  callExpr(ident("pool.give"), sa),
		})
	}

	var stmts []goast.Stmt
	{
		stmts = append(stmts, sastmts...)
		stmts = append(stmts, cargdefs...)
		stmts = append(stmts, go2cs...)
		stmts = append(stmts, ccall)
		stmts = append(stmts, c2gos...)
		if retstmt != nil {
			stmts = append(stmts, retstmt)
		}
	}

	{
		name := ident(fn.Name)
		typ := funcType(params, results)
		fndef := funcDecl(name, nil, typ, stmts...)
		g.target.addGo(fndef)
	}
}

func (g *generator) process(src Source) {
	for _, node := range src {
		switch n := node.(type) {
		case *cast.TypedefDecl:
		case *cast.RecordDecl:
		case *cast.EnumDecl:
		case *cast.FunctionDecl:
			g.genFunc(n)
		default:
			halt("Unkown node in source", node)
		}
	}
}

func checkTypeInfo(info *typeInfo, node cast.Node) {
	if info == nil {
		halt("Type info is nil", node)
	}
}

func ident(name string) *goast.Ident {
	return &goast.Ident{
		Name: name,
	}
}

func starExpr(expr goast.Expr) *goast.StarExpr {
	return &goast.StarExpr{
		X: expr,
	}
}

func parenExpr(expr goast.Expr) *goast.ParenExpr {
	return &goast.ParenExpr{
		X: expr,
	}
}

func exprStmt(expr goast.Expr) *goast.ExprStmt {
	return &goast.ExprStmt{
		X: expr,
	}
}

func callExpr(fn goast.Expr, args ...goast.Expr) *goast.CallExpr {
	return &goast.CallExpr{
		Fun:    fn,
		Lparen: token.Pos(1),
		Rparen: token.Pos(1),
		Args:   args,
	}
}

func blockStmt(stmts ...goast.Stmt) *goast.BlockStmt {
	return &goast.BlockStmt{
		Lbrace: token.Pos(1),
		Rbrace: token.Pos(1),
		List:   stmts,
	}
}

func assignStmt1n1(lh, rh goast.Expr) *goast.AssignStmt {
	return assignStmt([]goast.Expr{lh}, []goast.Expr{rh})
}

func assignStmt(lhs, rhs []goast.Expr) *goast.AssignStmt {
	return &goast.AssignStmt{
		Lhs: lhs,
		Rhs: rhs,
		Tok: token.ASSIGN,
	}
}

func assignStmt1n1D(lh, rh goast.Expr) *goast.AssignStmt {
	return assignStmtD([]goast.Expr{lh}, []goast.Expr{rh})
}

func assignStmtD(lhs, rhs []goast.Expr) *goast.AssignStmt {
	return &goast.AssignStmt{
		Lhs: lhs,
		Rhs: rhs,
		Tok: token.DEFINE,
	}
}

func varDeclStmt(typ goast.Expr, names ...*goast.Ident) *goast.DeclStmt {
	return &goast.DeclStmt{
		Decl: &goast.GenDecl{
			Tok: token.VAR,
			Specs: []goast.Spec{
				&goast.ValueSpec{
					Names: names,
					Type:  typ,
				},
			},
		},
	}
}

func field(typ goast.Expr, names ...*goast.Ident) *goast.Field {
	return &goast.Field{
		Names: names,
		Type:  typ,
	}
}

func funcType(params, results []*goast.Field) *goast.FuncType {
	return &goast.FuncType{
		Func: token.Pos(1),
		Params: &goast.FieldList{
			Opening: token.Pos(1),
			Closing: token.Pos(1),
			List:    params,
		},
		Results: &goast.FieldList{
			Opening: token.Pos(1),
			Closing: token.Pos(1),
			List:    results,
		},
	}
}

func funcDecl(name *goast.Ident, recv *goast.Field, typ *goast.FuncType, stmts ...goast.Stmt) *goast.FuncDecl {
	var list *goast.FieldList
	if recv != nil {
		list = &goast.FieldList{
			Opening: token.Pos(1),
			Closing: token.Pos(1),
			List: []*goast.Field{
				recv,
			},
		}
	}
	return &goast.FuncDecl{
		Name: name,
		Type: typ,
		Recv: list,
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List:   stmts,
			Rbrace: token.Pos(1),
		},
	}
}

func generate(src Source) (tgt Target) {
	g := &generator{}
	g.init()

	defer func() {
		if r := recover(); r != nil {
			if h, ok := r.(*halting); ok {
				fmt.Println(aurora.Red("Generating Halted:"), h.msg)
				deepPrint(h.node, 0)
				tgt = g.target
			} else {
				panic(r)
			}
		}
	}()

	g.process(src)

	return g.target
}

type halting struct {
	msg  string
	node cast.Node
}

func halt(msg string, node cast.Node) {
	panic(&halting{msg, node})
}

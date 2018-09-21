package main

import (
	"fmt"
	goast "go/ast"
	"go/token"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

type typeInfo struct {
	gotype    goast.Expr
	ctype     goast.Expr
	csize     goast.Expr
	go2cAlloc bool
	c2go      func(govar, cvar goast.Expr) goast.Stmt
	go2c      func(govar, cvar goast.Expr) goast.Stmt
	refc2go   func(govar, cvar goast.Expr) goast.Stmt
}

type compTypeInfo struct {
	gofields  []*goast.Field
	cfields   []*goast.Field
	go2cAlloc bool
	c2go      func(goscope, cscope goast.Expr) []goast.Stmt
	go2c      func(goscope, cscope goast.Expr) []goast.Stmt
	refc2go   func(goscope, cscope goast.Expr) []goast.Stmt
}

func saalloc(typ, size goast.Expr) *goast.CallExpr {
	fun := parenExpr(starExpr(typ))
	arg := callExpr(ident("_sa.alloc"), size)
	call := callExpr(fun, arg)
	return call
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
	case *cast.ElaboratedType:
		return g.genElaboratedType(n)
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
	if info, ok := g.types[n.Type]; ok {
		return info
	}
	var info *typeInfo

	decl := g.nodes[n.Type]
	if decl == nil {
		halt("No decl for record type", n)
	}

	g.types[n.Type] = info
	return info
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
			ctype:  ident("unsafe.Pointer"),
			gotype: ident("unsafe.Pointer"),
			csize:  ident("unsafe.Sizeof(uintptr(0))"),
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			},
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			},
		}
	}

	o := n.ChildNodes[0]
	oinfo := g.genType(o)
	if oinfo == nil {
		typ := starExpr(ident("[0]byte"))
		return &typeInfo{
			ctype:  typ,
			gotype: typ,
			csize:  callExpr(ident("unsafe.Sizeof"), typ),
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			},
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			},
		}
	} else {
		return &typeInfo{
			ctype:  starExpr(oinfo.ctype),
			gotype: starExpr(oinfo.gotype),
			csize:  callExpr(ident("unsafe.Sizeof"), starExpr(oinfo.ctype)),
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(govar, callExpr(ident("new"), oinfo.gotype))
				conv := oinfo.c2go(starExpr(govar), starExpr(cvar))
				return blockStmt(alloc, conv)
			},
			go2cAlloc: true,
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(cvar, saalloc(oinfo.ctype, oinfo.csize))
				conv := oinfo.go2c(starExpr(govar), starExpr(cvar))
				return blockStmt(alloc, conv)
			},
			refc2go: func(govar, cvar goast.Expr) goast.Stmt {
				return oinfo.c2go(starExpr(govar), starExpr(cvar))
			},
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
	ctype := ident("C." + n.Type)
	csize := ident("C.sizeof_" + n.Type)

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
		ctype:  ctype,
		gotype: gotype,
		csize:  csize,
		c2go: func(govar, cvar goast.Expr) goast.Stmt {
			_temp := ident("_temp")
			def := varDeclStmt(oinfo.gotype, _temp)
			convc := oinfo.c2go(_temp, callExpr(parenExpr(oinfo.ctype), cvar))
			convbase := assignStmt1n1(govar, callExpr(gotype, _temp))
			return blockStmt(def, convc, convbase)
		},

		go2c: func(govar, cvar goast.Expr) goast.Stmt {
			_temp := ident("_temp")
			def := varDeclStmt(oinfo.ctype, _temp)
			convc := oinfo.go2c(callExpr(parenExpr(oinfo.gotype), govar), _temp)
			convbase := assignStmt1n1(cvar, callExpr(ctype, _temp))
			return blockStmt(def, convc, convbase)
		},
		refc2go: oinfo.refc2go, //for pointer type case
	}

	g.types[n.Type] = info
	return info
}

func (g *generator) genElaboratedType(n *cast.ElaboratedType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}
	var info *typeInfo

	obj := n.ChildNodes[0]

	switch o := obj.(type) {
	case *cast.RecordType:
		info = g.genRecordType(o)
	case *cast.EnumType:
		info = g.genEnumType(o)
	default:
		halt("Unkown elaborated type", o)
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
	var m2 = map[string]string{
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
	gotypestr := m[n.Type]
	ctypestr := m2[n.Type]
	if gotypestr == "" || ctypestr == "" {
		return nil
	}

	gotype := ident(gotypestr)
	ctype := ident("C." + ctypestr)
	csize := ident("C.sizeof_" + ctypestr)

	return &typeInfo{
		ctype:  ctype,
		gotype: gotype,
		csize:  csize,
		c2go: func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(gotype, cvar))
		},
		go2c: func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(ctype, govar))
		},
	}
}

func (g *generator) genCompType(fieldDecls []*cast.FieldDecl) *compTypeInfo {
	info := &compTypeInfo{}

	go2cFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	c2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	refc2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}

	for _, fieldDecl := range fieldDecls {
		finfo := g.genType(fieldDecl.ChildNodes[0])
		cname := ident(fieldDecl.Name)
		goname := ident(fieldDecl.Name)
		info.cfields = append(info.cfields, field(finfo.ctype, cname))
		info.gofields = append(info.gofields, field(finfo.gotype, goname))
		go2cFns = append(go2cFns, func(goscope, cscope goast.Expr) goast.Stmt {
			return finfo.go2c(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
		})
		c2goFns = append(c2goFns, func(goscope, cscope goast.Expr) goast.Stmt {
			return finfo.c2go(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
		})
		if finfo.refc2go != nil {
			refc2goFns = append(refc2goFns, func(goscope, cscope goast.Expr) goast.Stmt {
				return finfo.refc2go(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
			})
		}
		if finfo.go2cAlloc {
			info.go2cAlloc = true
		}
	}

	info.c2go = func(goscope, cscope goast.Expr) []goast.Stmt {
		stmts := []goast.Stmt{}
		for _, fn := range c2goFns {
			stmts = append(stmts, fn(goscope, cscope))
		}
		return stmts
	}

	info.go2c = func(goscope, cscope goast.Expr) []goast.Stmt {
		stmts := []goast.Stmt{}
		for _, fn := range go2cFns {
			stmts = append(stmts, fn(goscope, cscope))
		}
		return stmts
	}

	info.refc2go = func(goscope, cscope goast.Expr) []goast.Stmt {
		stmts := []goast.Stmt{}
		for _, fn := range refc2goFns {
			stmts = append(stmts, fn(goscope, cscope))
		}
		return stmts
	}
	return info
}

func (g *generator) genConst() {
	// TODO:genConst
}

func (g *generator) genFunc(fn *cast.FunctionDecl) {
	if g.funcs[fn.Name] {
		return
	}

	var goparams, goresults []*goast.Field
	var cparams, cresults []*goast.Field

	var cdef goast.Stmt
	var allocdefs []goast.Stmt
	var pconvs []goast.Stmt
	var call goast.Stmt
	var rconvs []goast.Stmt
	var prconvs []goast.Stmt
	var retstmt goast.Stmt

	//Params
	{
		fieldDecls := []*cast.FieldDecl{}
		for _, param := range fn.ChildNodes[1:] {
			pvd, ok := param.(*cast.ParmVarDecl)
			if !ok {
				break
			}
			fieldDecl := &cast.FieldDecl{
				Name:       pvd.Name,
				ChildNodes: pvd.ChildNodes,
			}
			fieldDecls = append(fieldDecls, fieldDecl)
		}
		info := g.genCompType(fieldDecls)
		goparams = info.gofields
		cparams = info.cfields
		pconvs = info.go2c(nil, ident("c"))
		prconvs = info.refc2go(nil, ident("c"))

		if info.go2cAlloc {
			sa := ident("_sa")
			allocdefs = append(allocdefs,
				assignStmt1n1D(sa, callExpr(ident("pool.take"))))
			allocdefs = append(allocdefs, &goast.DeferStmt{
				Defer: token.Pos(1),
				Call:  callExpr(ident("pool.give"), sa),
			})
		}
	}

	//Results
	{
		info := g.genType(fn.ChildNodes[0])
		goname := ident("_ret")
		cname := ident("_ret")
		if info != nil {
			goresults = append(goresults, field(info.gotype, goname))
			cresults = append(cresults, field(info.ctype, cname))
			rconvs = append(rconvs, info.c2go(goname, selectorExpr(ident("c"), cname)))
		}
	}

	//Call
	{
		args := []goast.Expr{}
		for _, cparam := range cparams {
			args = append(args, selectorExpr(ident("c"), cparam.Names[0]))
		}
		cfn := ident("C." + fn.Name)
		if len(cresults) > 0 {
			rs := selectorExpr(ident("c"), cresults[0].Names[0])
			call = assignStmt1n1(rs, callExpr(cfn, args...))
			retstmt = &goast.ReturnStmt{}
		} else {
			call = exprStmt(callExpr(cfn, args...))
		}
	}

	//cdef
	{
		fields := []*goast.Field{}
		fields = append(fields, cparams...)
		fields = append(fields, cresults...)

		if len(fields) > 0 {
			cdef = &goast.DeclStmt{
				Decl: &goast.GenDecl{
					Tok: token.VAR,
					Specs: []goast.Spec{
						&goast.ValueSpec{
							Names: []*goast.Ident{
								ident("c"),
							},
							Type: &goast.StructType{
								Struct: token.Pos(1),
								Fields: &goast.FieldList{
									Opening: token.Pos(1),
									List:    fields,
									Closing: token.Pos(1),
								},
							},
						},
					},
				},
			}
		}
	}

	//Generate
	{
		var stmts []goast.Stmt
		if cdef != nil {
			stmts = append(stmts, cdef)
		}
		stmts = append(stmts, allocdefs...)
		stmts = append(stmts, pconvs...)
		stmts = append(stmts, call)
		stmts = append(stmts, rconvs...)
		stmts = append(stmts, prconvs...)
		if retstmt != nil {
			stmts = append(stmts, retstmt)
		}

		name := ident(fn.Name)
		gofntype := funcType(goparams, goresults)
		gofndef := funcDecl(name, nil, gofntype, stmts...)
		g.target.addGo(gofndef)
	}
}

func (g *generator) process(src Source) {
	for _, node := range src {
		switch n := node.(type) {
		case *cast.TypedefDecl:
			g.nodes[n.Type] = n
		case *cast.RecordDecl:
			g.nodes["struct "+n.Name] = n
		case *cast.EnumDecl:
			g.nodes["enum "+n.Name] = n
		case *cast.FunctionDecl:
			g.nodes[n.Name] = n
		default:
			halt("Unkown node in source", node)
		}
	}

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

func selectorExpr(x goast.Expr, name *goast.Ident) goast.Expr {
	if x != nil {
		return &goast.SelectorExpr{
			X:   x,
			Sel: name,
		}
	} else {
		return name
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

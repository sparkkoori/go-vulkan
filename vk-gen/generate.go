package main

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"strconv"
	"strings"

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

type generator struct {
	target Target
	nodes  map[string]cast.Node //name node map
	types  map[string]*typeInfo
	consts map[string]goast.Expr
	funcs  map[string]bool

	hints map[string][]fieldHint
}

func (g *generator) init() {
	g.nodes = make(map[string]cast.Node, 2048)
	g.types = make(map[string]*typeInfo, 512)
	g.consts = make(map[string]goast.Expr, 128)
	g.funcs = make(map[string]bool, 512)

	g.hints = make(map[string][]fieldHint, 512)
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
	var info *typeInfo = &typeInfo{}

	var recordDecl *cast.RecordDecl
	{
		node := g.nodes[n.Type]
		if node == nil {
			halt("No decl for record type", n)
		}
		recordDecl = node.(*cast.RecordDecl)
	}

	if !recordDecl.Definition {
		info.gotype = ident("C.struct_" + recordDecl.Name)
		info.ctype = ident("C.struct_" + recordDecl.Name)
		info.csize = nil
		g.types[n.Type] = info
		return info
	}

	var cinfo *compTypeInfo
	{
		fieldDecls := []*cast.FieldDecl{}
		for _, child := range recordDecl.ChildNodes {
			fieldDecls = append(fieldDecls, child.(*cast.FieldDecl))
		}
		cinfo = g.genCompType(fieldDecls, g.hints[n.Type])
	}

	if recordDecl.Name != "" {
		info.gotype = ident(recordDecl.Name)
		if _, ok := g.nodes[recordDecl.Name]; ok {
			//same name typedef exists
			info.ctype = ident("C." + recordDecl.Name)
			info.csize = ident("C.sizeof_" + recordDecl.Name)
		} else {
			info.ctype = ident("C.struct" + recordDecl.Name)
			info.csize = ident("C.sizeof_struct_" + recordDecl.Name)
		}

		g.target.addGo(&goast.GenDecl{
			Tok: token.TYPE,
			Specs: []goast.Spec{
				&goast.TypeSpec{
					Name: info.gotype.(*goast.Ident),
					Type: &goast.StructType{
						Struct: token.Pos(1),
						Fields: &goast.FieldList{
							Opening: token.Pos(1),
							List:    cinfo.gofields,
							Closing: token.Pos(1),
						},
					},
				},
			},
		})
	} else {
		halt("Unamed struct isn't implemented", n)
	}

	takeAddr := func(govar, cvar goast.Expr) (goast.Expr, goast.Expr) {
		if starX, ok := govar.(*goast.StarExpr); ok {
			govar = starX.X
		}
		if starX, ok := cvar.(*goast.StarExpr); ok {
			cvar = starX.X
		} else {
			cvar = &goast.UnaryExpr{Op: token.AND, X: cvar}
		}
		return govar, cvar
	}

	//go2c
	{
		recv := field(starExpr(info.gotype), ident("g"))
		var fntype *goast.FuncType
		if cinfo.go2cAlloc {
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				govar, cvar = takeAddr(govar, cvar)
				return exprStmt(callExpr(selectorExpr(govar, ident("toC")), cvar, ident("_sa")))
			}
			fntype = funcType([]*goast.Field{
				field(starExpr(info.ctype), ident("c")),
				field(starExpr(ident("stackAllocator")), ident("_sa")),
			}, nil)
		} else {
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				govar, cvar = takeAddr(govar, cvar)
				return exprStmt(callExpr(selectorExpr(govar, ident("toC")), cvar))
			}
			fntype = funcType([]*goast.Field{
				field(starExpr(info.ctype), ident("c")),
			}, nil)
		}
		stmts := cinfo.go2c(ident("g"), ident("c"))
		g.target.addGo(funcDecl(ident("toC"), recv, fntype, stmts...))
	}

	//c2go
	{
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			govar, cvar = takeAddr(govar, cvar)
			return exprStmt(callExpr(selectorExpr(govar, ident("fromC")), cvar))
		}
		recv := field(starExpr(info.gotype), ident("g"))
		fntype := funcType([]*goast.Field{
			field(starExpr(info.ctype), ident("c")),
		}, nil)
		stmts := cinfo.c2go(ident("g"), ident("c"))
		g.target.addGo(funcDecl(ident("fromC"), recv, fntype, stmts...))
	}

	//refc2go
	{
		if cinfo.refc2go != nil {
			info.refc2go = func(govar, cvar goast.Expr) goast.Stmt {
				stmts := cinfo.refc2go(govar, cvar)
				return blockStmt(stmts...)
			}
		}
	}

	g.types[n.Type] = info
	return info
}

func (g *generator) mapConstExpr(node cast.Node) goast.Expr {
	switch v := node.(type) {
	case *cast.IntegerLiteral:
		return &goast.BasicLit{Kind: token.INT, Value: v.Value}
	case *cast.DeclRefExpr:
		return g.consts[v.Name]
	case *cast.ParenExpr:
		return &goast.ParenExpr{
			Lparen: token.Pos(1),
			X:      g.mapConstExpr(v.ChildNodes[0]),
			Rparen: token.Pos(1),
		}
	case *cast.BinaryOperator:
		return &goast.BinaryExpr{
			X:  g.mapConstExpr(v.ChildNodes[0]),
			Op: mapOperator(v.Operator),
			Y:  g.mapConstExpr(v.ChildNodes[1]),
		}
	case *cast.UnaryOperator:
		return &goast.UnaryExpr{
			X:  g.mapConstExpr(v.ChildNodes[0]),
			Op: mapOperator(v.Operator),
		}
	default:
		halt("convConst()", node)
	}
	return nil
}

func (g *generator) genEnumType(n *cast.EnumType) *typeInfo {
	if info, ok := g.types[n.Name]; ok {
		return info
	}

	var enumDecl *cast.EnumDecl
	{
		node := g.nodes[n.Name]
		if node == nil {
			halt("No decl for record type", n)
		}
		enumDecl = node.(*cast.EnumDecl)
	}
	name := strings.Trim(enumDecl.Name, "enum ")
	info := &typeInfo{}

	if name != "" {
		info.gotype = ident("int")
		if _, ok := g.nodes[name]; ok {
			info.ctype = ident("C." + name)
			info.csize = ident("C.sizeof_" + name)
		} else {
			info.ctype = ident("C.enum" + name)
			info.csize = ident("C.sizeof_enum_" + name)
		}
	} else {
		halt("Unamed enum isn't implemented", n)
	}

	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		return assignStmt1n1(govar, callExpr(info.gotype, cvar))
	}
	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		return assignStmt1n1(cvar, callExpr(info.ctype, govar))
	}

	g.target.addGo(&goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: ident(name),
				Type: info.gotype,
			},
		},
	})

	//Consts
	{
		specs := []goast.Spec{}
		for _, child := range enumDecl.ChildNodes {
			decl := child.(*cast.EnumConstantDecl)
			goname := ident(decl.Name)
			g.consts[decl.Name] = goname

			var val = g.mapConstExpr(decl.ChildNodes[0])
			if val != nil {
				specs = append(specs, &goast.ValueSpec{
					Names:  []*goast.Ident{goname},
					Type:   info.gotype,
					Values: []goast.Expr{val},
				})
			}
		}
		g.target.addGo(&goast.GenDecl{
			Tok:    token.CONST,
			Lparen: token.Pos(1),
			Specs:  specs,
			Rparen: token.Pos(1),
		})
	}

	g.types[n.Name] = info
	return info
}

func (g *generator) genFunctionProtoType(n *cast.FunctionProtoType) *typeInfo {
	return nil
}

func (g *generator) genQualType(n *cast.QualType) *typeInfo {
	return g.genType(n.ChildNodes[0])
}

func (g *generator) genPointerType(n *cast.PointerType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}

	if n.Type == "void *" {
		info := &typeInfo{
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
		g.types[n.Type] = info
		return info
	} else if n.Type == "const char *" {
		info := &typeInfo{}
		info.gotype = ident("string")
		info.ctype = starExpr(ident("C.char"))
		info.csize = ident("unsafe.Sizeof(uintptr(0))")
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(ident("toGoString"), cvar))
		}
		info.go2cAlloc = true
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(ident("toCString"), govar, ident("_sa")))
		}
		info.refc2go = nil //Strings are immutable in Go
		g.types[n.Type] = info
		return info
	}

	o := n.ChildNodes[0]
	oinfo := g.genType(o)
	if oinfo == nil {
		//Pure pointer
		var typ goast.Expr
		if _, ok := o.(*cast.ElaboratedType); ok {
			typ = starExpr(ident("struct{}"))
		} else {
			typ = starExpr(ident("[0]byte"))
		}
		info := &typeInfo{
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
		g.types[n.Type] = info
		return info
	} else {
		//Reference pointer
		info := &typeInfo{
			ctype:  starExpr(oinfo.ctype),
			gotype: starExpr(oinfo.gotype),
			csize:  callExpr(ident("unsafe.Sizeof"), starExpr(oinfo.ctype)),
		}
		if oinfo.csize != nil {
			//Size is not 0
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(govar, callExpr(ident("new"), oinfo.gotype))
				ifstmt := &goast.IfStmt{
					If: token.Pos(1),
					Cond: &goast.BinaryExpr{
						X:  govar,
						Op: token.EQL,
						Y:  ident("nil"),
					},
					Body: blockStmt(alloc),
				}
				conv := oinfo.c2go(starExpr(govar), starExpr(cvar))
				return blockStmt(ifstmt, conv)
			}
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(cvar, saalloc(oinfo.ctype, oinfo.csize))
				conv := oinfo.go2c(starExpr(govar), starExpr(cvar))
				return blockStmt(alloc, conv)
			}
			info.refc2go = func(govar, cvar goast.Expr) goast.Stmt {
				return oinfo.c2go(starExpr(govar), starExpr(cvar))
			}
		} else {
			//No size,so no alloc
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			}
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			}
		}
		g.types[n.Type] = info
		return info
	}
}

func (g *generator) genConstantArrayType(n *cast.ConstantArrayType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}
	info := &typeInfo{}

	oinfo := g.genType(n.ChildNodes[0])
	if oinfo == nil {
		g.types[n.Type] = nil
		return nil
	}

	info.gotype = &goast.ArrayType{
		Len: &goast.BasicLit{
			Kind:  token.INT,
			Value: strconv.Itoa(n.Size),
		},
		Elt: oinfo.gotype,
	}
	info.ctype = &goast.ArrayType{
		Len: &goast.BasicLit{
			Kind:  token.INT,
			Value: strconv.Itoa(n.Size),
		},
		Elt: oinfo.ctype,
	}
	info.csize = &goast.BinaryExpr{
		X:  oinfo.csize,
		Op: token.MUL,
		Y: &goast.BasicLit{
			Kind:  token.INT,
			Value: strconv.Itoa(n.Size),
		},
	}

	rangeStmt := func(x goast.Expr, stmts ...goast.Stmt) *goast.RangeStmt {
		return &goast.RangeStmt{
			For:   token.Pos(1),
			Key:   ident("i"),
			Value: ident("_"),
			Tok:   token.DEFINE,
			X:     x,
			Body:  blockStmt(stmts...),
		}
	}

	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		stmt := oinfo.c2go(indexExpr(govar, ident("i")), indexExpr(cvar, ident("i")))
		return rangeStmt(govar, stmt)
	}

	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		stmt := oinfo.go2c(indexExpr(govar, ident("i")), indexExpr(cvar, ident("i")))
		return rangeStmt(govar, stmt)
	}

	info.go2cAlloc = oinfo.go2cAlloc

	if oinfo.refc2go != nil {
		info.refc2go = func(govar, cvar goast.Expr) goast.Stmt {
			stmt := oinfo.refc2go(indexExpr(govar, ident("i")), indexExpr(cvar, ident("i")))
			return rangeStmt(govar, stmt)
		}
	}

	g.types[n.Type] = info
	return info
}

func (g *generator) genTypedefType(n *cast.TypedefType) *typeInfo {
	if info, ok := g.types[n.Type]; ok {
		return info
	}
	info := &typeInfo{}
	g.types[n.Type] = info

	o := n.ChildNodes[1]
	oinfo := g.genType(o)
	checkTypeInfo(oinfo, o)
	//Skip same name struct/enum/union typedef
	if et, ok := o.(*cast.ElaboratedType); ok {
		name := et.Type
		name = strings.TrimPrefix(name, "enum ")
		name = strings.TrimPrefix(name, "struct ")
		name = strings.TrimPrefix(name, "union ")
		if name == n.Type {
			*info = *oinfo
			return info
		}
	}

	//standard typedef
	m := map[string]string{
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
	if std, ok := m[n.Type]; ok {
		info.gotype = ident(std)
		info.ctype = ident("C." + n.Type)
		info.csize = ident("C.sizeof_" + n.Type)
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(info.gotype, cvar))
		}
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(info.ctype, govar))
		}
		return info
	}

	//function pointer typedef
	if pt, ok := o.(*cast.PointerType); ok {
		if ptt, ok := pt.ChildNodes[0].(*cast.ParenType); ok {
			if fpt, ok := ptt.ChildNodes[0].(*cast.FunctionProtoType); ok {
				info.gotype = ident(n.Type)
				info.ctype = ident("C." + n.Type)
				info.csize = ident("C.sizeof_" + n.Type)
				info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
					return assignStmt1n1(selectorExpr(govar, ident("Raw")), cvar)
				}
				info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
					return assignStmt1n1(cvar, selectorExpr(govar, ident("Raw")))
				}

				g.target.addGo(typeDecl(info.gotype.(*goast.Ident), &goast.StructType{
					Fields: &goast.FieldList{
						List: []*goast.Field{
							field(ident("C."+n.Type), ident("Raw")),
						},
					},
				}))

				decl := g.nodes[n.Type].(*cast.TypedefDecl)
				g.genBridgeCall(decl, info, fpt)

				return info
			}
		}
	}

	info.gotype = ident(n.Type)
	info.ctype = ident("C." + n.Type)
	info.csize = ident("C.sizeof_" + n.Type)

	//forword style declaration
	if _, ok := oinfo.gotype.(*goast.StarExpr); ok {
		if oinfo.refc2go == nil {
			//means it's a pure pointer
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, callExpr(info.gotype, cvar))
			}
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, callExpr(info.ctype, govar))
			}

			g.target.addGo(typeDecl(info.gotype.(*goast.Ident), info.ctype))
			return info
		}
	}

	g.target.addGo(typeDecl(info.gotype.(*goast.Ident), oinfo.gotype))
	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		_temp := ident("_temp")
		def := varDeclStmt(oinfo.gotype, _temp)
		convc := oinfo.c2go(_temp, callExpr(parenExpr(oinfo.ctype), cvar))
		convbase := assignStmt1n1(govar, callExpr(info.gotype, _temp))
		return blockStmt(def, convc, convbase)
	}

	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		_temp := ident("_temp")
		def := varDeclStmt(oinfo.ctype, _temp)
		convc := oinfo.go2c(callExpr(parenExpr(oinfo.gotype), govar), _temp)
		convbase := assignStmt1n1(cvar, callExpr(info.ctype, _temp))
		return blockStmt(def, convc, convbase)
	}
	info.refc2go = oinfo.refc2go //for pointer type case
	return info
}

func (g *generator) genBridgeCall(decl *cast.TypedefDecl, info *typeInfo, fpt *cast.FunctionProtoType) {
	var goparams, goresults []*goast.Field
	var cparams, cresults []*goast.Field

	var cdef goast.Stmt
	var allocdefs []goast.Stmt
	var pconvs []goast.Stmt
	var call goast.Stmt
	var rconvs []goast.Stmt
	var prconvs []goast.Stmt
	var retstmt goast.Stmt

	//Param
	{
		fieldDecls := []*cast.FieldDecl{}
		for i, cn := range fpt.ChildNodes[1:] {
			fieldDecls = append(fieldDecls, &cast.FieldDecl{
				Name:       "arg" + strconv.Itoa(i),
				ChildNodes: []cast.Node{cn},
			})
		}

		info := g.genCompType(fieldDecls, g.hints[decl.Name])
		goparams = info.gofields
		cparams = info.cfields
		pconvs = info.go2c(nil, ident("c"))
		if info.refc2go != nil {
			prconvs = info.refc2go(nil, ident("c"))
		}

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
		info := g.genType(fpt.ChildNodes[0])
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
		args = append(args, selectorExpr(ident("p"), ident("Raw")))
		for _, cparam := range cparams {
			args = append(args, selectorExpr(ident("c"), cparam.Names[0]))
		}
		if len(cresults) > 0 {
			rs := selectorExpr(ident("c"), cresults[0].Names[0])
			call = assignStmt1n1(rs, callExpr(ident("C.call"+decl.Name), args...))
			retstmt = &goast.ReturnStmt{}
		} else {
			call = exprStmt(callExpr(ident("C.call"+decl.Name), args...))
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

		g.target.addC(decl)
		gofntype := funcType(goparams, goresults)
		gofndef := funcDecl(ident("Call"), field(info.gotype, ident("p")), gofntype, stmts...)
		g.target.addGo(gofndef)
	}
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
	var info *typeInfo

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
		"void":                   "",
	}
	gotypestr := m[n.Type]
	ctypestr := m2[n.Type]

	if gotypestr == "" || ctypestr == "" {
		g.types[n.Type] = nil
		return nil
	}

	gotype := ident(gotypestr)
	ctype := ident("C." + ctypestr)
	csize := ident("C.sizeof_" + ctypestr)

	info = &typeInfo{
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
	g.types[n.Type] = info
	return info
}

func (g *generator) genCompType(fieldDecls []*cast.FieldDecl, hints []fieldHint) *compTypeInfo {
	if len(fieldDecls) != len(hints) {
		halt("Length of hints is not match of fields", nil)
	}

	info := &compTypeInfo{}

	go2cFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	c2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	refc2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}

	for _, fieldDecl := range fieldDecls {
		ftn := fieldDecl.ChildNodes[0]
		finfo := g.genType(ftn)
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
			isPointerToConst := false
			if pt, ok := ftn.(*cast.PointerType); ok {
				if qt, ok := pt.ChildNodes[0].(*cast.QualType); ok {
					if qt.Kind == "const" {
						isPointerToConst = true
					}
				}
			}
			//Not restict, but it fit most cases
			if !isPointerToConst {
				refc2goFns = append(refc2goFns, func(goscope, cscope goast.Expr) goast.Stmt {
					return finfo.refc2go(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
				})
			}
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

	if len(refc2goFns) > 0 {
		info.refc2go = func(goscope, cscope goast.Expr) []goast.Stmt {
			stmts := []goast.Stmt{}
			for _, fn := range refc2goFns {
				stmts = append(stmts, fn(goscope, cscope))
			}
			return stmts
		}
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
		info := g.genCompType(fieldDecls, g.hints[fn.Name])
		goparams = info.gofields
		cparams = info.cfields
		pconvs = info.go2c(nil, ident("c"))
		if info.refc2go != nil {
			prconvs = info.refc2go(nil, ident("c"))
		}

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
	analyzeHints(g.hints, src)

	for _, node := range src {
		switch n := node.(type) {
		case *cast.TypedefDecl:
			g.nodes[n.Name] = n
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

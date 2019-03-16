package main

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/iancoleman/strcase"
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

type hint struct {
	isArray     map[string]bool
	isArraySize map[string]bool
	argNames    map[string][]string
}

func (h *hint) init() {
	h.isArray = make(map[string]bool, 128)
	h.isArraySize = make(map[string]bool, 128)
	h.argNames = make(map[string][]string, 128)
}

type generator struct {
	target Target
	nodes  map[string]cast.Node //name node map
	types  map[string]*typeInfo
	consts map[string]goast.Expr
	funcs  map[string]bool

	hint hint
}

func (g *generator) init() {
	g.nodes = make(map[string]cast.Node, 2048)
	g.types = make(map[string]*typeInfo, 512)
	g.consts = make(map[string]goast.Expr, 128)
	g.funcs = make(map[string]bool, 512)

	g.hint.init()
}

func (g *generator) mapType(node cast.Node, pid string) *typeInfo {
	switch n := node.(type) {
	case *cast.ParenType:
		return g.mapParenType(n, pid)
	case *cast.BuiltinType:
		return g.mapBuiltinType(n, pid)
	case *cast.PointerType:
		return g.mapPointerType(n, pid)
	case *cast.ConstantArrayType:
		return g.mapConstantArrayType(n, pid)
	case *cast.DecayedType:
		return g.mapDecayedType(n, pid)
	case *cast.QualType:
		return g.mapQualType(n, pid)
	case *cast.TypedefType:
		return g.mapTypedefType(n, pid)
	case *cast.ElaboratedType:
		return g.mapElaboratedType(n, pid)
	case *cast.RecordType:
		return g.mapRecordType(n, pid)
	case *cast.EnumType:
		return g.mapEnumType(n, pid)
	case *cast.FunctionProtoType:
		return nil
	case *cast.IncompleteArrayType:
		return nil
	default:
		halt("Unkown type to map", n)
		return nil
	}
}

func (g *generator) mapRecordType(n *cast.RecordType, pid string) *typeInfo {
	var decl *cast.RecordDecl
	{
		node := g.nodes[n.Type]
		if node == nil {
			halt("No decl for record type", n)
		}
		decl = node.(*cast.RecordDecl)
	}

	if strings.HasPrefix(n.Type, "struct ") {
		//Incomplete Struct
		if !decl.Definition {
			info := &typeInfo{}
			info.gotype = ident("C.struct_" + decl.Name)
			info.ctype = ident("C.struct_" + decl.Name)
			info.csize = nil
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			}
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, govar)
			}
			return info
		}
	}

	if decl.Name != "" {
		return g.genRecordType(decl)
	} else {
		halt("Unamed struct isn't implemented", n)
		return nil
	}
}

func (g *generator) genRecordType(decl *cast.RecordDecl) *typeInfo {
	switch decl.Kind {
	case "struct":
		return g.genRecordTypeStruct(decl)
	case "union":
		return g.genRecordTypeUnion(decl)
	}
	halt("Unkown record type", decl)
	return nil
}

func (g *generator) genRecordTypeUnion(decl *cast.RecordDecl) *typeInfo {
	if info, ok := g.types[decl.Name]; ok {
		return info
	}
	var info *typeInfo = &typeInfo{}
	g.types[decl.Name] = info

	info.gotype = ident(trimPrefixs(decl.Name, "Vk"))
	if _, ok := g.nodes[decl.Name]; ok {
		//same name typedef exists
		info.ctype = ident("C." + decl.Name)
		info.csize = ident("C.sizeof_" + decl.Name)
	} else {
		info.ctype = ident("C.union_" + decl.Name)
		info.csize = ident("C.sizeof_union_" + decl.Name)
	}

	g.target.addGo(typeDecl(info.gotype.(*goast.Ident), info.ctype))

	recv := field(starExpr(info.gotype), ident("g"))
	for _, fieldDecl := range decl.ChildNodes {
		dn := fieldDecl.(*cast.FieldDecl)
		id := decl.Name + "." + dn.Name
		finfo := g.mapType(dn.ChildNodes[0], id)
		gostr := toGoFieldName(dn.Name)

		//assign
		funcType0 := funcType([]*goast.Field{
			field(finfo.gotype, ident("v")),
		}, nil)
		stmts0 := []goast.Stmt{}
		{
			stmts0 = append(stmts0, varDeclStmt(finfo.ctype, ident("cv")))
			stmts0 = append(stmts0, finfo.go2c(ident("v"), ident("cv")))
			lhs := starExpr(callExpr(starExpr(finfo.ctype), callExpr(ident("unsafe.Pointer"), ident("g"))))
			stmts0 = append(stmts0, assignStmt1n1(lhs, ident("cv")))
		}
		g.target.addGo(funcDecl(ident("Assgin"+upFirst(gostr)), recv, funcType0, stmts0...))

		//get
		funcType1 := funcType(nil, []*goast.Field{
			field(finfo.gotype, ident("v")),
		})
		stmts1 := []goast.Stmt{}
		{
			rhs := starExpr(callExpr(starExpr(finfo.ctype), callExpr(ident("unsafe.Pointer"), ident("g"))))
			stmts1 = append(stmts1, assignStmt1n1D(ident("cv"), rhs))
			stmts1 = append(stmts1, finfo.c2go(ident("v"), ident("cv")))
			stmts1 = append(stmts1, &goast.ReturnStmt{})
		}
		g.target.addGo(funcDecl(ident(upFirst(gostr)), recv, funcType1, stmts1...))
	}

	info.go2cAlloc = false
	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		//no alloc
		return assignStmt1n1(cvar, callExpr(info.ctype, govar))
	}
	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		return assignStmt1n1(govar, callExpr(info.gotype, cvar))
	}
	info.refc2go = nil

	return info
}

func (g *generator) genRecordTypeStruct(decl *cast.RecordDecl) *typeInfo {
	if info, ok := g.types[decl.Name]; ok {
		return info
	}
	var info *typeInfo = &typeInfo{}
	g.types[decl.Name] = info

	var cinfo *compTypeInfo
	var isStructure bool
	{
		fields := []*cast.FieldDecl{}
		for i, child := range decl.ChildNodes {
			fd := child.(*cast.FieldDecl)
			if i == 0 && fd.Name == "sType" {
				isStructure = true
			}
			fields = append(fields, fd)
		}
		cinfo = g.mapCompType(fields, decl.Name)
	}

	info.gotype = ident(trimPrefixs(decl.Name, "Vk"))
	if _, ok := g.nodes[decl.Name]; ok {
		//same name typedef exists
		info.ctype = ident("C." + decl.Name)
		info.csize = ident("C.sizeof_" + decl.Name)
	} else {
		info.ctype = ident("C.struct_" + decl.Name)
		info.csize = ident("C.sizeof_struct_" + decl.Name)
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

	//go2c
	{
		recv := field(starExpr(info.gotype), ident("g"))
		var fntype *goast.FuncType
		if cinfo.go2cAlloc {
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				govar, cvar = takeAddr(govar, cvar)
				return exprStmt(callExpr(selectorExpr(govar, ident("toC")), cvar, ident("m")))
			}
			fntype = funcType([]*goast.Field{
				field(starExpr(info.ctype), ident("c")),
				field(starExpr(ident("cmemory")), ident("m")),
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

	if isStructure {
		g.genStructureMethods(decl.Name, info)
	}
	return info
}

func (g *generator) genStructureMethods(name string, info *typeInfo) {
	/*VkStructureType value is obtained by taking the name of the structure,
	  stripping the leading Vk, prefixing each capital letter with _,
	  converting the entire resulting string to upper case,
	  and prefixing it with VK_STRUCTURE_TYPE_.*/
	sType := strings.TrimPrefix(name, "struct ")
	sType = strings.TrimPrefix(sType, "Vk")
	sType = strcase.ToScreamingSnake(sType)
	sType = "VK_STRUCTURE_TYPE_" + sType

	//sType()
	{
		fnT := funcType(nil, []*goast.Field{field(ident("C.VkStructureType"))})
		ret := returnStmt(ident("C." + sType))
		fn := funcDecl(ident("sType"), field(starExpr(info.gotype), ident("s")), fnT, ret)
		g.target.addGo(fn)
	}

	//toCStructure()
	{
		fnT := funcType([]*goast.Field{
			field(starExpr(ident("cmemory")), ident("m")),
		}, []*goast.Field{
			field(ident("unsafe.Pointer")),
		})
		alloc := assignStmt1n1D(ident("c"), m_alloc(info.ctype, info.csize))
		conv := info.go2c(starExpr(ident("s")), starExpr(ident("c")))
		ret := returnStmt(callExpr(ident("unsafe.Pointer"), ident("c")))
		fn := funcDecl(ident("toCStructure"), field(starExpr(info.gotype), ident("s")), fnT, alloc, conv, ret)
		g.target.addGo(fn)
	}

	//fromCStructure()
	{
		fnT := funcType([]*goast.Field{field(ident("unsafe.Pointer"), ident("p"))}, nil)
		cast := assignStmt1n1D(ident("c"), callExpr(starExpr(info.ctype), ident("p")))
		conv := info.c2go(starExpr(ident("s")), starExpr(ident("c")))
		fn := funcDecl(ident("fromCStructure"), field(starExpr(info.gotype), ident("s")), fnT, cast, conv)
		g.target.addGo(fn)
	}

	//GetNext()
	{
		fnT := funcType(nil, []*goast.Field{field(ident("Structure"))})
		ret := returnStmt(selectorExpr(ident("s"), ident("Next")))
		fn := funcDecl(ident("GetNext"), field(starExpr(info.gotype), ident("s")), fnT, ret)
		g.target.addGo(fn)
	}

	//SetNext()
	{
		fnT := funcType([]*goast.Field{field(ident("Structure"), ident("n"))}, nil)
		ass := assignStmt1n1(selectorExpr(ident("s"), ident("Next")), ident("n"))
		fn := funcDecl(ident("SetNext"), field(starExpr(info.gotype), ident("s")), fnT, ass)
		g.target.addGo(fn)
	}
}

func (g *generator) convConst(node cast.Node) goast.Expr {
	switch v := node.(type) {
	case *cast.IntegerLiteral:
		return &goast.BasicLit{Kind: token.INT, Value: v.Value}
	case *cast.DeclRefExpr:
		return g.consts[v.Name]
	case *cast.ParenExpr:
		return &goast.ParenExpr{
			Lparen: token.Pos(1),
			X:      g.convConst(v.ChildNodes[0]),
			Rparen: token.Pos(1),
		}
	case *cast.BinaryOperator:
		return &goast.BinaryExpr{
			X:  g.convConst(v.ChildNodes[0]),
			Op: convOperator(v.Operator),
			Y:  g.convConst(v.ChildNodes[1]),
		}
	case *cast.UnaryOperator:
		return &goast.UnaryExpr{
			X:  g.convConst(v.ChildNodes[0]),
			Op: convOperator(v.Operator),
		}
	default:
		halt("convConst()", node)
	}
	return nil
}

func (g *generator) mapEnumType(n *cast.EnumType, pid string) *typeInfo {
	var decl *cast.EnumDecl
	{
		node := g.nodes[n.Name]
		if node == nil {
			halt("No decl for record type", n)
		}
		decl = node.(*cast.EnumDecl)
	}

	name := strings.TrimPrefix(decl.Name, "enum ")
	if name != "" {
		return g.genEnumType(decl)
	} else {
		halt("Unamed enum isn't implemented", n)
		return nil
	}
}

func (g *generator) genEnumType(decl *cast.EnumDecl) *typeInfo {
	if info, ok := g.types[decl.Name]; ok {
		return info
	}
	var info *typeInfo

	info = &typeInfo{}
	g.types[decl.Name] = info

	goname := trimPrefixs(decl.Name, "Vk")
	info.gotype = ident(goname)
	if _, ok := g.nodes[decl.Name]; ok {
		info.ctype = ident("C." + decl.Name)
		info.csize = ident("C.sizeof_" + decl.Name)
	} else {
		info.ctype = ident("C.enum_" + decl.Name)
		info.csize = ident("C.sizeof_enum_" + decl.Name)
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
				Name: info.gotype.(*goast.Ident),
				Type: ident("int"),
			},
		},
	})

	//Consts
	{
		specs := []goast.Spec{}
		for _, child := range decl.ChildNodes {
			d := child.(*cast.EnumConstantDecl)
			goname := ident(trimPrefixs(d.Name, "VK_"))
			g.consts[d.Name] = goname

			var val = g.convConst(d.ChildNodes[0])
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

	return info
}

func (g *generator) mapQualType(n *cast.QualType, pid string) *typeInfo {
	id := pid + "." + n.Kind
	return g.mapType(n.ChildNodes[0], id)
}

func (g *generator) mapParenType(n *cast.ParenType, pid string) *typeInfo {
	return g.mapType(n.ChildNodes[0], pid)
}

func (g *generator) mapPointerType(n *cast.PointerType, pid string) *typeInfo {
	isArray := g.hint.isArray[pid]

	//Array
	if isArray {
		o := n.ChildNodes[0]
		oinfo := g.mapType(o, pid+".*")
		level := strconv.Itoa(len(strings.Split(pid, ".")))

		info := &typeInfo{}
		if oinfo == nil {
			info.gotype = arrayType(ident("byte"), nil)
			info.ctype = ident("unsafe.Pointer")
			info.csize = ident("C.sizeof_void_pointer")
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				slice := ident("slice" + level)
				i := ident("i" + level)
				size := mulExpr(ident("C.sizeof_char"), uintLen(govar))
				toslice := asSliceExpr(nil, cvar, callExpr(ident("len"), govar))
				stmt1 := assignStmt1n1(cvar, m_alloc(nil, size))
				stmt2 := assignStmt1n1D(slice, toslice)
				stmt3s := assignStmt1n1(indexExpr(slice, i), indexExpr(govar, i))
				stmt3 := rangeStmti(govar, i, stmt3s)
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(callExpr(ident("len"), govar), ident("0"), token.NEQ),
					Body: blockStmt(stmt1, stmt2, stmt3),
					Else: assignStmt1n1(cvar, ident("nil")),
				}
			}
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				slice := ident("slice" + level)
				i := ident("i" + level)
				toslice := asSliceExpr(nil, cvar, callExpr(ident("len"), govar))
				assignSlice := assignStmt1n1D(slice, toslice)
				assignValue := assignStmt1n1(indexExpr(govar, i), indexExpr(slice, i))
				rangeSlice := rangeStmti(govar, i, assignValue)
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(callExpr(ident("len"), govar), ident("0"), token.NEQ),
					Body: blockStmt(assignSlice, rangeSlice),
				}
			}
			info.refc2go = info.c2go
		} else {
			info.gotype = arrayType(oinfo.gotype, nil)
			info.ctype = starExpr(oinfo.ctype)
			info.csize = ident("C.sizeof_void_pointer")
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				slice := ident("slice" + level)
				i := ident("i" + level)
				size := mulExpr(oinfo.csize, uintLen(govar))
				toslice := asSliceExpr(oinfo.ctype, cvar, callExpr(ident("len"), govar))
				stmt1 := assignStmt1n1(cvar, m_alloc(oinfo.ctype, size))
				stmt2 := assignStmt1n1D(slice, toslice)
				stmt3s := oinfo.go2c(indexExpr(govar, i), indexExpr(slice, i))
				stmt3 := rangeStmti(govar, i, stmt3s)
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(callExpr(ident("len"), govar), ident("0"), token.NEQ),
					Body: blockStmt(stmt1, stmt2, stmt3),
					Else: assignStmt1n1(cvar, ident("nil")),
				}
			}
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				slice := ident("slice" + level)
				i := ident("i" + level)
				toslice := asSliceExpr(oinfo.ctype, cvar, callExpr(ident("len"), govar))
				assignSlice := assignStmt1n1D(slice, toslice)
				assignValue := oinfo.c2go(indexExpr(govar, i), indexExpr(slice, i))
				rangeSlice := rangeStmti(govar, i, assignValue)
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(callExpr(ident("len"), govar), ident("0"), token.NEQ),
					Body: blockStmt(assignSlice, rangeSlice),
				}
			}
			info.refc2go = info.c2go
		}
		return info
	}

	//Map as unsafe.Pointer
	if n.Type == "void *" || n.Type == "const void *" {
		info := &typeInfo{
			ctype:  ident("unsafe.Pointer"),
			gotype: ident("unsafe.Pointer"),
			csize:  ident("C.sizeof_void_pointer"),
		}
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			//No alloc for unsafe.Pointer
			return assignStmt1n1(cvar, govar)
		}
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, cvar)
		}
		info.refc2go = nil
		return info
	}

	//Map as string
	if n.Type == "const char *" {
		info := &typeInfo{}
		info.gotype = ident("string")
		info.ctype = starExpr(ident("C.char"))
		info.csize = ident("C.sizeof_void_pointer")
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(ident("toGoString"), cvar))
		}
		info.go2cAlloc = true
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(ident("toCString"), govar, ident("m")))
		}
		info.refc2go = nil //Strings are immutable in Go
		return info
	}

	o := n.ChildNodes[0]
	oinfo := g.mapType(o, pid+".*")

	//Pure pointer
	if oinfo == nil {
		var typ goast.Expr
		if _, ok := o.(*cast.ElaboratedType); ok {
			typ = starExpr(ident("struct{}"))
		} else {
			typ = starExpr(ident("[0]byte"))
		}
		info := &typeInfo{
			ctype:  typ,
			gotype: typ,
			csize:  ident("C.sizeof_void_pointer"),
			go2c: func(govar, cvar goast.Expr) goast.Stmt {
				//no alloc
				return assignStmt1n1(cvar, govar)
			},
			c2go: func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			},
			refc2go: nil,
		}
		return info
	}

	//Reference pointer
	{
		info := &typeInfo{
			ctype:  starExpr(oinfo.ctype),
			gotype: starExpr(oinfo.gotype),
			csize:  ident("C.sizeof_void_pointer"),
		}
		if oinfo.csize != nil {
			//Size is not 0
			info.go2cAlloc = true
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				alloc := assignStmt1n1(cvar, m_alloc(oinfo.ctype, oinfo.csize))
				conv := oinfo.go2c(starExpr(govar), starExpr(cvar))
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(govar, ident("nil"), token.NEQ),
					Body: blockStmt(alloc, conv),
					Else: assignStmt1n1(cvar, ident("nil")),
				}
			}
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				stmt := oinfo.c2go(starExpr(govar), starExpr(cvar))
				block, ok := stmt.(*goast.BlockStmt)
				if !ok {
					block = blockStmt(stmt)
				}
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(govar, ident("nil"), token.NEQ),
					Body: block,
				}
			}
			info.refc2go = info.c2go
		} else {
			//No size
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				//no alloc
				return assignStmt1n1(cvar, govar)
			}
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, cvar)
			}
			info.refc2go = nil
		}
		return info
	}
}

func (g *generator) mapConstantArrayType(n *cast.ConstantArrayType, pid string) *typeInfo {
	id := pid + "." + fmt.Sprintf("[%d]", n.Size)
	oinfo := g.mapType(n.ChildNodes[0], id)

	if oinfo == nil {
		return nil
	}

	info := &typeInfo{}
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

	return info
}

func (g *generator) mapDecayedType(n *cast.DecayedType, pid string) *typeInfo {
	an := n.ChildNodes[0].(*cast.ConstantArrayType)
	id := pid + "." + fmt.Sprintf("[%d]", an.Size)
	oinfo := g.mapType(an.ChildNodes[0], id)

	if oinfo == nil {
		return nil
	}

	info := &typeInfo{}
	info.gotype = &goast.ArrayType{
		Len: &goast.BasicLit{
			Kind:  token.INT,
			Value: strconv.Itoa(an.Size),
		},
		Elt: oinfo.gotype,
	}
	info.ctype = starExpr(oinfo.ctype)
	info.csize = ident("C.sizeof_void_pointer")

	level := strconv.Itoa(len(strings.Split(pid, ".")))
	size := intLit(an.Size)
	info.go2cAlloc = true
	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		slice := ident("slice" + level)
		i := ident("i" + level)
		memSize := mulExpr(oinfo.csize, size)
		toslice := asSliceExpr(oinfo.ctype, cvar, size)
		stmt1 := assignStmt1n1(cvar, m_alloc(oinfo.ctype, memSize))
		stmt2 := assignStmt1n1D(slice, toslice)
		stmt3s := oinfo.go2c(indexExpr(govar, i), indexExpr(slice, i))
		stmt3 := rangeStmti(govar, i, stmt3s)
		return blockStmt(stmt1, stmt2, stmt3)
	}
	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		slice := ident("slice" + level)
		i := ident("i" + level)
		toslice := asSliceExpr(oinfo.ctype, cvar, size)
		assignSlice := assignStmt1n1D(slice, toslice)
		assignValue := oinfo.c2go(indexExpr(govar, i), indexExpr(slice, i))
		rangeSlice := rangeStmti(govar, i, assignValue)
		return blockStmt(assignSlice, rangeSlice)
	}

	if oinfo.refc2go != nil {
		info.refc2go = func(govar, cvar goast.Expr) goast.Stmt {
			slice := ident("slice" + level)
			i := ident("i" + level)
			toslice := asSliceExpr(oinfo.ctype, cvar, size)
			assignSlice := assignStmt1n1D(slice, toslice)
			assignValue := oinfo.refc2go(indexExpr(govar, i), indexExpr(slice, i))
			rangeSlice := rangeStmti(govar, i, assignValue)
			return blockStmt(assignSlice, rangeSlice)
		}
	}

	return info
}

func (g *generator) mapTypedefType(n *cast.TypedefType, pid string) *typeInfo {
	//Standard Typedef
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
		info := &typeInfo{}
		info.gotype = ident(std)
		info.ctype = ident("C." + n.Type)
		info.csize = ident("C.sizeof_" + n.Type)
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, callExpr(info.gotype, cvar))
		}
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(cvar, callExpr(info.ctype, govar))
		}
		info.refc2go = nil
		return info
	}

	if n.Type == "VkBool32" {
		info := &typeInfo{}
		info.gotype = ident("bool")
		info.ctype = ident("C." + n.Type)
		info.csize = ident("C.sizeof_" + n.Type)
		info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
			return assignStmt1n1(govar, binExpr(cvar, intLit(0), token.NEQ))
		}
		info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
			return &goast.IfStmt{
				Cond: govar,
				Body: blockStmt(assignStmt1n1(cvar, intLit(1))),
				Else: assignStmt1n1(cvar, intLit(0)),
			}
		}
		info.refc2go = nil
		return info
	}

	decl := g.nodes[n.Type].(*cast.TypedefDecl)
	o := decl.ChildNodes[0]

	//Same Name Typedef（struct/enum/union）
	if et, ok := o.(*cast.ElaboratedType); ok {
		name := et.Type
		name = strings.TrimPrefix(name, "enum ")
		name = strings.TrimPrefix(name, "struct ")
		name = strings.TrimPrefix(name, "union ")
		if name == decl.Name {
			oinfo := g.mapType(o, "")
			checkTypeInfoNotNil(oinfo, o)
			info := &typeInfo{}
			*info = *oinfo
			return info
		}
	}

	return g.genTypedefDecl(decl)
}

func (g *generator) genTypedefDecl(decl *cast.TypedefDecl) *typeInfo {
	if info, ok := g.types[decl.Name]; ok {
		return info
	}
	info := &typeInfo{}
	g.types[decl.Name] = info
	var name string
	if strings.HasPrefix(decl.Name, "PFN_vk") {
		name = trimPrefixs(decl.Name, "PFN_vk")
		name = "PFN" + name
	} else {
		name = trimPrefixs(decl.Name, "Vk")
	}
	info.gotype = ident(name)
	info.ctype = ident("C." + decl.Name)
	info.csize = ident("C.sizeof_" + decl.Name)

	o := decl.ChildNodes[0]

	//Function Pointer Typedef
	if pt, ok := o.(*cast.PointerType); ok {
		if ptt, ok := pt.ChildNodes[0].(*cast.ParenType); ok {
			if fpt, ok := ptt.ChildNodes[0].(*cast.FunctionProtoType); ok {

				info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
					cast := callExpr(info.gotype, cvar)
					return assignStmt1n1(govar, cast)
				}
				info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
					cast := callExpr(info.ctype, govar)
					return assignStmt1n1(cvar, cast)
				}
				info.refc2go = nil
				g.target.addGo(typeDecl(info.gotype.(*goast.Ident), info.ctype))
				g.genBridgeCall(decl, info, fpt, decl.Name)
				return info
			}
		}
	}

	oinfo := g.mapType(o, decl.Name)

	//Forword Typedef
	if _, ok := oinfo.gotype.(*goast.StarExpr); ok {
		if oinfo.refc2go == nil {
			//means it's a pure pointer
			info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(govar, callExpr(info.gotype, cvar))
			}
			info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
				return assignStmt1n1(cvar, callExpr(info.ctype, govar))
			}
			info.refc2go = nil

			g.target.addGo(typeDecl(info.gotype.(*goast.Ident), info.ctype))
			return info
		}
	}

	//Normal Typedef
	g.target.addGo(typeDecl(info.gotype.(*goast.Ident), oinfo.gotype))

	info.go2cAlloc = oinfo.go2cAlloc
	info.go2c = func(govar, cvar goast.Expr) goast.Stmt {
		temp := ident("temp_in_" + decl.Name)
		def := varDeclStmt(oinfo.ctype, temp)
		convc := oinfo.go2c(callExpr(parenExpr(oinfo.gotype), govar), temp)
		convbase := assignStmt1n1(cvar, callExpr(info.ctype, temp))
		return blockStmt(def, convc, convbase)
		// return oinfo.go2c(callExpr(parenExpr(oinfo.gotype), govar), callExpr(parenExpr(oinfo.ctype), cvar))
	}

	info.c2go = func(govar, cvar goast.Expr) goast.Stmt {
		temp := ident("temp_in_" + decl.Name)
		def := varDeclStmt(oinfo.gotype, temp)
		convc := oinfo.c2go(temp, callExpr(parenExpr(oinfo.ctype), cvar))
		convbase := assignStmt1n1(govar, callExpr(info.gotype, temp))
		return blockStmt(def, convc, convbase)
		// return oinfo.c2go(callExpr(parenExpr(oinfo.gotype), govar), callExpr(parenExpr(oinfo.ctype), cvar))
	}

	info.refc2go = oinfo.refc2go //for pointer type case
	return info
}

func (g *generator) genBridgeCall(decl *cast.TypedefDecl, info *typeInfo, fpt *cast.FunctionProtoType, id string) {
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
		argNames := g.hint.argNames[id+".params"]
		if argNames == nil {
			argNames = make([]string, len(fpt.ChildNodes[1:]))
			for i, _ := range argNames {
				argNames[i] = "arg" + strconv.Itoa(i)
			}
		}

		fieldDecls := []*cast.FieldDecl{}
		if len(argNames) != len(fpt.ChildNodes[1:]) {
			halt(fmt.Sprintf("The length of names hints %d are wrong", len(argNames)), decl)
		}

		for i, cn := range fpt.ChildNodes[1:] {
			fieldDecls = append(fieldDecls, &cast.FieldDecl{
				Name:       argNames[i],
				ChildNodes: []cast.Node{cn},
			})
		}

		info := g.mapCompType(fieldDecls, id+"."+"params")
		goparams = info.gofields
		cparams = info.cfields
		pconvs = info.go2c(nil, ident("c"))
		if info.refc2go != nil {
			prconvs = info.refc2go(nil, ident("c"))
		}

		if info.go2cAlloc {
			sa := ident("m")
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
		info := g.mapType(fpt.ChildNodes[0], id+"."+"result")
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
		args = append(args, callExpr(info.ctype, callExpr(ident("unsafe.Pointer"), ident("p"))))
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
		var name string
		if strings.HasPrefix(decl.Name, "PFN_vk") {
			name = trimPrefixs(decl.Name, "PFN_vk")
		} else {
			name = trimPrefixs(decl.Name, "Vk")
		}

		gofntype := funcType(goparams, goresults)
		g.target.addGo(typeDecl(ident("Func"+name), gofntype))

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
		def := funcDecl(ident("To"+name), nil, funcType(
			[]*goast.Field{field(ident("PFNVoidFunction"), ident("p"))},
			[]*goast.Field{field(ident("Func"+name), ident("fn"))},
		),
			returnStmt(&goast.FuncLit{
				Type: gofntype,
				Body: blockStmt(stmts...),
			}))
		g.target.addGo(def)
	}
}

func (g *generator) mapElaboratedType(n *cast.ElaboratedType, pid string) *typeInfo {
	var info *typeInfo
	switch o := n.ChildNodes[0].(type) {
	case *cast.RecordType:
		info = g.mapRecordType(o, pid)
	case *cast.EnumType:
		info = g.mapEnumType(o, pid)
	default:
		halt("Unkown elaborated type", o)
	}
	return info
}

func (g *generator) mapBuiltinType(n *cast.BuiltinType, pid string) *typeInfo {
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
	return info
}

func (g *generator) mapCompType(fieldDecls []*cast.FieldDecl, pid string) *compTypeInfo {
	info := &compTypeInfo{}

	go2cFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	c2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}
	refc2goFns := []func(goscope, cscope goast.Expr) goast.Stmt{}

	var pending func(arrtype goast.Expr, arrname *goast.Ident)
	for i, fieldDecl := range fieldDecls {
		ftn := fieldDecl.ChildNodes[0]
		id := pid + "." + strconv.Itoa(i)
		cname := ident(avoidGoKeyword(fieldDecl.Name))
		goname := ident("")
		if strings.HasSuffix(pid, ".params") {
			goname.Name = lowFirst(toGoFieldName(fieldDecl.Name))
			goname.Name = avoidGoKeyword(goname.Name)
		} else {
			goname.Name = upFirst(toGoFieldName(fieldDecl.Name))
		}
		isPointerToConst := false
		if pt, ok := ftn.(*cast.PointerType); ok {
			if qt, ok := pt.ChildNodes[0].(*cast.QualType); ok {
				if qt.Kind == "const" {
					isPointerToConst = true
				}
			}
		}
		finfo := g.mapType(ftn, id)

		//sType
		if i == 0 && fieldDecl.Name == "sType" {
			info.cfields = append(info.cfields, field(finfo.ctype, cname))
			go2cFns = append(go2cFns, func(goscope, cscope goast.Expr) goast.Stmt {
				return assignStmt1n1(selectorExpr(cscope, ident("sType")), selectorExpr(goscope, ident("sType()")))
			})
			continue
		}

		if pending != nil {
			pending(finfo.gotype, goname)
			pending = nil
		}

		//pNext
		if i == 1 && fieldDecl.Name == "pNext" {
			info.cfields = append(info.cfields, field(finfo.ctype, cname))
			info.gofields = append(info.gofields, field(ident("Structure"), goname))
			info.go2cAlloc = true
			go2cFns = append(go2cFns, func(goscope, cscope goast.Expr) goast.Stmt {
				fn := selectorExpr(selectorExpr(goscope, goname), ident("toCStructure"))
				call := callExpr(fn, ident("m"))
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(selectorExpr(goscope, goname), ident("nil"), token.NEQ),
					Body: blockStmt(assignStmt1n1(selectorExpr(cscope, cname), call)),
				}
			})
			c2go := func(goscope, cscope goast.Expr) goast.Stmt {
				fn := selectorExpr(selectorExpr(goscope, goname), ident("fromCStructure"))
				call := callExpr(fn, selectorExpr(cscope, cname))
				return &goast.IfStmt{
					If:   token.Pos(1),
					Cond: binExpr(selectorExpr(goscope, goname), ident("nil"), token.NEQ),
					Body: blockStmt(exprStmt(call)),
				}
			}
			c2goFns = append(c2goFns, c2go)
			if !isPointerToConst {
				//Not restict, but it fit most cases
				refc2goFns = append(refc2goFns, c2go)
			}
			continue
		}

		//Array Size
		if g.hint.isArraySize[id] {
			pending = func(arrtype goast.Expr, arrname *goast.Ident) {
				info.cfields = append(info.cfields, field(finfo.ctype, cname))
				//No go field
				go2cFns = append(go2cFns, func(goscope, cscope goast.Expr) goast.Stmt {
					n := callExpr(ident("len"), selectorExpr(goscope, arrname))
					cn := callExpr(finfo.ctype, n)
					return assignStmt1n1(selectorExpr(cscope, cname), cn)
				})
				c2goFns = append(c2goFns, func(goscope, cscope goast.Expr) goast.Stmt {
					n := callExpr(ident("int"), selectorExpr(cscope, cname))
					make := callExpr(ident("make"), arrtype, n)
					return assignStmt1n1(selectorExpr(goscope, arrname), make)
				})
				// no refc2goFn
			}
			continue
		}

		info.cfields = append(info.cfields, field(finfo.ctype, cname))
		info.gofields = append(info.gofields, field(finfo.gotype, goname))
		go2cFns = append(go2cFns, func(goscope, cscope goast.Expr) goast.Stmt {
			return finfo.go2c(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
		})
		c2goFns = append(c2goFns, func(goscope, cscope goast.Expr) goast.Stmt {
			return finfo.c2go(selectorExpr(goscope, goname), selectorExpr(cscope, cname))
		})
		if finfo.refc2go != nil && !isPointerToConst {
			//Not restict, but it fit most cases
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
		info := g.mapCompType(fieldDecls, fn.Name+"."+"params")
		goparams = info.gofields
		cparams = info.cfields
		pconvs = info.go2c(nil, ident("c"))
		if info.refc2go != nil {
			prconvs = info.refc2go(nil, ident("c"))
		}

		if info.go2cAlloc {
			sa := ident("m")
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
		info := g.mapType(fn.ChildNodes[0], fn.Name+"."+"result")
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

		name := ident(trimPrefixs(fn.Name, "vk"))
		gofntype := funcType(goparams, goresults)
		gofndef := funcDecl(name, nil, gofntype, stmts...)
		g.target.addGo(gofndef)
	}
}

func (g *generator) process(src Source) {
	analyzeHint(&g.hint, src)

	for _, node := range src {
		switch n := node.(type) {
		case *cast.TypedefDecl:
			g.nodes[n.Name] = n
		case *cast.RecordDecl:
			g.nodes[n.Kind+" "+n.Name] = n
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
			if strings.HasPrefix(n.Name, "PFN_") {
				// println(n.Name)
				g.genTypedefDecl(n)
			}
		case *cast.RecordDecl:
		case *cast.EnumDecl:
			if strings.HasPrefix(n.Name, "Vk") {
				// println(n.Name)
				g.genEnumType(n)
			}
		case *cast.FunctionDecl:
			re := regexp.MustCompile(`.*(KHR)|(EXT)|(AMD)|(ANDROID)|(GOOGLE)|(IMG)|(MVK)|(NN)|(NVX)|(NV)$`)
			if !re.MatchString(n.Name) {
				g.genFunc(n)
			}
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

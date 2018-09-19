package main

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"strings"

	cast "github.com/elliotchance/c2go/ast"
	"github.com/logrusorgru/aurora"
)

type typeInfo struct {
	gotype goast.Expr
	ctype  cast.Node
	c2go   func(goname, cname string, define bool) []goast.Stmt
	go2c   func(cname, goname string, define bool) []goast.Stmt
}

func (ti *typeInfo) cgotype() goast.Expr {
	t := cgotype(ti.ctype)
	return &goast.Ident{Name: t}
}

func (ti *typeInfo) sizeof_cgotype() goast.Expr {
	st := sizeof_cgotype(ti.ctype)
	return &goast.Ident{Name: st}
}

func allocType(node cast.Node) *goast.CallExpr {
	return &goast.CallExpr{
		Fun: &goast.ParenExpr{
			X: &goast.StarExpr{
				X: &goast.Ident{Name: cgotype(node)},
			},
		},
		Lparen: token.Pos(1),
		Rparen: token.Pos(1),
		Args: []goast.Expr{
			&goast.CallExpr{
				Fun:    &goast.Ident{Name: "_sa.alloc"},
				Lparen: token.Pos(1),
				Rparen: token.Pos(1),
				Args: []goast.Expr{
					&goast.Ident{Name: sizeof_cgotype(node)},
				},
			},
		},
	}
}

func cgotype(node cast.Node) string {
	switch n := node.(type) {
	case *cast.BuiltinType:
		{
			var m = map[string]string{
				"char":                   "C.char",
				"singed char":            "C.schar",
				"unsigned char":          "C.uchar",
				"short":                  "C.short",
				"unsigned short":         "C.ushort",
				"int":                    "C.int",
				"unsigned int":           "C.uint",
				"long":                   "C.long",
				"unsigned long":          "C.ulong",
				"long long int":          "C.longlong",
				"unsigned long long int": "C.ulonglong",
				"float":                  "C.float",
				"double":                 "C.double",
			}
			return m[n.Type]
		}
	case *cast.PointerType:
		return "C." + n.Type
	case *cast.ConstantArrayType:
		return "C." + n.Type
	case *cast.TypedefType:
		return "C." + n.Type
	default:
		halt("Unkown type for cgotype()", node)
		return ""
	}
}

func sizeof_cgotype(node cast.Node) string {
	switch n := node.(type) {
	case *cast.BuiltinType:
		return "C.sizeof_" + n.Type
	case *cast.PointerType:
		return "C.sizeof_" + n.Type
	case *cast.ConstantArrayType:
		return "C.sizeof_" + n.Type
	case *cast.TypedefType:
		return "C.sizeof_" + n.Type
	default:
		halt("Unkown type for sizeof_cgotype()", node)
		return ""
	}
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
	obj := n.ChildNodes[0]
	info := g.genType(obj)
	if info == nil {
		return &typeInfo{
			ctype:  n,
			gotype: &goast.Ident{Name: "unsafe.Pointer"},
			c2go: func(goname, cname string, define bool) []goast.Stmt {
				tok := token.ASSIGN
				if define {
					tok = token.DEFINE
				}
				return []goast.Stmt{
					&goast.AssignStmt{
						Lhs: []goast.Expr{&goast.Ident{Name: goname}},
						Tok: tok,
						Rhs: []goast.Expr{&goast.Ident{Name: cname}},
					},
				}
			},
			go2c: func(cname, goname string, define bool) []goast.Stmt {
				tok := token.ASSIGN
				if define {
					tok = token.DEFINE
				}
				return []goast.Stmt{
					&goast.AssignStmt{
						Lhs: []goast.Expr{&goast.Ident{Name: cname}},
						Tok: tok,
						Rhs: []goast.Expr{&goast.Ident{Name: goname}},
					},
				}
			},
		}
	}

	return &typeInfo{
		ctype:  n,
		gotype: &goast.StarExpr{X: info.gotype},
		c2go: func(goname, cname string, define bool) []goast.Stmt {
			tok := token.ASSIGN
			if define {
				tok = token.DEFINE
			}
			stmts := []goast.Stmt{}
			stmts = append(stmts, &goast.AssignStmt{
				Lhs: []goast.Expr{
					&goast.Ident{Name: goname},
				},
				Tok: tok,
				Rhs: []goast.Expr{
					&goast.CallExpr{
						Fun:    &goast.Ident{Name: "new"},
						Lparen: token.Pos(1),
						Rparen: token.Pos(1),
						Args: []goast.Expr{
							info.gotype,
						},
					},
				},
			})
			stmts = append(stmts, info.c2go("*"+goname, "*"+cname, false)...)
			return stmts
		},
		go2c: func(cname, goname string, define bool) []goast.Stmt {

			tok := token.ASSIGN
			if define {
				tok = token.DEFINE
			}
			stmts := []goast.Stmt{}
			stmts = append(stmts, &goast.AssignStmt{
				Lhs: []goast.Expr{
					&goast.Ident{Name: cname},
				},
				Tok: tok,
				Rhs: []goast.Expr{
					allocType(obj),
				},
			})
			stmts = append(stmts, info.go2c("*"+cname, "*"+goname, false)...)
			return stmts
		},
	}
}

func (g *generator) genConstantArrayType(n *cast.ConstantArrayType) *typeInfo {
	deepPrint(n, 0)
	return nil
}

func (g *generator) genTypedefType(n *cast.TypedefType) *typeInfo {
	deepPrint(n, 0)
	return nil
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
	gotype := m[n.Type]
	if gotype == "" {
		return nil
	}
	return &typeInfo{
		ctype:  n,
		gotype: &goast.Ident{Name: gotype},
		c2go: func(goname, cname string, define bool) []goast.Stmt {
			tok := token.ASSIGN
			if define {
				tok = token.DEFINE
			}
			return []goast.Stmt{
				&goast.AssignStmt{
					Lhs: []goast.Expr{&goast.Ident{Name: goname}},
					Tok: tok,
					Rhs: []goast.Expr{&goast.CallExpr{
						Fun:    &goast.Ident{Name: gotype},
						Lparen: token.Pos(1),
						Args: []goast.Expr{
							&goast.Ident{Name: cname},
						},
						Rparen: token.Pos(1),
					}},
				},
			}
		},
		go2c: func(cname, goname string, define bool) []goast.Stmt {
			tok := token.ASSIGN
			if define {
				tok = token.DEFINE
			}
			return []goast.Stmt{
				&goast.AssignStmt{
					Lhs: []goast.Expr{&goast.Ident{Name: cname}},
					Tok: tok,
					Rhs: []goast.Expr{&goast.CallExpr{
						Fun:    &goast.Ident{Name: cgotype(n)},
						Lparen: token.Pos(1),
						Args: []goast.Expr{
							&goast.Ident{Name: goname},
						},
						Rparen: token.Pos(1),
					}},
				},
			}
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

	decl := &goast.FuncDecl{
		Name: &goast.Ident{},
		Type: &goast.FuncType{
			Func: token.Pos(1),
			Params: &goast.FieldList{
				Opening: token.Pos(1),
				Closing: token.Pos(1),
				List:    []*goast.Field{},
			},
			Results: &goast.FieldList{
				Opening: token.Pos(1),
				Closing: token.Pos(1),
				List:    []*goast.Field{},
			},
		},
		Body: &goast.BlockStmt{
			Lbrace: token.Pos(1),
			List:   []goast.Stmt{},
			Rbrace: token.Pos(1),
		},
	}

	//Name
	{
		name := fn.Name
		name = strings.TrimPrefix(name, "vk")
		decl.Name.Name = name
	}

	hasPointerParam := false
	stmts := []goast.Stmt{}
	_stmts := []goast.Stmt{}
	args := []goast.Expr{}
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

		goarg := pvd.Name
		carg := "c_" + goarg
		decl.Type.Params.List = append(decl.Type.Params.List, &goast.Field{
			Names: []*goast.Ident{&goast.Ident{Name: goarg}},
			Type:  info.gotype,
		})
		stmts = append(stmts, info.go2c(carg, goarg, true)...)
		if _, ok := info.gotype.(*goast.StarExpr); ok {
			subinfo := g.genType(node.(*cast.PointerType).ChildNodes[0])
			_stmts = append(_stmts, subinfo.c2go("*"+goarg, "*"+carg, false)...)
			hasPointerParam = true
		}
		args = append(args, &goast.Ident{Name: carg})
	}

	//Results
	{
		rs := fn.ChildNodes[0]
		info := g.genType(rs)
		if info != nil {
			decl.Type.Results.List = append(decl.Type.Results.List, &goast.Field{
				Names: []*goast.Ident{&goast.Ident{Name: "ret"}},
				Type:  info.gotype,
			})
			stmts = append(stmts, &goast.AssignStmt{
				Lhs: []goast.Expr{&goast.Ident{Name: "c_ret"}},
				Tok: token.DEFINE,
				Rhs: []goast.Expr{
					&goast.CallExpr{
						Fun:    &goast.Ident{Name: "C." + fn.Name},
						Lparen: token.Pos(1),
						Args:   args,
						Rparen: token.Pos(1),
					},
				},
			})
			_stmts = append(_stmts, info.c2go("ret", "c_ret", false)...)
		} else {
			stmts = append(stmts, &goast.ExprStmt{
				X: &goast.CallExpr{
					Fun:    &goast.Ident{Name: "C." + fn.Name},
					Lparen: token.Pos(1),
					Args:   args,
					Rparen: token.Pos(1),
				},
			})
		}
	}

	//stackAllocator
	if hasPointerParam {
		stmts = append([]goast.Stmt{
			&goast.AssignStmt{
				Lhs: []goast.Expr{&goast.Ident{Name: "_sa"}},
				Tok: token.DEFINE,
				Rhs: []goast.Expr{
					&goast.CallExpr{
						Fun:    &goast.Ident{Name: "pool.take"},
						Lparen: token.Pos(1),
						Rparen: token.Pos(1),
					},
				},
			},
			&goast.DeferStmt{
				Defer: token.Pos(1),
				Call: &goast.CallExpr{
					Fun:    &goast.Ident{Name: "pool.give"},
					Lparen: token.Pos(1),
					Args: []goast.Expr{
						&goast.Ident{Name: "_sa"},
					},
					Rparen: token.Pos(1),
				},
			},
		}, stmts...)
	}

	stmts = append(stmts, _stmts...)
	stmts = append(stmts, &goast.ReturnStmt{
		Return: token.Pos(1),
	})

	decl.Body.List = stmts
	g.target.addGo(decl)
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

package main

import (
	goast "go/ast"
	"go/token"

	cast "github.com/elliotchance/c2go/ast"
)

type halting struct {
	msg  string
	node cast.Node
}

func halt(msg string, node cast.Node) {
	panic(&halting{msg, node})
}

func saalloc(typ, size goast.Expr) *goast.CallExpr {
	fun := parenExpr(starExpr(typ))
	arg := callExpr(ident("_sa.alloc"), size)
	call := callExpr(fun, arg)
	return call
}

func mapOperator(op string) token.Token {
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

func typeDecl(name *goast.Ident, typ goast.Expr) *goast.GenDecl {
	return &goast.GenDecl{
		Tok: token.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: name,
				Type: typ,
			},
		},
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

func indexExpr(x goast.Expr, idx *goast.Ident) goast.Expr {
	return &goast.IndexExpr{
		X:      x,
		Lbrack: token.Pos(1),
		Index:  idx,
		Rbrack: token.Pos(1),
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

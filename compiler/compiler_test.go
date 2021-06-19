package compiler

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/stretchr/testify/assert"
)

func TestCompilerCompile(t *testing.T) {
	type test struct {
		label  string
		source ast.Node
		exp    []Ins
	}

	tests := []test{
		{
			label: "while statement",
			source: &ast.SWhile{
				Condition: &ast.ELiteral{
					Type:   ast.LitTypeIdent,
					Lexeme: "true",
				},
				Body: &ast.SExpr{
					Expr: &ast.EAssign{
						Ident: "val",
						Expr: &ast.ELiteral{
							Type:   ast.LitTypeNumber,
							Lexeme: "123",
						},
					},
				},
			},
			exp: []Ins{
				{op.LoadName, "true"},
				{op.RelJumpIfFalse, 3},
				{op.LoadConst, float64(123)},
				{op.StoreName, "val"},
				{op.Jump, 0},
			},
		},
		{
			label: "if statement",
			source: &ast.SIf{
				Condition: &ast.EGroup{
					Expr: &ast.EBinary{
						LHS: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "1"},
						Op:  ast.BinOpAdd,
						RHS: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "1"},
					},
				},
				TrueBlock: &ast.SBlock{
					Stmts: []ast.Stmt{
						&ast.SExpr{
							Expr: &ast.EAssign{
								Ident: "x",
								Expr:  &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "1"},
							},
						},
					},
				},
				FalseBlock: &ast.SExpr{
					Expr: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "0"},
				},
			},
			exp: []Ins{
				{op.LoadConst, float64(1)},
				{op.LoadConst, float64(1)},
				{op.BinAdd, nil},
				{op.RelJumpIfFalse, 3},
				{op.LoadConst, float64(1)},
				{op.StoreName, "x"},
				{op.RelJump, 1},
				{op.LoadConst, float64(0)},
			},
		},
		{
			label: "func def",
			source: &ast.SBlock{
				Stmts: []ast.Stmt{
					&ast.SExpr{
						Expr: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "false"},
					},
					&ast.SExpr{
						Expr: &ast.EFuncDef{
							Args: []string{"name"},
							Body: &ast.SExpr{
								Expr: &ast.EAssign{
									Ident: "x",
									Expr:  &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "true"},
								},
							},
						},
					},
				},
			},
			exp: []Ins{
				{op.LoadName, "false"},
				{
					Op: op.MakeFunction,
					Arg: &FuncDef{
						Args: []string{"name"},
						Body: []Ins{
							{op.LoadName, "true"},
							{op.StoreName, "x"},
						},
					},
				},
			},
		},
		{
			label: "return",
			source: &ast.SReturn{
				Expr: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "123.1"},
			},
			exp: []Ins{
				{op.LoadConst, 123.1},
				{op.Return, nil},
			},
		},
		{
			label: "attribute get",
			source: &ast.SExpr{
				Expr: &ast.EAttrGet{
					Attr: "name",
					Expr: &ast.ELiteral{
						Type:   ast.LitTypeIdent,
						Lexeme: "x",
					},
				},
			},
			exp: []Ins{
				{op.LoadName, "x"},
				{op.AttrGet, "name"},
			},
		},
		{
			label: "call",
			source: &ast.SExpr{
				Expr: &ast.ECall{
					Expr: &ast.ELiteral{
						Type:   ast.LitTypeIdent,
						Lexeme: "f",
					},
					Args: []ast.Expr{
						&ast.ELiteral{
							Type:   ast.LitTypeNumber,
							Lexeme: "123",
						},
						&ast.ELiteral{
							Type:   ast.LitTypeIdent,
							Lexeme: "x",
						},
					},
				},
			},
			exp: []Ins{
				{op.LoadConst, 123.0},
				{op.LoadName, "x"},
				{op.LoadName, "f"},
				{op.Call, 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			var compiler Compiler

			assert := assert.New(t)

			got := compiler.Compile(ast.New(test.source))
			assert.Equal(test.exp, got)
		})
	}
}

func TestCompilerVisitEBinaryPanics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		var compiler Compiler

		compiler.VisitEBinary(&ast.EBinary{
			LHS: &ast.EAssign{},
			Op:  -1,
			RHS: &ast.EAssign{},
		})
	})
}

package parser

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/stretchr/testify/assert"
)

func TestParserParse(t *testing.T) {
	type test struct {
		source string
		exp    ast.Stmt
	}

	tests := []test{
		{
			source: "!(1 == true)",
			exp: &ast.SExpr{
				Expr: &ast.EUnary{
					Op: ast.UnaryOpNot,
					Expr: &ast.EGroup{
						Expr: &ast.EBinary{
							LHS: &ast.ELiteral{
								Type:   ast.LitTypeNumber,
								Lexeme: "1",
							},
							Op: ast.BinOpEq,
							RHS: &ast.ELiteral{
								Type:   ast.LitTypeIdent,
								Lexeme: "true",
							},
						},
					},
				},
			},
		},
		{
			source: "1+1",
			exp: &ast.SExpr{
				Expr: &ast.EBinary{
					LHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "1",
					},
					Op: ast.BinOpAdd,
					RHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "1",
					},
				},
			},
		},
		{
			source: "1 / 2",
			exp: &ast.SExpr{
				Expr: &ast.EBinary{
					LHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "1",
					},
					Op: ast.BinOpDiv,
					RHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "2",
					},
				},
			},
		},
		{
			source: "500 >= 10124",
			exp: &ast.SExpr{
				Expr: &ast.EBinary{
					LHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "500",
					},
					Op: ast.BinOpGreaterEq,
					RHS: &ast.ELiteral{
						Type:   ast.LitTypeNumber,
						Lexeme: "10124",
					},
				},
			},
		},
		{
			source: "x = true;",
			exp: &ast.SAssign{
				Ident: "x",
				Expr: &ast.ELiteral{
					Type:   ast.LitTypeIdent,
					Lexeme: "true",
				},
			},
		},
		{
			source: "func (x, y, z) x + y + z",
			exp: &ast.SExpr{
				Expr: &ast.EFuncDef{
					Args: []string{"x", "y", "z"},
					Body: &ast.SExpr{
						Expr: &ast.EBinary{
							LHS: &ast.EBinary{
								LHS: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "x"},
								Op:  ast.BinOpAdd,
								RHS: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "y"},
							},
							Op:  ast.BinOpAdd,
							RHS: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "z"},
						},
					},
				},
			},
		},
		{
			source: "func !1",
			exp: &ast.SExpr{
				Expr: &ast.EFuncDef{
					Body: &ast.SExpr{
						Expr: &ast.EUnary{
							Op: ast.UnaryOpNot,
							Expr: &ast.ELiteral{
								Type:   ast.LitTypeNumber,
								Lexeme: "1",
							},
						},
					},
				},
			},
		},
		{
			source: "if x 1 else 2",
			exp: &ast.SIf{
				Condition: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "x"},
				TrueBlock: &ast.SExpr{
					Expr: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "1"},
				},
				FalseBlock: &ast.SExpr{
					Expr: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "2"},
				},
			},
		},
		{
			source: "while (true) { 500 }",
			exp: &ast.SWhile{
				Condition: &ast.EGroup{
					Expr: &ast.ELiteral{Type: ast.LitTypeIdent, Lexeme: "true"},
				},
				Body: &ast.SBlock{
					Stmts: []ast.Stmt{
						&ast.SExpr{
							Expr: &ast.ELiteral{Type: ast.LitTypeNumber, Lexeme: "500"},
						}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.source, func(t *testing.T) {
			assert := assert.New(t)

			parser := New(Tokenize(test.source))

			exp := ast.New(&ast.SBlock{Stmts: []ast.Stmt{test.exp}})
			got := parser.Parse()

			assert.Equal(exp, got)
		})
	}

}

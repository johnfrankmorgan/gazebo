package parser

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/stretchr/testify/assert"
)

func TestParserParse(t *testing.T) {
	assert := assert.New(t)

	type test struct {
		source string
		exp    *ast.AST
	}

	tests := []test{
		{
			source: "!(1 == true)",
			exp: ast.New(&ast.EUnary{
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
			}),
		},
		{
			source: "1+1",
			exp: ast.New(&ast.EBinary{
				LHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "1",
				},
				Op: ast.BinOpAdd,
				RHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "1",
				},
			}),
		},
		{
			source: "1/2",
			exp: ast.New(&ast.EBinary{
				LHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "1",
				},
				Op: ast.BinOpDiv,
				RHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "2",
				},
			}),
		},
		{
			source: "500 >= 10124",
			exp: ast.New(&ast.EBinary{
				LHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "500",
				},
				Op: ast.BinOpGreaterEq,
				RHS: &ast.ELiteral{
					Type:   ast.LitTypeNumber,
					Lexeme: "10124",
				},
			}),
		},
	}

	for _, test := range tests {
		parser := New(Tokenize(test.source))
		got := parser.Parse()

		assert.Equal(test.exp, got)
	}

}

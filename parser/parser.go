package parser

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast"
)

type Parser struct {
	ts tstream
}

func New(tokens []Token) *Parser {
	return &Parser{ts: tstream{tokens: tokens}}
}

func (m *Parser) Parse() *ast.AST {
	node := &ast.SBlock{}

	for !m.ts.finished() {
		node.Append(m.parse())

		for m.ts.match(TSemicolon) {
			// consume semicolons following statements
		}
	}

	return ast.New(node)
}

func (m *Parser) parse() ast.Node {
	return m.statement()
}

func (m *Parser) statement() ast.Stmt {
	if m.ts.check(TIdent) && m.ts.peek(1).Is(TEqual) {
		return m.assignment()
	}

	return &ast.SExpr{Expr: m.expression()}
}

func (m *Parser) assignment() ast.Stmt {
	ident := m.ts.consume(TIdent).lexeme
	m.ts.consume(TEqual)
	return &ast.SAssign{Ident: ident, Expr: m.expression()}
}

func (m *Parser) expression() ast.Expr {
	return m.equality()
}

func (m *Parser) equality() ast.Expr {
	expr := m.comparison()

	for m.ts.match(TEqualEqual, TBangEqual) {
		op := m.ts.prev()
		rhs := m.comparison()
		expr = &ast.EBinary{
			LHS: expr,
			Op:  op.ToBinOp(),
			RHS: rhs,
		}
	}

	return expr
}

func (m *Parser) comparison() ast.Expr {
	expr := m.addition()

	for m.ts.match(TGreater, TGreaterEqual, TLess, TLessEqual) {
		op := m.ts.prev()
		rhs := m.addition()
		expr = &ast.EBinary{
			LHS: expr,
			Op:  op.ToBinOp(),
			RHS: rhs,
		}
	}

	return expr
}

func (m *Parser) addition() ast.Expr {
	expr := m.multiplication()

	for m.ts.match(TPlus, TMinus) {
		op := m.ts.prev()
		rhs := m.multiplication()
		expr = &ast.EBinary{
			LHS: expr,
			Op:  op.ToBinOp(),
			RHS: rhs,
		}
	}

	return expr
}

func (m *Parser) multiplication() ast.Expr {
	expr := m.unary()

	for m.ts.match(TStar, TSlash) {
		op := m.ts.prev()
		rhs := m.unary()
		expr = &ast.EBinary{
			LHS: expr,
			Op:  op.ToBinOp(),
			RHS: rhs,
		}
	}

	return expr
}

func (m *Parser) unary() ast.Expr {
	if m.ts.match(TBang, TMinus) {
		op := m.ts.prev()
		expr := m.unary()
		return &ast.EUnary{
			Op:   op.ToUnaryOp(),
			Expr: expr,
		}
	}

	return m.literal()
}

func (m *Parser) literal() ast.Expr {
	token := m.ts.next()

	switch token.kind {
	case TIdent:
		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeIdent}

	case TNumber:
		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeNumber}

	case TString:
		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeString}
	}

	if token.Is(TParenOpen) {
		expr := m.expression()
		m.ts.consume(TParenClose)
		return &ast.EGroup{Expr: expr}
	}

	panic(
		fmt.Errorf(
			"token %s is not a literal",
			token.kind,
		),
	)
}

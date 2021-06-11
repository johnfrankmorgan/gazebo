package parser

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast"
)

type Parser struct {
	tokens   []Token
	position int
}

func New(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (m *Parser) Parse() *ast.AST {
	root := m.parse()

	return ast.New(root)
}

func (m *Parser) finished() bool {
	return m.position >= len(m.tokens)
}

func (m *Parser) advance() {
	m.position++
}

func (m *Parser) peek(n int) Token {
	token := m.tokens[m.position+n] // FIXME: add bounds check

	if token.Is(TWhitespace, TComment) {
		m.advance()
		return m.peek(n)
	}

	return token
}

func (m *Parser) prev() Token {
	return m.peek(-1)
}

func (m *Parser) next() Token {
	defer m.advance()

	return m.peek(0)
}

func (m *Parser) check(kinds ...TKind) bool {
	if m.finished() {
		return false
	}

	return m.peek(0).Is(kinds...)
}

func (m *Parser) match(kinds ...TKind) bool {
	if m.check(kinds...) {
		m.next()
		return true
	}

	return false
}

func (m *Parser) consume(kinds ...TKind) {
	if m.match(kinds...) {
		return
	}

	panic(
		fmt.Errorf(
			"expected one of %s, got %s",
			kinds,
			m.peek(0).kind,
		),
	)
}

func (m *Parser) parse() ast.Node {
	return m.expression()
}

func (m *Parser) expression() ast.Expr {
	return m.equality()
}

func (m *Parser) equality() ast.Expr {
	expr := m.comparison()

	for m.match(TEqualEqual, TBangEqual) {
		op := m.prev()
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

	for m.match(TGreater, TGreaterEqual, TLess, TLessEqual) {
		op := m.prev()
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

	for m.match(TPlus, TMinus) {
		op := m.prev()
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

	for m.match(TStar, TSlash) {
		op := m.prev()
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
	if m.match(TBang, TMinus) {
		op := m.prev()
		expr := m.unary()
		return &ast.EUnary{
			Op:   op.ToUnaryOp(),
			Expr: expr,
		}
	}

	return m.literal()
}

func (m *Parser) literal() ast.Expr {
	token := m.next()

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
		m.consume(TParenClose)
		return &ast.EGroup{Expr: expr}
	}

	panic(
		fmt.Errorf(
			"token %s is not a literal",
			token.kind,
		),
	)
}

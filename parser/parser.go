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
	}

	return ast.New(node)
}

func (m *Parser) parse() ast.Node {
	return m.statement()
}

func (m *Parser) statement() ast.Stmt {
	switch m.ts.peek(0).kind {
	case TBraceOpen:
		return m.block()

	case TIf:
		return m.conditional()

	case TWhile:
		return m.while()
	}

	defer m.ts.terminate()

	if m.ts.match(TReturn) {
		// FIXME: we should allow empty returns
		return &ast.SReturn{Expr: m.expression()}
	}

	return &ast.SExpr{Expr: m.expression()}
}

func (m *Parser) block() ast.Stmt {
	var block ast.SBlock

	m.ts.consume(TBraceOpen)

	for !m.ts.match(TBraceClose) {
		block.Append(m.statement())
	}

	return &block
}

func (m *Parser) conditional() ast.Stmt {
	var stmt ast.SIf

	m.ts.consume(TIf)

	stmt.Condition = m.expression()
	stmt.TrueBlock = m.statement()

	if m.ts.match(TElse) {
		stmt.FalseBlock = m.statement()
	}

	return &stmt
}

func (m *Parser) while() ast.Stmt {
	var stmt ast.SWhile

	m.ts.consume(TWhile)

	stmt.Condition = m.expression()
	stmt.Body = m.statement()

	return &stmt
}

func (m *Parser) expression() ast.Expr {
	expr := m.equality()

	for m.ts.check(TParenOpen, TDot) {
		if m.ts.match(TParenOpen) {
			args := []ast.Expr{}

			for !m.ts.check(TParenClose) {
				args = append(args, m.expression())

				// FIXME: allow trailing commas on function calls
				if !m.ts.check(TParenClose) {
					m.ts.consume(TComma)
				}
			}

			m.ts.consume(TParenClose)
			expr = &ast.ECall{
				Expr: expr,
				Args: args,
			}
		}

		if m.ts.match(TDot) {
			m.ts.consume(TIdent)

			expr = &ast.EAttrGet{
				Expr: expr,
				Attr: m.ts.prev().lexeme,
			}
		}
	}

	if expr, ok := expr.(*ast.EAttrGet); ok && m.ts.match(TEqual) {
		return &ast.EAttrSet{
			Expr:  expr.Expr,
			Attr:  expr.Attr,
			Value: m.expression(),
		}
	}

	return expr
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
		if m.ts.match(TEqual) {
			return &ast.EAssign{Ident: token.lexeme, Expr: m.expression()}
		}

		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeIdent}

	case TNumber:
		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeNumber}

	case TString:
		return &ast.ELiteral{Lexeme: token.lexeme, Type: ast.LitTypeString}

	case TFunc:
		return m.funcdef()
	}

	if token.Is(TParenOpen) {
		expr := m.expression()
		m.ts.consume(TParenClose)
		return &ast.EGroup{Expr: expr}
	}

	panic(
		fmt.Errorf(
			"token %q is not a literal",
			token.kind,
		),
	)
}

func (m *Parser) funcdef() ast.Expr {
	expr := &ast.EFuncDef{}

	// FIXME: this fails to parse function definitions with parentheses but no arguments
	//        func { ... } and func () { ... } should both be acceptable
	if m.ts.match(TParenOpen) {
		m.ts.consume(TIdent)
		expr.AddArg(m.ts.prev().lexeme)

		for m.ts.match(TComma) {
			m.ts.consume(TIdent)
			expr.AddArg(m.ts.prev().lexeme)
		}

		m.ts.consume(TParenClose)
	}

	expr.Body = m.statement()
	return expr
}

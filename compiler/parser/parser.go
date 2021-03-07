package parser

import (
	"strconv"

	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/stmt"
)

type Parser struct {
	stream *TokenStream
}

func New(tokens lexer.Tokens) *Parser {
	return &Parser{
		stream: NewTokenStream(tokens),
	}
}

func (m *Parser) Parse() (st stmt.Statement, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			if perr, ok := perr.(error); ok {
				err = perr
			}
		}
	}()

	var statements []stmt.Statement

	m.stream.Reset()

	for !m.stream.Finished() {
		if m.stream.Peek().Is(lexer.TkEOF) {
			break
		}

		statements = append(statements, m.parse())

		for m.stream.Match(lexer.TkSemicolon) {
			//
		}
	}

	return &stmt.Block{Statements: statements}, nil
}

func (m *Parser) parse() stmt.Statement {
	semicolon := func() {
		if m.stream.Check(lexer.TkBraceOpen) {
			return
		}

		m.stream.Expect(lexer.TkSemicolon)
	}

	switch token := m.stream.Peek(); token.Type {
	case lexer.TkEOF:
		panic(UnexpectedEOF())

	case lexer.TkIdent:
		if m.stream.Peek(1).Is(lexer.TkEqual) {
			defer semicolon()
			m.stream.Advance()
			m.stream.Advance()
			return &stmt.Assignment{
				Name: token.Value,
				Expr: m.expression(),
			}
		}

	case lexer.TkPass:
		defer semicolon()
		m.stream.Advance()
		return &stmt.Pass{}

	case lexer.TkReturn:
		defer semicolon()
		m.stream.Advance()
		return &stmt.Return{
			Expr: m.expression(),
		}

	case lexer.TkBraceOpen:
		if m.stream.Peek(2).Is(lexer.TkColon) {
			break
		}
		stmt := &stmt.Block{}
		m.stream.Advance()
		for !m.stream.Finished() {
			if m.stream.Match(lexer.TkBraceClose) {
				return stmt
			}
			stmt.Statements = append(stmt.Statements, m.parse())
		}
		panic(UnexpectedEOF())

	case lexer.TkDel:
		defer semicolon()
		return m.delstmt()

	case lexer.TkIf:
		m.stream.Advance()
		stmt := &stmt.If{
			Condition: m.expression(),
			TruePath:  m.parse(),
		}
		if m.stream.Match(lexer.TkElse) {
			stmt.FalsePath = m.parse()
		}
		return stmt

	case lexer.TkWhile:
		m.stream.Advance()
		return &stmt.While{
			Condition: m.expression(),
			Body:      m.parse(),
		}

	case lexer.TkFor:
		return m.forstmt()

	case lexer.TkLoad:
		defer semicolon()
		m.stream.Advance()
		stmt := &stmt.Load{}
		for m.stream.Match(lexer.TkIdent) {
			stmt.Names = append(stmt.Names, m.stream.Prev().Value)
		}
		return stmt

	case lexer.TkBreak:
		defer semicolon()
		m.stream.Advance()
		return &stmt.Break{}

	case lexer.TkContinue:
		defer semicolon()
		m.stream.Advance()
		return &stmt.Continue{}
	}

	defer semicolon()

	ex := m.expression()

	if getattr, ok := ex.(*expr.GetAttr); ok {
		if m.stream.Match(lexer.TkEqual) {
			return &stmt.SetAttr{
				Expr:  getattr.Expr,
				Name:  getattr.Name,
				Value: m.expression(),
			}
		}
	}

	return &stmt.Expression{
		Expr: ex,
	}
}

func (m *Parser) delstmt() stmt.Statement {
	m.stream.Expect(lexer.TkDel)

	ex := m.expression()

	switch ex := ex.(type) {
	case *expr.Literal:
		if !ex.Token.Is(lexer.TkIdent) {
			panic(UnexpectedToken(ex.Token, lexer.TkIdent))
		}

	case *expr.GetAttr:
		return &stmt.DelAttr{
			Expr: ex.Expr,
			Name: ex.Name,
		}
	}

	panic(UnexpectedExpression(ex))
}

func (m *Parser) forstmt() stmt.Statement {
	m.stream.Expect(lexer.TkFor)

	init := m.parse()
	cond := m.expression()
	m.stream.Expect(lexer.TkSemicolon)
	incr := m.parse()
	body := m.parse()

	return &stmt.Block{
		Statements: []stmt.Statement{
			init,
			&stmt.While{
				Condition: cond,
				Body: &stmt.Block{
					Statements: []stmt.Statement{
						body,
						incr,
					},
				},
			},
		},
	}
}

func (m *Parser) binary(next func() expr.Expression, expected ...lexer.TokenType) expr.Expression {
	ex := next()

	for m.stream.Match(expected...) {
		ex = &expr.Binary{
			Op:    m.stream.Prev(),
			Left:  ex,
			Right: next(),
		}
	}

	return ex
}

func (m *Parser) expression() expr.Expression {
	return m.logical()
}

func (m *Parser) logical() expr.Expression {
	return m.binary(m.contains, lexer.TkAnd, lexer.TkOr)
}

func (m *Parser) contains() expr.Expression {
	ex := m.equality()

	if m.stream.Match(lexer.TkIn) {
		ex = &expr.Binary{
			Op:    m.stream.Prev(),
			Left:  m.expression(),
			Right: ex,
		}
	}

	return ex
}

func (m *Parser) equality() expr.Expression {
	return m.binary(m.comparison, lexer.TkEqualEqual, lexer.TkBangEqual)
}

func (m *Parser) comparison() expr.Expression {
	return m.binary(
		m.addition,
		lexer.TkGreater,
		lexer.TkGreaterEqual,
		lexer.TkLess,
		lexer.TkLessEqual,
	)
}

func (m *Parser) addition() expr.Expression {
	return m.binary(m.multiplication, lexer.TkPlus, lexer.TkMinus)
}

func (m *Parser) multiplication() expr.Expression {
	return m.binary(m.unary, lexer.TkStar, lexer.TkSlash)
}

func (m *Parser) unary() expr.Expression {
	if m.stream.Match(lexer.TkBang, lexer.TkMinus) {
		return &expr.Unary{
			Op:    m.stream.Prev(),
			Right: m.unary(),
		}
	}

	return m.funcall()
}

func (m *Parser) funcall() expr.Expression {
	ex := m.fundef()

	for m.stream.Check(lexer.TkParenOpen, lexer.TkDot) {
		if m.stream.Match(lexer.TkParenOpen) {
			funcall := &expr.FunCall{
				Function: ex,
			}

			for !m.stream.Check(lexer.TkParenClose) {
				funcall.Args = append(funcall.Args, m.expression())

				if !m.stream.Check(lexer.TkParenClose) {
					m.stream.Expect(lexer.TkComma)
				}
			}

			m.stream.Expect(lexer.TkParenClose)
			ex = funcall
		}

		if m.stream.Match(lexer.TkDot) {
			ex = &expr.GetAttr{
				Name: m.stream.Expect(lexer.TkIdent).Value,
				Expr: ex,
			}
		}
	}

	return ex

}

func (m *Parser) fundef() expr.Expression {
	if !m.stream.Match(lexer.TkFun) {
		return m.primary()
	}

	ex := &expr.Fun{}

	if m.stream.Match(lexer.TkParenOpen) {
		for !m.stream.Finished() {
			if m.stream.Match(lexer.TkParenClose) {
				break
			}

			arg := m.stream.Expect(lexer.TkIdent)
			ex.Args = append(ex.Args, arg.Value)

			if !m.stream.Check(lexer.TkParenClose) {
				m.stream.Expect(lexer.TkComma)
			}
		}
	}

	ex.Body = m.parse()

	return ex
}

func (m *Parser) primary() expr.Expression {
	if m.stream.Match(lexer.TkBracketOpen) {
		ex := &expr.List{}

		for !m.stream.Finished() {
			if m.stream.Check(lexer.TkBracketClose) {
				break
			}

			ex.Values = append(ex.Values, m.expression())

			if !m.stream.Check(lexer.TkBracketClose) {
				m.stream.Expect(lexer.TkComma)
			}
		}

		m.stream.Expect(lexer.TkBracketClose)

		return ex
	}

	if m.stream.Match(lexer.TkBraceOpen) {
		ex := &expr.Map{}

		for !m.stream.Finished() {
			if m.stream.Check(lexer.TkBraceClose) {
				break
			}

			key := m.stream.Expect(lexer.TkIdent, lexer.TkString, lexer.TkNumber)

			if key.Is(lexer.TkIdent) {
				key.Value = strconv.Quote(key.Value)
				key.Type = lexer.TkString
			}

			m.stream.Expect(lexer.TkColon)

			value := m.expression()

			ex.Keys = append(ex.Keys, &expr.Literal{Token: key})
			ex.Values = append(ex.Values, value)

			if !m.stream.Check(lexer.TkBraceClose) {
				m.stream.Expect(lexer.TkComma)
			}
		}

		m.stream.Expect(lexer.TkBraceClose)

		return ex
	}

	if m.stream.Match(lexer.TkIdent, lexer.TkString, lexer.TkNumber) {
		return &expr.Literal{
			Token: m.stream.Prev(),
		}
	}

	if m.stream.Match(lexer.TkParenOpen) {
		defer m.stream.Expect(lexer.TkParenClose)
		return &expr.Group{
			Expr: m.expression(),
		}
	}

	panic(UnexpectedToken(
		m.stream.Peek(),
		lexer.TkBracketOpen,
		lexer.TkIdent,
		lexer.TkString,
		lexer.TkNumber,
		lexer.TkParenOpen,
	))
}

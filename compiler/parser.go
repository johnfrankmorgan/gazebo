package compiler

import (
	"strings"

	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
)

func parse(source string) []statement {
	tokens := tokenize(source)

	if debug.Enabled() {
		tokens.dump()
	}

	parser := parser{tokens: tokens}

	return parser.parse()
}

type parser struct {
	tokens   tokens
	position int
}

func (m *parser) unexpectedeof() {
	errors.ErrEOF.Panic(
		"unexpected eof at token offset %d, token %#v",
		m.position,
		m.peek(),
	)
}

func (m *parser) finished() bool {
	return m.peek().is(tkeof)
}

func (m *parser) peek(pos ...int) token {
	if len(pos) == 0 {
		pos = []int{0}
	}

	if m.position+pos[0] >= len(m.tokens) {
		return token{typ: tkinvalid}
	}

	return m.tokens[m.position+pos[0]]
}

func (m *parser) prev() token {
	return m.tokens[m.position-1]
}

func (m *parser) next() token {
	token := m.tokens[m.position]
	m.position++
	return token
}

func (m *parser) check(typ ...tokentype) bool {
	if m.finished() {
		return false
	}

	token := m.peek()

	for _, typ := range typ {
		if token.is(typ) {
			return true
		}
	}

	return false
}

func (m *parser) match(typ ...tokentype) bool {
	if m.check(typ...) {
		m.next()
		return true
	}

	return false
}

func (m *parser) expect(typ ...tokentype) token {
	if !m.match(typ...) {
		names := make([]string, len(typ))

		for i, typ := range typ {
			names[i] = typ.name()
		}

		errors.ErrParse.Panic(
			"expected one of (%s), got %s",
			strings.Join(names, ", "),
			m.peek().typ.name(),
		)
	}

	return m.prev()
}

func (m *parser) expression() expression {
	return m.equality()
}

func (m *parser) binary(next func() expression, expected ...tokentype) expression {
	expr := next()

	for m.match(expected...) {
		expr = &binary{
			op:    m.prev(),
			left:  expr,
			right: next(),
		}
	}

	return expr
}

func (m *parser) equality() expression {
	return m.binary(m.comparison, tkequalequal, tkbangequal)
}

func (m *parser) comparison() expression {
	return m.binary(m.addition, tkgreater, tkgreaterequal, tkless, tklessequal)
}

func (m *parser) addition() expression {
	return m.binary(m.multiplication, tkplus, tkminus)
}

func (m *parser) multiplication() expression {
	return m.binary(m.unary, tkstar, tkslash)
}

func (m *parser) unary() expression {
	if m.match(tkbang, tkminus) {
		return &unary{op: m.prev(), right: m.unary()}
	}

	return m.primary()
}

func (m *parser) primary() expression {
	if m.match(tkident, tkstring, tknumber) {
		return &literal{token: m.prev()}
	}

	if m.match(tkparenopen) {
		expr := m.expression()
		m.expect(tkparenclose)
		return &group{expr: expr}
	}

	errors.ErrParse.Panic(
		"unexpected %s %s",
		m.peek().typ.name(),
		m.peek().value,
	)
	return nil
}

func (m *parser) statement() statement {
	var stmt statement

	switch m.peek().typ {
	case tkeof:
		m.unexpectedeof()

	case tkbraceopen:
		stmt = m.block()

	case tklet:
		stmt = m.assignment()

	default:
		stmt = &exprstmt{expr: m.expression()}
	}

	m.expect(tksemicolon, tknewline)
	return stmt
}

func (m *parser) block() statement {
	var statements []statement

	m.expect(tkbraceopen)

	for !m.finished() {
		if m.match(tknewline) {
			continue
		}

		if m.match(tkbraceclose) {
			return &block{
				statements: statements,
			}
		}

		statements = append(statements, m.statement())
	}

	m.unexpectedeof()
	return nil
}

func (m *parser) assignment() statement {
	m.expect(tklet)

	name := m.expect(tkident)

	m.expect(tkequal)

	return &assign{
		name: name,
		expr: m.expression(),
	}
}

func (m *parser) parse() []statement {
	statements := []statement{}

	for !m.finished() {
		statements = append(statements, m.statement())
	}

	return statements
}

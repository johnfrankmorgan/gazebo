package compiler

import (
	"strings"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
)

func parse(source string) expression {
	tokens := tokenize(source)

	if debug.Enabled() {
		tokens.dump()
	}

	parser := parser{tokens: tokens}

	return parser.parse()
}

func dumpexpression(expr expression, depth int) {
	indent := strings.Repeat("  ", depth)

	limit := func(str string, count int) string {
		if len(str) > count {
			return str[:count]
		}

		return str
	}

	switch expr := expr.(type) {
	case *binary:
		debug.Printf("BIN  %s%s %s {\n", indent, expr.op.value, expr.op.typ.name())
		dumpexpression(expr.left, depth+1)
		dumpexpression(expr.right, depth+1)
		debug.Printf("/BIN %s}\n", indent)

	case *unary:
		debug.Printf("UNY  %s%s %s {\n", indent, expr.op.value, expr.op.typ.name())
		dumpexpression(expr.right, depth+1)
		debug.Printf("/UNY %s}\n", indent)

	case *literal:
		debug.Printf("LIT  %s%s %s\n", indent, limit(expr.token.value, 10), expr.token.typ.name())

	case *group:
		debug.Printf("GRP  %s() {\n", indent)
		dumpexpression(expr.expr, depth+1)
		debug.Printf("/GRP %s}\n", indent)

	default:
		assert.Unreached("unknown expression type %T", expr)
	}
}

type parser struct {
	tokens   tokens
	position int
}

func (m *parser) unexpectedeof() expression {
	errors.ErrEOF.Panic(
		"unexpected eof at token offset %d, token %#v",
		m.position,
		m.peek(),
	)
	return nil
}

func (m *parser) finished() bool {
	return m.peek().is(tkeof)
}

func (m *parser) peek() token {
	return m.tokens[m.position]
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

func (m *parser) expect(typ ...tokentype) {
	if !m.match(typ...) {
		names := make([]string, len(typ))

		for i, typ := range typ {
			names[i] = typ.name()
		}

		errors.ErrParse.Panic(
			"expected %s, got %s",
			strings.Join(names, ", "),
			m.peek().typ.name(),
		)
	}
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

func (m *parser) parse() expression {
	return m.expression()
}

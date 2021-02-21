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

	return m.funcall()
}

func (m *parser) funcall() expression {
	expr := m.fundef()

	for !m.finished() {
		if !m.match(tkparenopen) {
			break
		}

		funcall := &funcall{name: expr}
		expr = funcall

		for !m.finished() {
			if m.match(tkparenclose) {
				break
			}

			funcall.args = append(funcall.args, m.expression())

			if !m.check(tkparenclose) {
				m.expect(tkcomma)
			}
		}

		if m.match(tkparenclose) {
			break
		}
	}

	return expr
}

func (m *parser) fundef() expression {
	if !m.match(tkfun) {
		return m.primary()
	}

	fundef := &fundef{args: []string{}}

	if m.match(tkparenopen) {
		for !m.finished() {
			if m.match(tkparenclose) {
				break
			}

			arg := m.expect(tkident)
			fundef.args = append(fundef.args, arg.value)

			if !m.check(tkparenclose) {
				m.expect(tkcomma)
			}
		}
	}

	fundef.body = m.statement()
	return fundef
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
	switch m.peek().typ {
	case tkeof:
		m.unexpectedeof()
		return nil

	case tkpass:
		m.next()
		return &pass{}

	case tkreturn:
		m.next()
		return &returnstmt{expr: m.expression()}

	case tkbraceopen:
		return m.block()

	case tklet:
		return m.assignment()

	case tkif:
		return m.ifstatement()

	case tkwhile:
		return m.while()

	case tkload:
		return m.load()
	}

	stmt := &exprstmt{expr: m.expression()}

	m.expect(tknewline, tksemicolon)

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

	expr := m.expression()

	m.expect(tknewline, tksemicolon)

	return &assign{
		name: name,
		expr: expr,
	}
}

func (m *parser) ifstatement() statement {
	var falsestmt statement

	m.expect(tkif)

	condition := m.expression()
	truestmt := m.statement()

	if m.match(tkelse) {
		falsestmt = m.statement()
	}

	return &ifstmt{condition: condition, truestmt: truestmt, falsestmt: falsestmt}
}

func (m *parser) while() statement {
	m.expect(tkwhile)

	condition := m.expression()
	body := m.statement()

	return &while{
		condition: condition,
		body:      body,
	}
}

func (m *parser) load() statement {
	var modules []string

	m.expect(tkload)

	for !m.finished() {
		if !m.match(tkident) {
			break
		}

		modules = append(modules, m.prev().value)
	}

	return &load{modules: modules}
}

func (m *parser) parse() []statement {
	statements := []statement{}

	for !m.finished() {
		if m.match(tknewline) {
			continue
		}

		statements = append(statements, m.statement())
	}

	return statements
}

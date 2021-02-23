package compiler

import (
	"strings"

	"github.com/johnfrankmorgan/gazebo/errors"
)

type parser struct {
	tokens   tokens
	position int
}

func (m *parser) reset() {
	m.position = 0
}

func (m *parser) desugar() {
	defer m.reset()

	tokens := tokens{}

	for !m.finished() {
		tk := m.peek()

		if tk.is(tkfun) && m.peek(1).is(tkident) {
			// rewrite function definition statement into assignment

			m.next()
			name := m.expect(tkident)

			tokens = append(tokens, name, token{typ: tkequal, value: "="}, tk)

			continue
		}

		if !tk.is(tkplusequal, tkminusequal, tkstarequal, tkslashequal) {
			tokens = append(tokens, m.next())
			continue
		}

		// rewrite compound assignment into simple assignment
		// i += 1  ->  i = i + 1

		prev := m.prev()
		m.next()

		errors.ErrParse.Expect(
			prev.is(tkident),
			"%s must be preceded by tkident",
			tk.typ.name(),
		)

		ops := map[tokentype]string{
			tkplus:  "+",
			tkminus: "-",
			tkstar:  "*",
			tkslash: "/",
		}

		tokens = append(
			tokens,
			token{typ: tkequal, value: "="},
			prev,
			token{typ: tk.typ - 1, value: ops[tk.typ-1]},
		)
	}

	tokens = append(tokens, token{typ: tkeof})

	m.tokens = tokens
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
		expr = &exprbinary{
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
		return &exprunary{op: m.prev(), right: m.unary()}
	}

	return m.funcall()
}

func (m *parser) funcall() expression {
	expr := m.fundef()

	for m.check(tkparenopen, tkdot) {
		if m.match(tkparenopen) {
			funcall := &exprfuncall{name: expr}

			for !m.check(tkparenclose) {
				funcall.args = append(funcall.args, m.expression())
				if !m.check(tkparenclose) {
					m.expect(tkcomma)
				}
			}

			m.expect(tkparenclose)
			expr = funcall
		}

		if m.match(tkdot) {
			expr = &exprgetattr{
				name: m.expect(tkident).value,
				expr: expr,
			}
		}
	}

	return expr
}

func (m *parser) fundef() expression {
	if !m.match(tkfun) {
		return m.primary()
	}

	expr := &exprfun{args: []string{}}

	if m.match(tkparenopen) {
		for !m.finished() {
			if m.match(tkparenclose) {
				break
			}

			arg := m.expect(tkident)
			expr.args = append(expr.args, arg.value)

			if !m.check(tkparenclose) {
				m.expect(tkcomma)
			}
		}
	}

	expr.body = m.statement()

	return expr
}

func (m *parser) primary() expression {
	if m.match(tkbracketopen) {
		return m.list()
	}

	if m.match(tkident, tkstring, tknumber) {
		return &exprliteral{token: m.prev()}
	}

	if m.match(tkparenopen) {
		expr := m.expression()
		m.expect(tkparenclose)
		return &exprgroup{expr: expr}
	}

	errors.ErrParse.Panic(
		"unexpected %s %s near token offset %d",
		m.peek().typ.name(),
		m.peek().value,
		m.position,
	)
	return nil
}

func (m *parser) list() expression {
	expr := &exprlist{}

	for !m.finished() {
		if m.check(tkbracketclose) {
			break
		}

		expr.expressions = append(expr.expressions, m.expression())

		if !m.check(tkbracketclose) {
			m.expect(tkcomma)
		}
	}

	m.expect(tkbracketclose)

	return expr
}

func (m *parser) statement() statement {
	switch m.peek().typ {
	case tkeof:
		m.unexpectedeof()
		return nil

	case tkident:
		if m.peek(1).is(tkequal) {
			return m.assignment()
		}

	case tkpass:
		m.next()
		return &stmtpass{}

	case tkreturn:
		m.next()
		return &stmtreturn{expr: m.expression()}

	case tkbraceopen:
		return m.block()

	case tkunset:
		return m.unset()

	case tkif:
		return m.ifstmt()

	case tkwhile:
		return m.while()

	case tkfor:
		return m.forstmt()

	case tkload:
		return m.load()
	}

	expr := m.expression()

	if getattr, ok := expr.(*exprgetattr); ok {
		if m.match(tkequal) {
			return &stmtsetattr{
				expr:  getattr.expr,
				name:  getattr.name,
				value: m.expression(),
			}
		}
	}

	return &stmtexpr{expr: expr}
}

func (m *parser) block() statement {
	var statements []statement

	m.expect(tkbraceopen)

	for !m.finished() {
		if m.match(tkbraceclose) {
			return &stmtblock{
				statements: statements,
			}
		}

		statements = append(statements, m.statement())
	}

	m.unexpectedeof()
	return nil
}

func (m *parser) assignment() statement {
	name := m.expect(tkident)

	m.expect(tkequal)

	expr := m.expression()

	return &stmtassign{
		name: name.value,
		expr: expr,
	}
}

func (m *parser) unset() statement {
	var stmt stmtunset

	m.expect(tkunset)

	for !m.finished() {
		if m.match(tksemicolon) {
			break
		}

		if !m.check(tkident) {
			break
		}

		stmt.names = append(stmt.names, m.next().value)
	}

	return &stmt
}

func (m *parser) ifstmt() statement {
	var falsestmt statement

	m.expect(tkif)

	condition := m.expression()
	truestmt := m.statement()

	if m.match(tkelse) {
		falsestmt = m.statement()
	}

	return &stmtif{condition: condition, truestmt: truestmt, falsestmt: falsestmt}
}

func (m *parser) while() statement {
	m.expect(tkwhile)

	condition := m.expression()
	body := m.statement()

	return &stmtwhile{
		condition: condition,
		body:      body,
	}
}

func (m *parser) forstmt() statement {
	m.expect(tkfor)

	hasparen := m.match(tkparenopen)

	init := m.statement()
	m.expect(tksemicolon)
	cond := m.expression()
	m.expect(tksemicolon)
	incr := m.statement()

	if hasparen {
		m.expect(tkparenclose)
	}

	body := m.statement()

	return &stmtblock{
		statements: []statement{
			init,
			&stmtwhile{
				condition: cond,
				body: &stmtblock{
					statements: []statement{
						body,
						incr,
					},
				},
			},
		},
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

	return &stmtload{modules: modules}
}

func (m *parser) parse() []statement {
	m.reset()
	m.desugar()

	statements := []statement{}

	for !m.finished() {
		statements = append(statements, m.statement())
		m.match(tksemicolon)
	}

	return statements
}

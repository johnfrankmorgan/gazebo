package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
)

type statement interface {
	compiler
}

type exprstmt struct {
	expr expression
}

func (m *exprstmt) compile() Code {
	return m.expr.compile()
}

type assign struct {
	name token
	expr expression
}

func (m *assign) compile() Code {
	errors.ErrCompile.Expect(m.name.is(tkident), "expected tkident, got %s", m.name.typ.name())

	return append(m.expr.compile(), op.StoreName.Ins(m.name.value))
}

type block struct {
	statements []statement
}

func (m *block) compile() Code {
	code := Code{}

	for _, stmt := range m.statements {
		code = append(code, stmt.compile()...)
	}

	return code
}

type ifstmt struct {
	condition expression
	truebody  statement
	falsebody statement
}

func (m *ifstmt) compile() Code {
	return nil
}

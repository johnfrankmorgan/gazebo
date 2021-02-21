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
	truestmt  statement
	falsestmt statement
}

func (m *ifstmt) compile() Code {
	truecode := m.truestmt.compile()
	falsecode := Code{}

	if m.falsestmt != nil {
		falsecode = m.falsestmt.compile()
	}

	falsecode = append(falsecode, op.RelJump.Ins(len(truecode)))

	condition := append(m.condition.compile(), op.RelJumpIfTrue.Ins(len(falsecode)))

	code := append(condition, falsecode...)
	return append(code, truecode...)
}

type while struct {
	condition expression
	body      statement
}

func (m *while) compile() Code {
	body := m.body.compile()
	cond := append(m.condition.compile(), op.RelJumpIfFalse.Ins(len(body)+1))
	body = append(body, op.RelJump.Ins(-len(body)-len(cond)-1))
	return append(cond, body...)
}

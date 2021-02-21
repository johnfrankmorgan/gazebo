package compiler

import (
	"strconv"

	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

type expression interface {
	compiler
}

type binary struct {
	op    token
	left  expression
	right expression
}

func (m *binary) compile() Code {
	fun, ok := protocols.BinaryOperators[m.op.value]
	errors.ErrCompile.Expect(ok, "unknown binary operator %s %s", m.op.value, m.op.typ.name())

	code := Code{op.LoadName.Ins(fun)}

	code = append(code, m.left.compile()...)
	code = append(code, m.right.compile()...)

	return append(code, op.CallFunc.Ins(2))
}

type unary struct {
	op    token
	right expression
}

func (m *unary) compile() Code {
	fun, ok := protocols.UnaryOperators[m.op.value]
	errors.ErrCompile.Expect(ok, "unknown unary operator %s %s", m.op.value, m.op.typ.name())

	code := Code{op.LoadName.Ins(fun)}

	code = append(code, m.right.compile()...)

	return append(code, op.CallFunc.Ins(1))
}

type literal struct {
	token token
}

func (m *literal) compile() Code {
	switch m.token.typ {
	case tknumber:
		value, err := strconv.ParseFloat(m.token.value, 64)
		errors.ErrCompile.ExpectNil(err, "%v", err)
		return Code{op.LoadConst.Ins(value)}

	case tkstring:
		value := m.token.value[1 : len(m.token.value)-1]
		return Code{op.LoadConst.Ins(value)}

	case tkident:
		return Code{op.LoadName.Ins(m.token.value)}
	}

	errors.ErrCompile.Panic("unknown literal: %s %s", m.token.typ.name(), m.token.value)
	return nil
}

type group struct {
	expr expression
}

func (m *group) compile() Code {
	return m.expr.compile()
}

type funcall struct {
	name expression
	args []expression
}

func (m *funcall) compile() Code {
	code := m.name.compile()
	argc := 0

	for _, arg := range m.args {
		code = append(code, arg.compile()...)
		argc++
	}

	return append(code, op.CallFunc.Ins(argc))
}

type fundef struct {
	args []string
	body statement
}

func (m *fundef) compile() Code {
	return Code{
		op.PushValue.Ins(m.args),
		op.PushValue.Ins(m.body.compile()),
		op.MakeFunc.Ins(len(m.args)),
	}
}

type attributelookup struct {
	expr expression
	name string
}

func (m *attributelookup) compile() Code {
	return append(m.expr.compile(), op.AttributeGet.Ins(m.name))
}

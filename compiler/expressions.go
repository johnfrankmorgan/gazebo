package compiler

import (
	"strconv"

	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

type expression interface {
	compiler
}

type exprbinary struct {
	op    token
	left  expression
	right expression
}

func (m *exprbinary) compile() Code {
	fun, ok := protocols.BinaryOperators[m.op.value]
	errors.ErrCompile.Expect(ok, "unknown binary operator %s %s", m.op.value, m.op.typ.name())

	code := m.left.compile()
	code = append(code, op.GetAttr.Ins(fun))
	code = append(code, m.right.compile()...)

	return append(code, op.CallFunc.Ins(1))
}

type exprunary struct {
	op    token
	right expression
}

func (m *exprunary) compile() Code {
	fun, ok := protocols.UnaryOperators[m.op.value]
	errors.ErrCompile.Expect(ok, "unknown unary operator %s %s", m.op.value, m.op.typ.name())

	code := m.right.compile()
	code = append(code, op.GetAttr.Ins(fun))

	return append(code, op.CallFunc.Ins(0))
}

type exprliteral struct {
	token token
}

func (m *exprliteral) compile() Code {
	switch m.token.typ {
	case tknumber:
		value, err := strconv.ParseFloat(m.token.value, 64)
		errors.ErrCompile.ExpectNilError(err)
		return Code{op.LoadConst.Ins(value)}

	case tkstring:
		value, err := strconv.Unquote(m.token.value)
		errors.ErrCompile.ExpectNilError(err)
		return Code{op.LoadConst.Ins(value)}

	case tkident:
		return Code{op.GetName.Ins(m.token.value)}
	}

	errors.ErrCompile.Panic("unknown literal: %s %s", m.token.typ.name(), m.token.value)
	return nil
}

type exprgroup struct {
	expr expression
}

func (m *exprgroup) compile() Code {
	return m.expr.compile()
}

type exprfuncall struct {
	name expression
	args []expression
}

func (m *exprfuncall) compile() Code {
	code := m.name.compile()
	argc := 0

	for _, arg := range m.args {
		code = append(code, arg.compile()...)
		argc++
	}

	return append(code, op.CallFunc.Ins(argc))
}

type exprfun struct {
	args []string
	body statement
}

func (m *exprfun) compile() Code {
	return Code{
		op.PushValue.Ins(m.args),
		op.PushValue.Ins(m.body.compile()),
		op.MakeFunc.Ins(len(m.args)),
	}
}

type exprgetattr struct {
	expr expression
	name string
}

func (m *exprgetattr) compile() Code {
	return append(m.expr.compile(), op.GetAttr.Ins(m.name))
}

type exprlist struct {
	expressions []expression
}

func (m *exprlist) compile() Code {
	code := Code{}

	for _, expr := range m.expressions {
		code = append(code, expr.compile()...)
	}

	return append(code, op.MakeList.Ins(len(m.expressions)))
}

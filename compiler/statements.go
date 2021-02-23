package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

type statement interface {
	compiler
}

type stmtexpr struct {
	expr expression
}

func (m *stmtexpr) compile() Code {
	return m.expr.compile()
}

type stmtassign struct {
	name string
	expr expression
}

func (m *stmtassign) compile() Code {
	return append(m.expr.compile(), op.StoreName.Ins(m.name))
}

type stmtunset struct {
	names []string
}

func (m *stmtunset) compile() Code {
	code := Code{}

	for _, name := range m.names {
		code = append(code, op.RemoveName.Ins(name))
	}

	return code
}

type stmtblock struct {
	statements []statement
}

func (m *stmtblock) compile() Code {
	code := Code{}

	for _, stmt := range m.statements {
		code = append(code, stmt.compile()...)
	}

	return code
}

type stmtif struct {
	condition expression
	truestmt  statement
	falsestmt statement
}

func (m *stmtif) compile() Code {
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

type stmtwhile struct {
	condition expression
	body      statement
}

func (m *stmtwhile) compile() Code {
	body := m.body.compile()
	cond := append(m.condition.compile(), op.RelJumpIfFalse.Ins(len(body)+1))
	body = append(body, op.RelJump.Ins(-len(body)-len(cond)-1))
	return append(cond, body...)
}

type stmtload struct {
	modules []string
}

func (m *stmtload) compile() Code {
	code := Code{}

	for _, module := range m.modules {
		code = append(code, op.LoadModule.Ins(module))
	}

	return code
}

type stmtpass struct{}

func (m *stmtpass) compile() Code {
	return Code{op.NoOp.Ins(nil)}
}

type stmtreturn struct {
	expr expression
}

func (m *stmtreturn) compile() Code {
	return append(m.expr.compile(), op.Return.Ins(nil))
}

type stmtsetattr struct {
	expr  expression
	name  string
	value expression
}

func (m *stmtsetattr) compile() Code {
	code := m.expr.compile()
	code = append(code, m.value.compile()...)
	return append(code, op.AttributeSet.Ins(m.name))
}

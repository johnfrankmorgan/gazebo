package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

type statement interface {
	compiler
}

func fillbreak(code Code) Code {
	for i, ins := range code {
		if ins.Opcode != op.Placeholder || ins.Arg.(int) != op.PlaceholderBreak {
			continue
		}

		code[i] = op.RelJump.Ins(len(code) - i - 1)
	}

	return code
}

func fillcontinue(code Code) Code {
	for i, ins := range code {
		if ins.Opcode != op.Placeholder || ins.Arg.(int) != op.PlaceholderContinue {
			continue
		}

		code[i] = op.RelJump.Ins(-i - 1)
	}

	return code
}

func checkloop(code Code) Code {
	code = fillbreak(code)
	return fillcontinue(code)
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
	return append(m.expr.compile(), op.SetName.Ins(m.name))
}

type stmtdel struct {
	name string
}

func (m *stmtdel) compile() Code {
	return Code{op.DelName.Ins(m.name)}
}

type stmtdelattr struct {
	expr expression
	name string
}

func (m *stmtdelattr) compile() Code {
	return append(m.expr.compile(), op.DelAttr.Ins(m.name))
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
	return checkloop(append(cond, body...))
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
	return append(code, op.SetAttr.Ins(m.name))
}

type stmtbreak struct{}

func (m *stmtbreak) compile() Code {
	return Code{op.Placeholder.Ins(op.PlaceholderBreak)}
}

type stmtcontinue struct{}

func (m *stmtcontinue) compile() Code {
	return Code{op.Placeholder.Ins(op.PlaceholderContinue)}
}

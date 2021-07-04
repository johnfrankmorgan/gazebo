package compiler

import (
	"fmt"
	"strconv"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

// FIXME: add a way for instructions to create new scopes
//        it would be cool if SIf, SWhile, etc didn't pollute
//        their outer scope
//
// Eg:
//    if (...) {
//        x = 1;
//    }
//    # x is undefined here

var _ ast.Visitor = &Compiler{}

type Compiler struct {
	code code
}

func New() *Compiler {
	return &Compiler{}
}

type FuncDef struct {
	Args []string
	Body []Ins
}

func (m *Compiler) todo() {
	panic(fmt.Errorf("todo"))
}

func (m *Compiler) emit(op op.Op, arg ...interface{}) label {
	ins := Ins{Op: op}

	if len(arg) > 0 {
		ins.Arg = arg[0]
	}

	label := m.code.label()
	m.code.append(ins)
	return label
}

func (m *Compiler) compile(node ast.Node) code {
	defer func(code code) {
		m.code = code
	}(m.code)

	m.code = code{}
	node.Accept(m)

	return m.code
}

func (m *Compiler) Compile(ast *ast.AST) []Ins {
	ast.Traverse(m)

	return m.code.instructions()
}

func (m *Compiler) VisitEAssign(stmt *ast.EAssign) {
	stmt.Expr.Accept(m)
	m.emit(op.StoreName, stmt.Ident)
}

func (m *Compiler) VisitEGroup(expr *ast.EGroup) {
	expr.Expr.Accept(m)
}

func (m *Compiler) VisitEBinary(expr *ast.EBinary) {
	expr.LHS.Accept(m)
	expr.RHS.Accept(m)

	ops := map[ast.BinOp]op.Op{
		ast.BinOpAdd:       op.BinAdd,
		ast.BinOpSub:       op.BinSub,
		ast.BinOpMul:       op.BinMul,
		ast.BinOpDiv:       op.BinDiv,
		ast.BinOpEq:        op.BinEq,
		ast.BinOpNEq:       op.BinNEq,
		ast.BinOpGreater:   op.BinGreater,
		ast.BinOpGreaterEq: op.BinGreaterEq,
		ast.BinOpLess:      op.BinLess,
		ast.BinOpLessEq:    op.BinLessEq,
	}

	if op, ok := ops[expr.Op]; ok {
		m.emit(op)
		return
	}

	panic(
		fmt.Errorf(
			"unknown binary operator %q",
			expr.Op,
		),
	)
}

func (m *Compiler) VisitEUnary(expr *ast.EUnary) {
	m.todo()
}

func (m *Compiler) VisitELiteral(expr *ast.ELiteral) {
	switch expr.Type {
	case ast.LitTypeIdent:
		m.emit(op.LoadName, expr.Lexeme)

	case ast.LitTypeNumber:
		// FIXME: this shouldn't fail but we should probably still check for an error
		val, _ := strconv.ParseFloat(expr.Lexeme, 64)
		m.emit(op.LoadConst, val)

	case ast.LitTypeString:
		m.todo()
	}
}

func (m *Compiler) VisitEFuncDef(expr *ast.EFuncDef) {
	body := m.compile(expr.Body)

	m.emit(op.MakeFunction, &FuncDef{
		Args: expr.Args,
		Body: body.instructions(),
	})
}

func (m *Compiler) VisitECall(expr *ast.ECall) {
	for _, arg := range expr.Args {
		arg.Accept(m)
	}

	expr.Expr.Accept(m)
	m.emit(op.Call, len(expr.Args))
}

func (m *Compiler) VisitEAttrGet(expr *ast.EAttrGet) {
	expr.Expr.Accept(m)
	m.emit(op.AttrGet, expr.Attr)
}

func (m *Compiler) VisitEAttrSet(expr *ast.EAttrSet) {
	expr.Value.Accept(m)
	expr.Expr.Accept(m)
	m.emit(op.AttrSet, expr.Attr)
}

func (m *Compiler) VisitSBlock(stmt *ast.SBlock) {
	for _, stmt := range stmt.Stmts {
		stmt.Accept(m)
	}
}

func (m *Compiler) VisitSExpr(stmt *ast.SExpr) {
	stmt.Expr.Accept(m)
}

func (m *Compiler) VisitSIf(stmt *ast.SIf) {
	// compile condition
	stmt.Condition.Accept(m)

	// label for the instruction to jump over "true block" if the condition is false
	jumpOverTrueBlock := m.emit(op.RelJumpIfFalse)

	// compile "true block" and calculate its size
	sizeOfTrueBlock := m.code.len()
	stmt.TrueBlock.Accept(m)
	// if the "true block" is executed, we need to make sure we jump over the "false block"
	jumpOverFalseBlock := m.emit(op.RelJump)
	sizeOfTrueBlock = m.code.len() - sizeOfTrueBlock

	// now we can tell the RelJumpIfFalse instruction how far it needs to jump
	m.code.labelled(jumpOverTrueBlock).Arg = sizeOfTrueBlock

	// same as above, compile "false block" and calculate its size
	sizeOfFalseBlock := m.code.len()
	if stmt.FalseBlock != nil {
		stmt.FalseBlock.Accept(m)
	}
	sizeOfFalseBlock = m.code.len() - sizeOfFalseBlock

	// update the RelJump over the "false block"'s code path with the correct size
	m.code.labelled(jumpOverFalseBlock).Arg = sizeOfFalseBlock
}

func (m *Compiler) VisitSWhile(stmt *ast.SWhile) {
	// create a label pointing at the while statement's condition
	// we'll unconditionally jump here at the end of the while loop's body
	condition := m.code.label()
	stmt.Condition.Accept(m)

	// jump over the body if the condition evaluates to false
	jumpOverBody := m.emit(op.RelJumpIfFalse)

	// calculate the size of the loop's body and update the previous instruction
	sizeOfBody := m.code.len()
	stmt.Body.Accept(m)
	m.emit(op.Jump, int(condition))
	sizeOfBody = m.code.len() - sizeOfBody

	m.code.labelled(jumpOverBody).Arg = sizeOfBody
}

func (m *Compiler) VisitSReturn(stmt *ast.SReturn) {
	stmt.Expr.Accept(m)
	m.emit(op.Return)
}

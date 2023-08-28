package compiler

import (
	"gazebo/ast"
	"gazebo/op"
)

func Compile(program *ast.Program) *Code {
	compiler := &compiler{code: new(Code)}

	program.Traverse(compiler)

	return compiler.code
}

type compiler struct {
	code *Code
}

var _ ast.Visitor = (*compiler)(nil)

// statements

func (c *compiler) VisitAssignment(node *ast.Assignment) {
	node.Expression.Accept(c)
	c.code.Emit(op.StoreName, c.code.Name(node.Identifier))
}

func (c *compiler) VisitBlock(node *ast.Block) {
	body := new(Code)

	c.code.Emit(op.ExecuteChild, c.code.Child(body))
	c.code = body

	for _, statement := range node.Statements {
		statement.Accept(c)
	}

	c.code = body.Parent
}

func (*compiler) VisitComment(*ast.Comment) {
	//
}

func (c *compiler) VisitExpressionStatement(node *ast.ExpressionStatement) {
	node.Expression.Accept(c)
}

func (c *compiler) VisitIf(node *ast.If) {
	node.Condition.Accept(c)

	els := new(Code)
	c.code.Emit(op.RelativeJumpIfTrue, 4)
	c.code.Emit(op.ExecuteChild, c.code.Child(els))
	c.code.Emit(op.RelativeJump, 2)

	if node.Else != nil {
		c.code = els
		node.Else.Accept(c)
		c.code = els.Parent
	}

	body := new(Code)
	c.code.Emit(op.ExecuteChild, c.code.Child(body))
	c.code = body

	node.Body.Accept(c)

	c.code = body.Parent
}

func (c *compiler) VisitWhile(node *ast.While) {
	pc := len(c.code.Opcodes)

	node.Condition.Accept(c)

	body := new(Code)
	c.code.Emit(op.RelativeJumpIfFalse, 2)
	c.code.Emit(op.ExecuteChild, c.code.Child(body))
	c.code = body

	node.Body.Accept(c)

	c.code = body.Parent
	c.code.Emit(op.Jump, pc)
}

// expressions

func (c *compiler) VisitBinary(node *ast.Binary) {
	node.Left.Accept(c)
	node.Right.Accept(c)

	c.code.Emit(op.Binaries[node.Op])
}

func (c *compiler) VisitCall(node *ast.Call) {
	for i := len(node.Arguments) - 1; i >= 0; i-- {
		node.Arguments[i].Accept(c)
	}

	node.Expression.Accept(c)

	c.code.Emit(op.Call, len(node.Arguments))
}

func (c *compiler) VisitFalse(*ast.False) {
	c.code.Emit(op.LoadFalse)
}

func (c *compiler) VisitFloat(node *ast.Float) {
	c.constant(node)
}

func (c *compiler) VisitGroup(node *ast.Group) {
	node.Expression.Accept(c)
}

func (c *compiler) VisitIdentifier(node *ast.Identifier) {
	c.code.Emit(op.LoadName, c.code.Name(node.Name))
}

func (c *compiler) VisitInteger(node *ast.Integer) {
	c.constant(node)
}

func (c *compiler) VisitNull(*ast.Null) {
	c.code.Emit(op.LoadNull)
}

func (c *compiler) VisitString(node *ast.String) {
	c.constant(node)
}

func (c *compiler) VisitTrue(*ast.True) {
	c.code.Emit(op.LoadTrue)
}

func (c *compiler) VisitUnary(node *ast.Unary) {
	node.Right.Accept(c)

	c.code.Emit(op.Unaries[node.Op])
}

func (c *compiler) constant(node ast.Expression) {
	c.code.Emit(op.LoadConstant, c.code.Constant(node))
}

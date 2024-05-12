package compile

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
)

func (c *compiler) compileStmt(node ast.Stmt) {
	switch node := node.(type) {
	case stmt.Assign:
		c.compileStmtAssign(node)

	case stmt.Block:
		c.compileStmtBlock(node)

	case stmt.Break:
		c.compileStmtBreak(node)

	case stmt.Continue:
		c.compileStmtContinue(node)

	case stmt.Expr:
		c.compileStmtExpr(node)

	case stmt.If:
		c.compileStmtIf(node)

	case stmt.Return:
		c.compileStmtReturn(node)

	case stmt.While:
		c.compileStmtWhile(node)

	default:
		panic(fmt.Errorf("compile: unknown statement type: %T", node))
	}
}

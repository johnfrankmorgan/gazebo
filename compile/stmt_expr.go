package compile

import "github.com/johnfrankmorgan/gazebo/ast/stmt"

func (c *compiler) compileStmtExpr(node stmt.Expr) {
	c.compileExpr(node.Inner)
}

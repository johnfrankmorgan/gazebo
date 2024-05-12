package compile

import "github.com/johnfrankmorgan/gazebo/ast/expr"

func (c *compiler) compileExprGroup(node expr.Group) {
	c.compileExpr(node.Inner)
}

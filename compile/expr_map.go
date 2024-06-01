package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprMap(node expr.Map) {
	for _, pair := range node.Items {
		c.compileExpr(pair.Key)
		c.compileExpr(pair.Value)
	}

	c.emit(opcode.MakeMap, len(node.Items))
}

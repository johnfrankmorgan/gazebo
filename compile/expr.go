package compile

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/johnfrankmorgan/gazebo/ast/expr"
)

func (c *compiler) compileExpr(node ast.Expr) {
	switch node := node.(type) {
	case expr.Attr:
		c.compileExprAttr(node)

	case expr.Binary:
		c.compileExprBinary(node)

	case expr.False:
		c.compileExprFalse(node)

	case expr.Float:
		c.compileExprFloat(node)

	case expr.Group:
		c.compileExprGroup(node)

	case expr.Ident:
		c.compileExprIdent(node)

	case expr.Index:
		c.compileExprIndex(node)

	case expr.Int:
		c.compileExprInt(node)

	case expr.List:
		c.compileExprList(node)

	case expr.Map:
		c.compileExprMap(node)

	case expr.Nil:
		c.compileExprNil(node)

	case expr.String:
		c.compileExprString(node)

	case expr.Ternary:
		c.compileExprTernary(node)

	case expr.True:
		c.compileExprTrue(node)

	case expr.Tuple:
		c.compileExprTuple(node)

	case expr.Unary:
		c.compileExprUnary(node)

	default:
		panic(fmt.Errorf("compile: unknown expression type: %T", node))
	}
}

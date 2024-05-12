package compile

import "github.com/johnfrankmorgan/gazebo/ast/stmt"

func (c *compiler) compileStmtBlock(node stmt.Block) {
	for _, stmt := range node.Statements {
		c.compileStmt(stmt)
	}
}

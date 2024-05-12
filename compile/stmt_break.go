package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtBreak(stmt.Break) {
	c.emit(opcode.Jump, pendingBreak)
}

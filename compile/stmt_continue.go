package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtContinue(stmt.Continue) {
	c.emit(opcode.Jump, pendingContinue)
}

package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Del struct {
	Name string
}

func (m *Del) Compile() code.Code {
	return code.Code{op.DelName.Ins(m.Name)}
}

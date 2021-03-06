package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Load struct {
	Names []string
}

func (m *Load) Compile() code.Code {
	code := code.Code{}

	for _, name := range m.Names {
		code = append(code, op.LoadModule.Ins(name))
	}

	return code
}

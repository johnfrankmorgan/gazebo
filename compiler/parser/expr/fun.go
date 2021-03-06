package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Fun struct {
	Args []string
	Body Expression
}

func (m *Fun) Compile() code.Code {
	return code.Code{
		op.PushValue.Ins(m.Args),
		op.PushValue.Ins(m.Body.Compile()),
		op.MakeFunc.Ins(len(m.Args)),
	}
}

package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type List struct {
	Values []Expression
}

func (m *List) Compile() code.Code {
	code := code.Code{}

	for _, value := range m.Values {
		code = append(code, value.Compile()...)
	}

	return append(code, op.MakeList.Ins(len(m.Values)))
}

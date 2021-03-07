package expr

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Map struct {
	Keys   []Expression
	Values []Expression
}

func (m *Map) Compile() code.Code {
	if len(m.Keys) != len(m.Values) {
		panic(fmt.Errorf(
			"expected %d keys and values, got %d keys and %d values",
			len(m.Keys),
			len(m.Keys),
			len(m.Values),
		))
	}

	code := code.Code{}

	for idx, key := range m.Keys {
		value := m.Values[idx]
		code = append(code, key.Compile()...)
		code = append(code, value.Compile()...)
	}

	return append(code, op.MakeMap.Ins(len(m.Keys)))
}

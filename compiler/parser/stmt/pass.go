package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Pass struct{}

func (_ *Pass) Compile() code.Code {
	return code.Code{op.NoOp.Ins(nil)}
}

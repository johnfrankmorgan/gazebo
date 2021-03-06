package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Break struct{}

func (_ *Break) Compile() code.Code {
	return code.Code{op.Placeholder.Ins(op.PlaceholderBreak)}
}

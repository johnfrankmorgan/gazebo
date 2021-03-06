package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Continue struct{}

func (_ *Continue) Compile() code.Code {
	return code.Code{op.Placeholder.Ins(op.PlaceholderContinue)}
}

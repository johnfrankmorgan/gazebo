package stmt

import "github.com/johnfrankmorgan/gazebo/compiler/code"

type Statement interface {
	code.Compiler
}

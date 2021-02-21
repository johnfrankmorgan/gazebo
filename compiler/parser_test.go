package compiler

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/debug"
)

func TestParse(t *testing.T) {
	debug.Enable()
	defer debug.Disable()

	expr := parse("1 == (-2 + 3) != true")

	dumpexpression(expr, 0)
}

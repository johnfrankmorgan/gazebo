package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
)

type Code []op.Instruction

func (m Code) dump() {
	for idx, ins := range m {
		debug.Printf(
			"%6d %18s (0x%02x) %v\n",
			idx,
			ins.Opcode.Name(),
			int(ins.Opcode),
			ins.Arg,
		)
	}
}

func Compile(source string) (code Code, err error) {
	defer func() {
		recovered := recover()

		if recovered == nil {
			return
		}

		if gerr, ok := recovered.(*errors.Error); ok {
			err = gerr
			return
		}

		panic(recovered)
	}()

	expr := parse(source)

	if debug.Enabled() {
		dumpexpression(expr, 0)
	}

	code = expr.compile()

	if debug.Enabled() {
		code.dump()
	}

	return code, nil
}

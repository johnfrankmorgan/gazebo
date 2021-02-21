package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/kr/pretty"
)

type compiler interface {
	compile() Code
}

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

	for _, stmt := range parse(source) {
		code = append(code, stmt.compile()...)

		if debug.Enabled() {
			debug.Printf("%s", pretty.Sprintf("%# v", stmt))
		}
	}

	if debug.Enabled() {
		code.dump()
	}

	return code, nil
}

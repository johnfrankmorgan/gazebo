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
	defer errors.Handle(&err)

	parser := parser{tokens: tokenize(source)}

	if debug.Enabled() {
		debug.Printf("TOKENS\n")
		parser.tokens.dump()
		debug.Printf("\n")
	}

	ast := parser.parse()

	if debug.Enabled() {
		debug.Printf("DESUGARED\n")
		parser.tokens.dump()
		debug.Printf("\n")

		debug.Printf("AST\n")
		debug.Printf("%s\n\n", pretty.Sprintf("%# v", ast))
	}

	for _, stmt := range ast {
		code = append(code, stmt.compile()...)
	}

	if debug.Enabled() {
		debug.Printf("BYTECODE\n")
		code.dump()
		debug.Printf("\n")
	}

	return code, nil
}

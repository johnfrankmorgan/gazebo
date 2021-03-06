package compiler

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
	"github.com/johnfrankmorgan/gazebo/compiler/parser"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/kr/pretty"
)

func Compile(source string) (code code.Code, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			if perr, ok := perr.(error); ok {
				err = perr
			}
		}
	}()

	tokens, err := lexer.New([]byte(source)).Lex()
	if err != nil {
		return nil, err
	}

	if debug.Enabled() {
		debug.Printf("TOKENS\n")
		tokens.Dump()
		debug.Printf("\n")
	}

	tokens = parser.NewDesugarer(tokens).Desugar()

	if debug.Enabled() {
		debug.Printf("DESUGARED\n")
		tokens.Dump()
		debug.Printf("\n")
	}

	ast, err := parser.New(tokens).Parse()
	if err != nil {
		return nil, err
	}

	if debug.Enabled() {
		debug.Printf("AST\n")
		debug.Printf("%s\n\n", pretty.Sprintf("%# v", ast))
	}

	code = ast.Compile()

	if debug.Enabled() {
		debug.Printf("BYTECODE\n")
		code.Dump()
		debug.Printf("\n")
	}

	return code, nil
}

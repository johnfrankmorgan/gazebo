package parser

import (
	"gazebo/ast"
)

//go:generate goyacc grammar.y

func Parse(source string, file string) (*ast.Program, error) {
	p := &lexer{
		source:  source,
		program: &ast.Program{Source: source},
	}

	p.position.current = ast.Position{File: file, Line: 1, Column: 1}

	yyParse(p)

	return p.program, p.err
}

func ParseBytes(source []byte, file string) (*ast.Program, error) {
	return Parse(string(source), file)
}

func init() {
	yyErrorVerbose = true
}

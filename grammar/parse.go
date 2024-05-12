package grammar

import (
	"errors"
	"fmt"
	"io"

	"github.com/johnfrankmorgan/gazebo/ast"
)

type Program struct {
	File       string
	Source     string
	Tokens     []Token
	Statements []ast.Stmt
}

var ErrParse = errors.New("grammar: parse error")

func Parse(rd io.Reader) (*Program, error) {
	source, err := io.ReadAll(rd)
	if err != nil {
		return nil, fmt.Errorf("grammar: failed to read source: %w", err)
	}

	tokens, err := Lex(string(source))
	if err != nil {
		return nil, fmt.Errorf("grammar: failed to lex source: %w", err)
	}

	file := ""

	if named, ok := rd.(interface{ Name() string }); ok {
		file = named.Name()
	}

	lexer := lexer{
		tokens: tokens,
		program: &Program{
			File:   file,
			Source: string(source),
			Tokens: append([]Token(nil), tokens...),
		},
	}

	yyParse(&lexer)

	if lexer.err != nil {
		return nil, lexer.err
	}

	return lexer.program, nil
}

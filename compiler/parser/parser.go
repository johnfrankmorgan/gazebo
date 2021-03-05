package parser

import "github.com/johnfrankmorgan/gazebo/compiler/lexer"

type Parser struct {
	stream *TokenStream
}

func New(tokens lexer.Tokens) *Parser {
	return &Parser{
		stream: NewTokenStream(tokens),
	}
}

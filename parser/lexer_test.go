package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerLex(t *testing.T) {
	assert := assert.New(t)

	source := "#comment\n(){}.,;!<<===>>==+-*/"

	expected := []TKind{
		TComment,
		TParenOpen,
		TParenClose,
		TBraceOpen,
		TBraceClose,
		TDot,
		TComma,
		TSemicolon,
		TBang,
		TLess,
		TLessEqual,
		TEqualEqual,
		TGreater,
		TGreaterEqual,
		TEqual,
		TPlus,
		TMinus,
		TStar,
		TSlash,
		TEOF,
	}

	lexer := Lexer{source: source}
	tokens := lexer.Lex()

	got := []TKind{}

	for _, tk := range tokens {
		got = append(got, tk.kind)
	}

	assert.Equal(expected, got)
}

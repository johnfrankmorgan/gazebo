package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerLex(t *testing.T) {
	assert := assert.New(t)

	source := "#comment\n()  \tif{}.,;an_identifier!<<===>>==+-*1234/"

	expected := []TKind{
		TComment,
		TParenOpen,
		TParenClose,
		TWhitespace,
		TIf,
		TBraceOpen,
		TBraceClose,
		TDot,
		TComma,
		TSemicolon,
		TIdent,
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
		TNumber,
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

func TestLexerLexPanics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		lexer := Lexer{source: "$"}
		lexer.Lex()
	})
}

package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lex(source string) (Tokens, error) {
	return New([]byte(source)).Lex()
}

func TestLexerLex(t *testing.T) {
	assert := assert.New(t)

	tokens, err := lex(`
		# comment
		// comment
		; . , if else return while for break
		continue fun del load pass
		in and or () {} [] = == ! !=
		+ += - -= * *= / /= < <= > >=
		"string" 84.59 ident
	`)

	assert.Nil(err)

	got := []TokenType{}

	for _, t := range tokens {
		got = append(got, t.Type)
	}

	exp := []TokenType{
		TkSemicolon,
		TkDot,
		TkComma,
		TkIf,
		TkElse,
		TkReturn,
		TkWhile,
		TkFor,
		TkBreak,
		TkContinue,
		TkFun,
		TkDel,
		TkLoad,
		TkPass,
		TkIn,
		TkAnd,
		TkOr,
		TkParenOpen,
		TkParenClose,
		TkBraceOpen,
		TkBraceClose,
		TkBracketOpen,
		TkBracketClose,
		TkEqual,
		TkEqualEqual,
		TkBang,
		TkBangEqual,
		TkPlus,
		TkPlusEqual,
		TkMinus,
		TkMinusEqual,
		TkStar,
		TkStarEqual,
		TkSlash,
		TkSlashEqual,
		TkLess,
		TkLessEqual,
		TkGreater,
		TkGreaterEqual,
		TkString,
		TkNumber,
		TkIdent,
		TkEOF,
	}

	assert.Equal(exp, got)
}

func TestLexerLexUnexpectedEOF(t *testing.T) {
	assert := assert.New(t)

	_, err := lex(`"string`)
	assert.ErrorIs(err, ErrUnexpectedEOF)
}

func TestLexerLexUnexpectedRune(t *testing.T) {
	assert := assert.New(t)

	_, err := lex(`'`)
	assert.ErrorIs(err, ErrUnexpectedRune)
}

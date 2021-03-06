package lexer

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/debug"
)

type TokenType int

func (m TokenType) String() string {
	if name, ok := _tokens[m]; ok {
		return name
	}

	return "TkUnknown"
}

type Token struct {
	Type     TokenType
	Value    string
	Position int
}

func (m Token) String() string {
	return fmt.Sprintf("%s{%q, %d}", m.Type, m.Value, m.Position)
}

func (m Token) Is(types ...TokenType) bool {
	for _, t := range types {
		if m.Type == t {
			return true
		}
	}

	return false
}

type Tokens []Token

func (m Tokens) Dump() {
	for idx, token := range m {
		debug.Printf(
			"%6d: %02x :: %16s :: %q\n",
			idx,
			int(token.Type),
			token.Type.String(),
			token.Value,
		)
	}
}

const (
	TkInvalid TokenType = iota
	TkComment
	TkNewline
	TkWhitespace
	TkSemicolon
	TkDot
	TkComma
	TkIf
	TkElse
	TkReturn
	TkWhile
	TkFor
	TkBreak
	TkContinue
	TkFun
	TkDel
	TkLoad
	TkPass
	TkIn
	TkAnd
	TkOr
	TkParenOpen
	TkParenClose
	TkBraceOpen
	TkBraceClose
	TkBracketOpen
	TkBracketClose
	TkEqual
	TkEqualEqual
	TkBang
	TkBangEqual
	TkPlus
	TkPlusEqual
	TkMinus
	TkMinusEqual
	TkStar
	TkStarEqual
	TkSlash
	TkSlashEqual
	TkLess
	TkLessEqual
	TkGreater
	TkGreaterEqual
	TkString
	TkNumber
	TkIdent
	TkEOF
)

var _tokens = map[TokenType]string{
	TkInvalid:      "TkInvalid",
	TkComment:      "TkComment",
	TkNewline:      "TkNewline",
	TkWhitespace:   "TkWhitespace",
	TkSemicolon:    "TkSemicolon",
	TkDot:          "TkDot",
	TkComma:        "TkComma",
	TkIf:           "TkIf",
	TkElse:         "TkElse",
	TkReturn:       "TkReturn",
	TkWhile:        "TkWhile",
	TkFor:          "TkFor",
	TkBreak:        "TkBreak",
	TkContinue:     "TkContinue",
	TkFun:          "TkFun",
	TkDel:          "TkDel",
	TkLoad:         "TkLoad",
	TkPass:         "TkPass",
	TkIn:           "TkIn",
	TkAnd:          "TkAnd",
	TkOr:           "TkOr",
	TkParenOpen:    "TkParenOpen",
	TkParenClose:   "TkParenClose",
	TkBraceOpen:    "TkBraceOpen",
	TkBraceClose:   "TkBraceClose",
	TkBracketOpen:  "TkBracketOpen",
	TkBracketClose: "TkBracketClose",
	TkEqual:        "TkEqual",
	TkEqualEqual:   "TkEqualEqual",
	TkBang:         "TkBang",
	TkBangEqual:    "TkBangEqual",
	TkPlus:         "TkPlus",
	TkPlusEqual:    "TkPlusEqual",
	TkMinus:        "TkMinus",
	TkMinusEqual:   "TkMinusEqual",
	TkStar:         "TkStar",
	TkStarEqual:    "TkStarEqual",
	TkSlash:        "TkSlash",
	TkSlashEqual:   "TkSlashEqual",
	TkLess:         "TkLess",
	TkLessEqual:    "TkLessEqual",
	TkGreater:      "TkGreater",
	TkGreaterEqual: "TkGreaterEqual",
	TkString:       "TkString",
	TkNumber:       "TkNumber",
	TkIdent:        "TkIdent",
	TkEOF:          "TkEOF",
}

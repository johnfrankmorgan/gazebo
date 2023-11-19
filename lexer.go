package gazebo

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

type Lexer struct {
	source   string
	ast      *AST
	position struct {
		current, last ASTPosition
	}
	rules []struct {
		token  int
		regexp *regexp.Regexp
	}
	err error
}

var _ yyLexer = (*Lexer)(nil)

func NewLexer(source string, ast *AST) *Lexer {
	l := &Lexer{
		source: source,
		ast:    ast,
	}

	l.position.current = ASTPosition{Line: 1, Column: 1}
	l.position.last = l.position.current

	l.rules = []struct {
		token  int
		regexp *regexp.Regexp
	}{
		{TKDot, l.lit(`.`)},
		{TKSemicolon, l.lit(`;`)},
		{TKComma, l.lit(`,`)},
		{TKComment, l.exp(`//.*`)},
		{TKBraceOpen, l.lit(`{`)},
		{TKBraceClose, l.lit(`}`)},
		{TKParenOpen, l.lit(`(`)},
		{TKParenClose, l.lit(`)`)},
		{TKPlus, l.lit(`+`)},
		{TKMinus, l.lit(`-`)},
		{TKStar, l.lit(`*`)},
		{TKSlash, l.lit(`/`)},
		{TKPercent, l.lit(`%`)},
		{TKEqualEqual, l.lit(`==`)},
		{TKEqual, l.lit(`=`)},
		{TKBangEqual, l.lit(`!=`)},
		{TKBang, l.lit(`!`)},
		{TKLessEqual, l.lit(`<=`)},
		{TKLess, l.lit(`<`)},
		{TKGreaterEqual, l.lit(`>=`)},
		{TKGreater, l.lit(`>`)},
		{TKAnd, l.lit(`and`)},
		{TKOr, l.lit(`or`)},
		{TKIf, l.lit(`if`)},
		{TKElse, l.lit(`else`)},
		{TKWhile, l.lit(`while`)},
		{TKFunc, l.lit(`func`)},
		{TKLambda, l.lit(`lambda`)},
		{TKReturn, l.lit(`return`)},
		{TKReturn, l.lit(`continue`)},
		{TKReturn, l.lit(`break`)},
		{TKInteger, l.exp(`[0-9]+`)},
		{TKString, l.exp(`"(?:\\.|[^"])*"`)},
		{TKIdentifier, l.exp(`[a-zA-Z_][a-zA-Z0-9_]*`)},
	}

	return l
}

func (l *Lexer) exp(s string) *regexp.Regexp { return regexp.MustCompile("^" + s) }

func (l *Lexer) lit(s string) *regexp.Regexp { return l.exp(regexp.QuoteMeta(s)) }

func (l *Lexer) Lex(lval *yySymType) int {
	const eof = 0

	for l.position.current.Offset < len(l.source) && unicode.IsSpace(rune(l.source[l.position.current.Offset])) {
		if l.source[l.position.current.Offset] == '\n' {
			l.position.current.Line++
			l.position.current.Column = 0
		}

		l.position.current.Column++
		l.position.current.Offset++
	}

	lval.Position = l.position.current
	l.position.last = l.position.current

	for _, rule := range l.rules {
		if rule.regexp.MatchString(l.source[l.position.current.Offset:]) {
			lval.Lexeme = rule.regexp.FindString(l.source[l.position.current.Offset:])

			l.position.current.Column += len(lval.Lexeme)
			l.position.current.Offset += len(lval.Lexeme)

			return rule.token
		}
	}

	return eof
}

func (l *Lexer) Error(s string) {
	l.err = errors.Join(fmt.Errorf("gazebo: %s: %s", l.position.last, s))
}

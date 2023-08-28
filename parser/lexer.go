package parser

import (
	"errors"
	"fmt"
	"gazebo/ast"
	"regexp"
	"unicode"
)

type lexer struct {
	source   string
	position struct {
		last    ast.Position
		current ast.Position
	}
	program *ast.Program
	err     error
}

var _ yyLexer = (*lexer)(nil)

func (l *lexer) Lex(lval *yySymType) int {
	const eof = 0

	for len(l.source) > 0 && unicode.IsSpace(rune(l.source[0])) {
		if l.source[0] == '\n' {
			l.position.current.Line++
			l.position.current.Column = 0
		}

		l.position.current.Column++
		l.position.current.Offset++

		l.source = l.source[1:]
	}

	lval.Position = l.position.current
	l.position.last = l.position.current

	for _, rule := range rules {
		if rule.regexp.MatchString(l.source) {
			lval.Lexeme = rule.regexp.FindString(l.source)

			l.position.current.Column += len(lval.Lexeme)
			l.position.current.Offset += len(lval.Lexeme)

			l.source = l.source[len(lval.Lexeme):]

			return rule.token
		}
	}

	return eof
}

func (l *lexer) Error(s string) {
	l.err = errors.Join(fmt.Errorf("parser: %s: %s", l.position.last, s))
}

func exp(s string) *regexp.Regexp { return regexp.MustCompile("^" + s) }

func lit(s string) *regexp.Regexp { return exp(regexp.QuoteMeta(s)) }

var rules = []struct {
	token  int
	regexp *regexp.Regexp
}{
	{SEMICOLON, lit(`;`)},
	{COMMA, lit(`,`)},
	{COMMENT, exp(`//.*`)},
	{BRACE_OPEN, lit(`{`)},
	{BRACE_CLOSE, lit(`}`)},
	{PAREN_OPEN, lit(`(`)},
	{PAREN_CLOSE, lit(`)`)},
	{BRACKET_OPEN, lit(`[`)},
	{BRACKET_CLOSE, lit(`]`)},
	{PLUS, lit(`+`)},
	{MINUS, lit(`-`)},
	{STAR, lit(`*`)},
	{SLASH, lit(`/`)},
	{PERCENT, lit(`%`)},
	{EQUAL_EQUAL, lit(`==`)},
	{EQUAL, lit(`=`)},
	{BANG_EQUAL, lit(`!=`)},
	{BANG, lit(`!`)},
	{LESS_EQUAL, lit(`<=`)},
	{LESS, lit(`<`)},
	{GREATER_EQUAL, lit(`>=`)},
	{GREATER, lit(`>`)},
	{NULL, lit(`null`)},
	{FALSE, lit(`false`)},
	{TRUE, lit(`true`)},
	{AND, lit(`and`)},
	{IF, lit(`if`)},
	{ELSE, lit(`else`)},
	{WHILE, lit(`while`)},
	{OR, lit(`or`)},
	{RETURN, lit(`return`)},
	{FLOAT, exp(`[0-9]*\.[0-9]+`)},
	{INTEGER, exp(`[0-9]+`)},
	{STRING, exp(`("(?:\\.|[^"])*"|'(?:\\.|[^'])*')`)},
	{IDENTIFIER, exp(`[a-zA-Z_][a-zA-Z0-9_]*`)},
}

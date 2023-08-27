package parser

import (
	"fmt"
	"gazebo/ast"
	"os"
	"regexp"
	"unicode"
)

//go:generate goyacc grammar.y

func Parse(source string) *ast.Program {
	p := &parser{
		source:   source,
		position: ast.Position{Line: 1, Column: 1},
		program:  &ast.Program{Source: source},
	}

	yyParse(p)

	return p.program
}

type parser struct {
	source   string
	position ast.Position
	program  *ast.Program
}

var _ yyLexer = (*parser)(nil)

func (p *parser) Lex(lval *yySymType) int {
	const eof = 0

	for len(p.source) > 0 && unicode.IsSpace(rune(p.source[0])) {
		if p.source[0] == '\n' {
			p.position.Line++
			p.position.Column = 0
		}

		p.position.Column++
		p.position.Offset++

		p.source = p.source[1:]
	}

	lval.Position = p.position

	for _, rule := range rules {
		if rule.regexp.MatchString(p.source) {
			lval.Lexeme = rule.regexp.FindString(p.source)

			p.position.Column += len(lval.Lexeme)
			p.position.Offset += len(lval.Lexeme)

			p.source = p.source[len(lval.Lexeme):]

			return rule.token
		}
	}

	return eof
}

func (*parser) Error(s string) { fmt.Fprintln(os.Stderr, s) }

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

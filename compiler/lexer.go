package compiler

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
)

func tokenize(source string) tokens {
	var tokens tokens

	lexer := lexer{source: []byte(strings.TrimSpace(source))}

	for {
		tk := lexer.lex()

		if !tk.is(tkwhitespace, tkcomment, tknewline) {
			tokens = append(tokens, tk)
		}

		if tk.is(tkeof) {
			break
		}
	}

	return tokens
}

type tokens []token

func (m tokens) dump() {
	for idx, tk := range m {
		debug.Printf(
			"%6d: %02x :: %16s :: %q\n",
			idx,
			int(tk.typ),
			tk.typ.name(),
			tk.value,
		)
	}
}

type tokentype int

const (
	tkinvalid tokentype = iota
	tkcomment
	tknewline
	tkwhitespace
	tksemicolon
	tkdot
	tkcomma
	tkif
	tkelse
	tkreturn
	tkwhile
	tkfor
	tkbreak
	tkcontinue
	tkfun
	tkdel
	tkload
	tkpass
	tkin
	tkand
	tkor
	tkparenopen
	tkparenclose
	tkbraceopen
	tkbraceclose
	tkbracketopen
	tkbracketclose
	tkequal
	tkequalequal
	tkbang
	tkbangequal
	tkplus
	tkplusequal
	tkminus
	tkminusequal
	tkstar
	tkstarequal
	tkslash
	tkslashequal
	tkless
	tklessequal
	tkgreater
	tkgreaterequal
	tkstring
	tknumber
	tkident
	tkeof
)

func (m tokentype) name() string {
	names := map[tokentype]string{
		tkinvalid:      "tkinvalid",
		tkcomment:      "tkcomment",
		tknewline:      "tknewline",
		tkwhitespace:   "tkwhitespace",
		tksemicolon:    "tksemicolon",
		tkdot:          "tkdot",
		tkcomma:        "tkcomma",
		tkif:           "tkif",
		tkelse:         "tkelse",
		tkreturn:       "tkreturn",
		tkwhile:        "tkwhile",
		tkfor:          "tkfor",
		tkbreak:        "tkbreak",
		tkcontinue:     "tkcontinue",
		tkfun:          "tkfun",
		tkdel:          "tkdel",
		tkload:         "tkload",
		tkpass:         "tkpass",
		tkin:           "tkin",
		tkand:          "tkand",
		tkor:           "tkor",
		tkparenopen:    "tkparenopen",
		tkparenclose:   "tkparenclose",
		tkbraceopen:    "tkbraceopen",
		tkbraceclose:   "tkbraceclose",
		tkbracketopen:  "tkbracketopen",
		tkbracketclose: "tkbracketclose",
		tkequal:        "tkequal",
		tkequalequal:   "tkequalequal",
		tkbang:         "tkbang",
		tkbangequal:    "tkbangequal",
		tkplus:         "tkplus",
		tkplusequal:    "tkplusequal",
		tkminus:        "tkminus",
		tkminusequal:   "tkminusequal",
		tkstar:         "tkstar",
		tkstarequal:    "tkstarequal",
		tkslash:        "tkslash",
		tkslashequal:   "tkslashequal",
		tkless:         "tkless",
		tklessequal:    "tklessequal",
		tkgreater:      "tkgreater",
		tkgreaterequal: "tkgreaterequal",
		tkstring:       "tkstring",
		tknumber:       "tknumber",
		tkident:        "tkident",
		tkeof:          "tkeof",
	}

	if name, ok := names[m]; ok {
		return name
	}

	return "tkunknown"
}

var keywords = map[string]tokentype{
	"if":       tkif,
	"else":     tkelse,
	"return":   tkreturn,
	"while":    tkwhile,
	"for":      tkfor,
	"break":    tkbreak,
	"continue": tkcontinue,
	"fun":      tkfun,
	"del":      tkdel,
	"load":     tkload,
	"pass":     tkpass,
	"in":       tkin,
	"and":      tkand,
	"or":       tkor,
}

type token struct {
	typ   tokentype
	value string
}

func (m token) is(types ...tokentype) bool {
	for _, typ := range types {
		if m.typ == typ {
			return true
		}
	}

	return false
}

type lexer struct {
	source   []byte
	position int
	buffer   bytes.Buffer
}

func (m *lexer) unexpectedeof() token {
	errors.ErrEOF.Panic("unexpected eof at byte offset %d, %v", m.position, m.peek())
	return m.token(tkinvalid)
}

func (m *lexer) finished() bool {
	return m.position >= len(m.source)
}

func (m *lexer) peek() rune {
	ch, _ := utf8.DecodeRune(m.source[m.position:])
	return ch
}

func (m *lexer) next() rune {
	ch, width := utf8.DecodeRune(m.source[m.position:])
	m.buffer.WriteRune(ch)
	m.position += width
	return ch
}

func (m *lexer) match(ch rune) bool {
	if m.finished() {
		return false
	}

	if m.peek() == ch {
		m.next()
		return true
	}

	return false
}

func (m *lexer) token(typ tokentype) token {
	tk := token{typ: typ, value: m.buffer.String()}
	m.buffer.Reset()
	return tk
}

func (m *lexer) isdigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (m *lexer) isalpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (m *lexer) isidentchar(ch rune) bool {
	if ch >= 0x1f600 { // >= 😀
		return true
	}

	if m.isalpha(ch) || m.isdigit(ch) {
		return true
	}

	for _, identch := range "!?@_$" {
		if identch == ch {
			return true
		}
	}

	return false
}

func (m *lexer) isnewline(ch rune) bool {
	return ch == '\n'
}

func (m *lexer) iswhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func (m *lexer) line(typ tokentype) token {
	for !m.finished() {
		if m.isnewline(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(typ)
}

func (m *lexer) lstring() token {
	for !m.finished() {
		if m.match('"') {
			return m.token(tkstring)
		}

		m.next()
	}

	return m.unexpectedeof()
}

func (m *lexer) lnumber() token {
	var isfloat bool

	for !m.finished() {
		ch := m.peek()

		if ch == '.' && !isfloat {
			m.next()
			isfloat = true
			continue
		}

		if !m.isdigit(ch) {
			break
		}

		m.next()
	}

	return m.token(tknumber)
}

func (m *lexer) lident() token {
	for !m.finished() {
		if !m.isidentchar(m.peek()) {
			break
		}

		m.next()
	}

	if typ, ok := keywords[m.buffer.String()]; ok {
		return m.token(typ)
	}

	return m.token(tkident)
}

func (m *lexer) lwhitespace() token {
	for !m.finished() {
		if !m.iswhitespace(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(tkwhitespace)
}

func (m *lexer) ifmatch(ch rune, typ, fallback tokentype) token {
	if m.match(ch) {
		return m.token(typ)
	}

	return m.token(fallback)
}

func (m *lexer) lex() token {
	if m.finished() {
		return m.token(tkeof)
	}

	ch := m.next()

	switch ch {
	case ';':
		return m.token(tksemicolon)

	case '.':
		return m.token(tkdot)

	case ',':
		return m.token(tkcomma)

	case '(':
		return m.token(tkparenopen)

	case ')':
		return m.token(tkparenclose)

	case '{':
		return m.token(tkbraceopen)

	case '}':
		return m.token(tkbraceclose)

	case '[':
		return m.token(tkbracketopen)

	case ']':
		return m.token(tkbracketclose)

	case '+':
		return m.ifmatch('=', tkplusequal, tkplus)

	case '-':
		return m.ifmatch('=', tkminusequal, tkminus)

	case '*':
		return m.ifmatch('=', tkstarequal, tkstar)

	case '/':
		if m.match('/') {
			return m.line(tkcomment)
		}

		return m.ifmatch('=', tkslashequal, tkslash)

	case '=':
		return m.ifmatch('=', tkequalequal, tkequal)

	case '!':
		return m.ifmatch('=', tkbangequal, tkbang)

	case '<':
		return m.ifmatch('=', tklessequal, tkless)

	case '>':
		return m.ifmatch('=', tkgreaterequal, tkgreater)

	case '"':
		return m.lstring()
	}

	if m.isnewline(ch) {
		return m.token(tknewline)
	}

	if m.isdigit(ch) {
		return m.lnumber()
	}

	if m.isidentchar(ch) {
		return m.lident()
	}

	if m.iswhitespace(ch) {
		return m.lwhitespace()
	}

	errors.ErrParse.Panic("unexpected rune %c at byte offset %d", ch, m.position)
	return m.token(tkinvalid)
}

package parser

import (
	"bytes"
	"fmt"
)

type Lexer struct {
	source   string
	position int
	buffer   bytes.Buffer
}

func (m *Lexer) token(kind TKind) Token {
	defer m.buffer.Reset()

	return Token{
		kind:     kind,
		lexeme:   m.buffer.String(),
		position: m.position,
	}
}

func (m *Lexer) isdigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (m *Lexer) isnewline(ch byte) bool {
	return ch == '\n'
}

func (m *Lexer) iswhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || m.isnewline(ch)
}

func (m *Lexer) isalpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (m *Lexer) isident(ch byte) bool {
	return ch == '_' || m.isalpha(ch)
}

func (m *Lexer) finished() bool {
	return m.position >= len(m.source)
}

func (m *Lexer) advance() {
	m.position++
}

func (m *Lexer) peek() byte {
	if m.finished() {
		return 0
	}

	return m.source[m.position]
}

func (m *Lexer) next() byte {
	defer m.advance()

	ch := m.peek()
	m.buffer.WriteByte(ch)

	return ch
}

func (m *Lexer) match(ch byte) bool {
	if m.peek() == ch {
		m.next()
		return true
	}

	return false
}

func (m *Lexer) line(kind TKind) Token {
	for !m.finished() {
		if m.isnewline(m.peek()) {
			m.next()
			break
		}

		m.next()
	}

	return m.token(kind)
}

func (m *Lexer) whitespace() Token {
	for !m.finished() {
		if !m.iswhitespace(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(TWhitespace)
}

func (m *Lexer) number() Token {
	for !m.finished() {
		if !m.isdigit(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(TNumber)
}

func (m *Lexer) ident() Token {
	keywords := map[string]TKind{
		"if":     TIf,
		"else":   TElse,
		"while":  TWhile,
		"return": TReturn,
	}

	for !m.finished() {
		if !m.isident(m.peek()) {
			break
		}

		m.next()
	}

	if kind, ok := keywords[m.buffer.String()]; ok {
		return m.token(kind)
	}

	return m.token(TIdent)
}

func (m *Lexer) lex() Token {
	if m.finished() {
		return m.token(TEOF)
	}

	ch := m.next()

	switch ch {
	case '#':
		return m.line(TComment)

	case '(':
		return m.token(TParenOpen)

	case ')':
		return m.token(TParenClose)

	case '{':
		return m.token(TBraceOpen)

	case '}':
		return m.token(TBraceClose)

	case '.':
		return m.token(TDot)

	case ',':
		return m.token(TComma)

	case ';':
		return m.token(TSemicolon)

	case '!':
		if m.match('=') {
			return m.token(TBangEqual)
		}

		return m.token(TBang)

	case '=':
		if m.match('=') {
			return m.token(TEqualEqual)
		}

		return m.token(TEqual)

	case '<':
		if m.match('=') {
			return m.token(TLessEqual)
		}

		return m.token(TLess)

	case '>':
		if m.match('=') {
			return m.token(TGreaterEqual)
		}

		return m.token(TGreater)

	case '+':
		return m.token(TPlus)

	case '-':
		return m.token(TMinus)

	case '*':
		return m.token(TStar)

	case '/':
		return m.token(TSlash)
	}

	if m.iswhitespace(ch) {
		return m.whitespace()
	}

	if m.isdigit(ch) {
		return m.number()
	}

	if m.isident(ch) {
		return m.ident()
	}

	panic(
		fmt.Errorf(
			"failed to lex token at position %d near %q",
			m.position,
			m.source[m.position-3:m.position+3], // FIXME: add bounds check
		),
	)
}

func (m *Lexer) Lex() []Token {
	var tokens []Token

	for !m.finished() {
		tokens = append(tokens, m.lex())
	}

	return append(tokens, m.token(TEOF))
}

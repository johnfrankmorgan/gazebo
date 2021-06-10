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
		if m.match('\n') {
			break
		}

		m.next()
	}

	return m.token(kind)
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

	panic(
		fmt.Errorf(
			"failed to lex token at position %d near %q",
			m.position,
			m.buffer.String(),
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

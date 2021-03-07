package lexer

import (
	"bytes"
	"unicode/utf8"
)

type Lexer struct {
	source   []byte
	buffer   bytes.Buffer
	position int
}

func New(source []byte) *Lexer {
	return &Lexer{source: source}
}

func (m *Lexer) Lex() (tokens Tokens, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			if perr, ok := perr.(error); ok {
				err = perr
			}
		}
	}()

	for {
		token := m.lex()

		if !token.Is(TkWhitespace, TkComment, TkNewline) {
			tokens = append(tokens, token)
		}

		if token.Is(TkEOF) {
			break
		}
	}

	return tokens, nil
}

func (m *Lexer) lex() Token {
	if m.finished() {
		return m.token(TkEOF)
	}

	ch := m.next()

	switch ch {
	case ':':
		return m.token(TkColon)

	case ';':
		return m.token(TkSemicolon)

	case '.':
		return m.token(TkDot)

	case ',':
		return m.token(TkComma)

	case '(':
		return m.token(TkParenOpen)

	case ')':
		return m.token(TkParenClose)

	case '{':
		return m.token(TkBraceOpen)

	case '}':
		return m.token(TkBraceClose)

	case '[':
		return m.token(TkBracketOpen)

	case ']':
		return m.token(TkBracketClose)

	case '+':
		return m.ifmatch('=', TkPlusEqual, TkPlus)

	case '-':
		return m.ifmatch('=', TkMinusEqual, TkMinus)

	case '*':
		return m.ifmatch('=', TkStarEqual, TkStar)

	case '/':
		if m.match('/') {
			return m.line(TkComment)
		}

		return m.ifmatch('=', TkSlashEqual, TkSlash)

	case '=':
		return m.ifmatch('=', TkEqualEqual, TkEqual)

	case '!':
		return m.ifmatch('=', TkBangEqual, TkBang)

	case '<':
		return m.ifmatch('=', TkLessEqual, TkLess)

	case '>':
		return m.ifmatch('=', TkGreaterEqual, TkGreater)

	case '"':
		return m.lstring()

	case '#':
		return m.line(TkComment)

	}

	if isNewline(ch) {
		return m.token(TkNewline)
	}

	if isDigit(ch) {
		return m.lnumber()
	}

	if isIdentChar(ch) {
		return m.lident()
	}

	if isWhitespace(ch) {
		return m.lwhitespace()
	}

	panic(ErrUnexpectedRune.WithContext(m.source, m.position))
}

func (m *Lexer) token(typ TokenType) Token {
	defer m.buffer.Reset()

	return Token{
		Type:     typ,
		Value:    m.buffer.String(),
		Position: m.position - m.buffer.Len(),
	}
}

func (m *Lexer) finished() bool {
	return m.position >= len(m.source)
}

func (m *Lexer) next() rune {
	if m.finished() {
		return 0
	}

	ch, width := utf8.DecodeRune(m.source[m.position:])
	m.buffer.WriteRune(ch)
	m.position += width

	return ch
}

func (m *Lexer) peek(offset ...int) rune {
	if m.finished() {
		return 0
	}

	if len(offset) == 0 {
		offset = []int{0}
	}

	ch, _ := utf8.DecodeRune(m.source[m.position+offset[0]:])
	return ch
}

func (m *Lexer) match(ch rune) bool {
	if m.peek() == ch {
		m.next()
		return true
	}

	return false
}

func (m *Lexer) ifmatch(ch rune, typ, fallback TokenType) Token {
	if m.match(ch) {
		return m.token(typ)
	}

	return m.token(fallback)
}

func (m *Lexer) line(t TokenType) Token {
	for !isNewline(m.peek()) {
		m.next()
	}

	return m.token(t)
}

func (m *Lexer) lstring() Token {
	for !m.finished() {
		if m.match('"') {
			return m.token(TkString)
		}

		m.next()
	}

	panic(ErrUnexpectedEOF.WithContext(m.source, m.position))
}

func (m *Lexer) lnumber() Token {
	var isfloat bool

	for !m.finished() {
		if !isfloat && m.peek() == '.' && isDigit(m.peek(1)) {
			m.next()
			m.next()
			isfloat = true
			continue
		}

		if !isDigit(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(TkNumber)
}

func (m *Lexer) lident() Token {
	for isIdentChar(m.peek()) {
		m.next()
	}

	if t, ok := keywords[m.buffer.String()]; ok {
		return m.token(t)
	}

	return m.token(TkIdent)
}

func (m *Lexer) lwhitespace() Token {
	for !m.finished() {
		if !isWhitespace(m.peek()) {
			break
		}

		m.next()
	}

	return m.token(TkWhitespace)
}

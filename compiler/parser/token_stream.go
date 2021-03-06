package parser

import "github.com/johnfrankmorgan/gazebo/compiler/lexer"

type TokenStream struct {
	tokens   lexer.Tokens
	position int
}

func NewTokenStream(tokens lexer.Tokens) *TokenStream {
	return &TokenStream{tokens: tokens}
}

func (m *TokenStream) Position() int {
	return m.position
}

func (m *TokenStream) Advance() {
	m.position++
}

func (m *TokenStream) Reset() {
	m.position = 0
}

func (m *TokenStream) Finished() bool {
	return m.position >= len(m.tokens)
}

func (m *TokenStream) Prev() lexer.Token {
	return m.tokens[m.position-1]
}

func (m *TokenStream) Peek(offset ...int) lexer.Token {
	if len(offset) == 0 {
		offset = []int{0}
	}

	return m.tokens[m.position+offset[0]]
}

func (m *TokenStream) Next() lexer.Token {
	if m.Finished() {
		return m.tokens[len(m.tokens)-1]
	}

	defer m.Advance()

	return m.Peek()
}

func (m *TokenStream) Check(types ...lexer.TokenType) bool {
	if m.Finished() {
		return false
	}

	for _, t := range types {
		if m.Peek().Is(t) {
			return true
		}
	}

	return false
}

func (m *TokenStream) Match(types ...lexer.TokenType) bool {
	if m.Check(types...) {
		m.Next()
		return true
	}

	return false
}

func (m *TokenStream) Expect(types ...lexer.TokenType) lexer.Token {
	if !m.Match(types...) {
		panic(UnexpectedToken(m.Peek(), types...))
	}

	return m.Prev()
}

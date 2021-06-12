package parser

import "fmt"

type tstream struct {
	tokens   []Token
	position int
}

func (m *tstream) len() int {
	return len(m.tokens)
}

func (m *tstream) finished() bool {
	return m.position >= m.len() || m.tokens[m.position].Is(TEOF)
}

func (m *tstream) advance() {
	m.position++
}

func (m *tstream) peek(n int) Token {
	defer func(position int) {
		m.position = position
	}(m.position)

	token := m.next()

	for i := 0; i < n; i++ {
		token = m.next()
	}

	return token
}

func (m *tstream) prev() Token {
	return m.tokens[m.position-1]
}

func (m *tstream) next() Token {
	for {
		token := m.tokens[m.position]
		m.advance()

		if !token.Is(TWhitespace, TComment) {
			return token
		}
	}
}

func (m *tstream) check(kinds ...TKind) bool {
	return m.peek(0).Is(kinds...)
}

func (m *tstream) match(kinds ...TKind) bool {
	if m.check(kinds...) {
		m.next()
		return true
	}

	return false
}

func (m *tstream) consume(kinds ...TKind) Token {
	if m.match(kinds...) {
		return m.prev()
	}

	panic(
		fmt.Errorf(
			"expected one of %s, got %s",
			kinds,
			m.peek(0).kind,
		),
	)
}

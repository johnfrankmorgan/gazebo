package lexer

import "fmt"

var (
	ErrUnexpectedEOF  = &Error{message: "unexpected EOF"}
	ErrUnexpectedRune = &Error{message: "unexpected character"}
)

type Error struct {
	message string
	context string
}

func (m *Error) WithContext(source []byte, position int) *Error {
	m.context = string(limit(source, position, 10))
	return m
}

func (m *Error) Error() string {
	str := fmt.Sprintf("lexer: %s", m.message)

	if m.context != "" {
		str = fmt.Sprintf("%s: %s <<", str, m.context)
	}

	return str
}

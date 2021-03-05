package parser

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
)

type Error struct {
	message string
}

func UnexpectedToken(token lexer.Token, expected []lexer.TokenType) error {
	return &Error{
		message: fmt.Sprintf("unexpected token %s, expected %s", token, expected),
	}
}

func (m *Error) Error() string {
	return fmt.Sprintf("parser: %s", m.message)
}

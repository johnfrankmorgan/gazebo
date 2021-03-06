package parser

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type Error struct {
	message string
}

func UnexpectedToken(token lexer.Token, expected ...lexer.TokenType) error {
	return &Error{
		message: fmt.Sprintf("unexpected token %s, expected %s", token, expected),
	}
}

func UnexpectedEOF() error {
	return &Error{message: "unexpected EOF"}
}

func UnexpectedExpression(expr expr.Expression) error {
	return &Error{
		message: fmt.Sprintf("unexpected expression: %v", expr),
	}
}

func (m *Error) Error() string {
	return fmt.Sprintf("parser: %s", m.message)
}

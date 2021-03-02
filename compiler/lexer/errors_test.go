package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert := assert.New(t)

	str := []byte("if 1 == 1 ' testing")
	err := ErrUnexpectedRune.WithContext(str, 10)
	exp := "lexer: unexpected character: if 1 == 1 ' <<"

	assert.Equal(exp, err.Error())
}

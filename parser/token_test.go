package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenToBinOpPanics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		t := Token{kind: TIf}
		t.ToBinOp()
	})
}

func TestTokenToUnaryOpPanics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		t := Token{kind: TIf}
		t.ToUnaryOp()
	})
}

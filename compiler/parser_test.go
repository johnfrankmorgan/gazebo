package compiler

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)

	source, err := ioutil.ReadFile("../tests/gaz/test-fixture-1.gaz")
	assert.Nil(err)

	parser := parser{tokens: tokenize(string(source))}
	assert.Len(parser.parse(), 9)
}

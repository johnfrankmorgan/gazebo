package compiler

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {
	assert := assert.New(t)

	source, err := ioutil.ReadFile("../tests/gaz/test-fixture-1.gaz")
	assert.Nil(err)

	_, err = Compile(string(source))
	assert.Nil(err)
}

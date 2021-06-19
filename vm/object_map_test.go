package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)

	m := NewMap()
	m.Set(NewString("test"), NewNumber(123.1))

	assert.Len(m.values, 1)
	assert.Len(m.keys, 1)
	assert.True(m.Has(NewString("test")))
	assert.Equal(NewNumber(123.1), m.Get(NewString("test")))
	assert.Equal(NewNil(), m.Get(NewString("other")))

	m.Del(NewString("test"))
	assert.Len(m.values, 0)
	assert.Len(m.keys, 0)

	assert.Same(m, m.Attrs())
}

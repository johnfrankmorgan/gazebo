package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsExpectsExactly(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		var args Args
		args.ExpectsExactly(10)
	})
}

func TestArgsParse(t *testing.T) {
	assert := assert.New(t)

	args := Args{
		NewString("1"),
		NewString("2"),
		NewBool(true),
		NewBool(false),
		NewNumber(1.1),
		NewNumber(20.2),
		NewNumber(300.0),
		NewMap(),
	}

	var got struct {
		s1 *String
		s2 string
		b1 *Bool
		b2 bool
		n1 *Number
		n2 float64
		n3 int
		o1 Object
	}

	args.Parse(
		&got.s1,
		&got.s2,
		&got.b1,
		&got.b2,
		&got.n1,
		&got.n2,
		&got.n3,
		&got.o1,
	)

	assert.Equal("1", got.s1.String())
	assert.Equal("2", got.s2)
	assert.Equal(true, got.b1.Bool())
	assert.Equal(false, got.b2)
	assert.Equal(1.1, got.n1.Float())
	assert.Equal(20.2, got.n2)
	assert.Equal(300, got.n3)
	assert.Equal(NewMap(), got.o1)
}

func TestArgsParsePanics(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		var t struct{}

		args := Args{NewString("test")}
		args.Parse(&t)
	})
}

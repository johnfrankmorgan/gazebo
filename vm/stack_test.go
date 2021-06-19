package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push(NewBool(true))
	assert.Equal(t, 1, stack.Size())
}

func TestStackPop(t *testing.T) {
	stack := NewStack()
	stack.Push(NewString("test"))
	stack.Push(NewNil())
	got := stack.Pop()
	assert.Equal(t, NewNil(), got)
	assert.Equal(t, 1, stack.Size())
}

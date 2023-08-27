package ds_test

import (
	"gazebo/util/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Parallel()

	s := ds.NewStack[int]()

	t.Run("push", func(t *testing.T) {
		s.Push(1)
		s.Push(2)
		s.Push(3)

		assert.Equal(t, 3, s.Size())
	})

	t.Run("pop", func(t *testing.T) {
		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 2, s.Size())
	})

	t.Run("peek", func(t *testing.T) {
		assert.Equal(t, 2, s.Peek())
	})
}

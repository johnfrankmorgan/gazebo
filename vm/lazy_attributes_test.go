package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyAttribtutesAttrs(t *testing.T) {
	var attrs LazyAttributes

	assert := assert.New(t)
	assert.NotNil(attrs.Attrs())
}

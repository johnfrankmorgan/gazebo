package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithError("parse error: oh no!", func() {
		ErrParse.WithMessage("won't panic").Expect(true)
		ErrParse.WithMessage("won't panic").ExpectAtLeast([]string{"test"}, 0)
		ErrParse.WithMessage("won't panic").ExpectLen([]string{}, 0)
		ErrParse.WithMessage("won't panic").ExpectNil(nil)
		ErrParse.WithMessage("won't panic").ExpectNilError(nil)
		ErrParse.WithMessage("oh no!").Expect(false)
	})
}

package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTStreamNext(t *testing.T) {
	assert := assert.New(t)

	ts := tstream{tokens: Tokenize("    x #c\n123")}

	assert.True(ts.next().Is(TIdent))
	assert.True(ts.next().Is(TNumber))
}

func TestTStreamPrev(t *testing.T) {
	assert := assert.New(t)

	ts := tstream{tokens: Tokenize("if x()")}

	ts.next()
	assert.True(ts.prev().Is(TIf))
}

func TestTStreamPeek(t *testing.T) {
	assert := assert.New(t)

	ts := tstream{tokens: Tokenize("    x ")}

	assert.True(ts.peek(0).Is(TIdent))
	assert.True(ts.peek(1).Is(TEOF))
	assert.Equal(0, ts.position)
}

func TestTStreamMatch(t *testing.T) {
	assert := assert.New(t)

	ts := tstream{tokens: Tokenize("1")}

	assert.True(ts.match(TNumber))
	assert.False(ts.match(TIf))
	assert.Equal(1, ts.position)
}

func TestTStreamFinished(t *testing.T) {
	var ts tstream

	assert.True(t, ts.finished())
}

func TestTStreamConsumePanics(t *testing.T) {
	assert := assert.New(t)

	ts := tstream{tokens: Tokenize("test")}

	assert.Panics(func() {
		ts.consume(TIf)
	})
}

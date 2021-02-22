package compiler

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/protocols"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	assert := assert.New(t)

	debug.Enable()
	defer debug.Disable()

	source := `
		// this is a comment
		if (true) {
			return "test"[0.0];
		}
		! = != == > >= < <= + - * /
	`

	expected := []tokentype{
		tkif,
		tkparenopen,
		tkident,
		tkparenclose,
		tkbraceopen,
		tkreturn,
		tkstring,
		tkbracketopen,
		tknumber,
		tkbracketclose,
		tksemicolon,
		tkbraceclose,
		tkbang,
		tkequal,
		tkbangequal,
		tkequalequal,
		tkgreater,
		tkgreaterequal,
		tkless,
		tklessequal,
		tkplus,
		tkminus,
		tkstar,
		tkslash,
		tkeof,
	}

	tokenize(source).dump()

	got := []tokentype{}
	for _, token := range tokenize(source) {
		got = append(got, token.typ)
	}

	assert.Equal(expected, got)
}

func TestProtocolMethodsAreValidIdentifiers(t *testing.T) {
	for _, protocol := range protocols.All() {
		t.Run(protocol, func(t *testing.T) {
			assert := assert.New(t)
			tokens := tokenize(protocol)

			assert.Len(tokens, 2)
			assert.True(tokens[0].is(tkident))
			assert.True(tokens[1].is(tkeof))
		})
	}
}

package compiler

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	assert := assert.New(t)

	debug.Enable()
	defer debug.Disable()

	source := `
		# this is a comment
		if (true) {
			return "test"[0.0];
		}
		! = != == > >= < <= + - * /
	`

	expected := []tokentype{
		tkcomment,
		tknewline,
		tkif,
		tkparenopen,
		tkident,
		tkparenclose,
		tkbraceopen,
		tknewline,
		tkreturn,
		tkstring,
		tkbracketopen,
		tknumber,
		tkbracketclose,
		tksemicolon,
		tknewline,
		tkbraceclose,
		tknewline,
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

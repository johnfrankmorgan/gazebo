package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvResolve(t *testing.T) {
	env1 := NewEnv(nil, nil)
	env2 := NewEnv(nil, env1)

	env1.Assign("v1", NewString("v1"))
	env2.Assign("v2", NewString("v2"))

	env := env2

	assert := assert.New(t)

	assert.Same(env1, env.Resolve("v1"))
	assert.Same(env2, env.Resolve("v2"))
	assert.Nil(env.Resolve("flurb"))
}

func TestEnvDefined(t *testing.T) {
	env := NewEnv(nil, nil)
	env.Assign("test", NewString("test"))

	assert.True(t, env.Defined("test"))
}

func TestEnvAssign(t *testing.T) {
	parent := NewEnv(nil, nil)
	env := NewEnv(nil, parent)

	parent.Assign("test", NewBool(false))
	env.Assign("test", NewBool(true))

	assert.True(t, parent.Lookup("test").(*Bool).Bool())
}

func TestEnvLookupReturnsNil(t *testing.T) {
	env := NewEnv(nil, nil)
	assert.Equal(t, NewNil(), env.Lookup("test"))
}

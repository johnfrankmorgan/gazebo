package modules

import (
	"reflect"
	"strings"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

var strfuncs = map[string]interface{}{
	"cmp":       strings.Compare,
	"contains?": strings.Contains,
	"count":     strings.Count,
	"ends?":     strings.HasSuffix,
	"lower":     strings.ToLower,
	"pos":       strings.Index,
	"replace":   strings.ReplaceAll,
	"starts?":   strings.HasPrefix,
	"trim":      strings.TrimSpace,
	"upper":     strings.ToUpper,
}

func wrapfunction(f interface{}) g.Func {
	fun := reflect.ValueOf(f)

	assert.True(fun.Kind() == reflect.Func)

	return g.Func(func(args g.Args) g.Object {
		args.Expects(fun.Type().NumIn())

		converted := make([]reflect.Value, len(args))

		for i, arg := range args {
			str := g.EnsureString(arg.Call(protocols.String, nil)).String()
			converted[i] = reflect.ValueOf(str)
		}

		return g.NewObject(fun.Call(converted)[0].Interface())
	})
}

// Str holds the definitions for the str module
var Str = &Module{
	Name: "str",
	Init: func(m *Module) {
		for name, f := range strfuncs {
			m.Values[name] = g.NewObjectInternalFunc(wrapfunction(f))
		}
	},
	Values: map[string]g.Object{},
}

package vm

import "github.com/johnfrankmorgan/gazebo/runtime"

type Variables struct {
	Parent *Variables
	Values map[string]runtime.Object
}

func (v *Variables) init() {
	if len(v.Values) == 0 {
		v.Values = make(map[string]runtime.Object)
	}
}

func (v *Variables) Get(name string) (runtime.Object, bool) {
	if value, ok := v.Values[name]; ok {
		return value, true
	}

	if v.Parent != nil {
		return v.Parent.Get(name)
	}

	return runtime.Nil, false
}

func (v *Variables) Set(name string, value runtime.Object) {
	v.init()

	v.Values[name] = value
}

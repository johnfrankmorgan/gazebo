package vm

import (
	"gazebo/objects"
	"gazebo/util/ds"
)

type Variables struct {
	parent *Variables
	values *ds.Map[string, *objects.Object]
}

func NewVariables(parent *Variables) *Variables {
	return &Variables{
		parent: parent,
		values: ds.NewMap[string, *objects.Object](),
	}
}

func (s *Variables) Parent() *Variables {
	return s.parent
}

func (s *Variables) Resolve(name string) *Variables {
	for s := s; s != nil; s = s.Parent() {
		if s.values.Has(name) {
			return s
		}
	}

	return nil
}

func (s *Variables) Load(name string) *objects.Object {
	return s.values.Get(name)
}

func (s *Variables) Store(name string, value *objects.Object) {
	s.values.Set(name, value)
}

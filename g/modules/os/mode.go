package os

import (
	"os"

	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Mode{}

type Mode struct {
	g.Number
}

func NewMode(value os.FileMode) *Mode {
	object := &Mode{Number: *g.NewNumberFromInt(int(value))}
	object.SetSelf(object)
	return object
}

func NewModeFromNumber(value *g.Number) *Mode {
	return NewMode(os.FileMode(value.Int()))
}

func (m *Mode) Mode() os.FileMode {
	return os.FileMode(m.Int())
}

// GAZEBO MODE OBJECT PROTOCOLS

func (m *Mode) G_repr() *g.String {
	return m.Base.G_repr()
}

func (m *Mode) G_str() *g.String {
	return g.NewString(m.Mode().String())
}

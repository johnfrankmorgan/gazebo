package time

import (
	"time"

	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Time{}

type Time struct {
	g.Base
	value time.Time
}

func NewTime(value time.Time) *Time {
	object := &Time{value: value}
	object.SetSelf(object)
	object.SetAttr("default_format", g.NewString("2 Jan 2006 15:04:05"))
	return object
}

func (m *Time) Value() interface{} {
	return m.value
}

func (m *Time) Time() time.Time {
	return m.value
}

// GAZEBO TIME OBJECT PROTOCOLS

func (m *Time) G_str() *g.String {
	format := m.GetAttr("default_format").G_str()
	return m.G_format(format)
}

// GAZEBO TIME OBJECT METHODS

func (m *Time) G_format(format *g.String) *g.String {
	return g.NewString(m.value.Format(format.String()))
}

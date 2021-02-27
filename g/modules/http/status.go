package http

import (
	"net/http"

	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Status{}

type Status struct {
	g.Number
}

func NewStatus(value int) *Status {
	object := &Status{
		Number: *g.NewNumberFromInt(value),
	}
	object.SetSelf(object)
	return object
}

func (m *Status) G_str() *g.String {
	return m.G_name()
}

func (m *Status) G_name() *g.String {
	if name := http.StatusText(m.Int()); name != "" {
		return g.NewString(http.StatusText(m.Int()))
	}

	return m.Number.G_str()
}

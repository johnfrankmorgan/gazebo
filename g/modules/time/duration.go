package time

import (
	"time"

	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Duration{}

type Duration struct {
	g.Number
}

func NewDuration(value time.Duration) *Duration {
	object := &Duration{Number: *g.NewNumber(float64(value))}
	object.SetSelf(object)
	return object
}

func NewDurationFromObject(object g.Object) *Duration {
	return NewDuration(time.Duration(object.G_num().Int64()) * time.Millisecond)
}

func (m *Duration) Value() interface{} {
	return m.Duration()
}

func (m *Duration) Duration() time.Duration {
	return time.Duration(m.Int64())
}

// GAZEBO DURATION OBJECT PROTOCOLS

func (m *Duration) G_repr() *g.String {
	return m.Number.Base.G_repr()
}

func (m *Duration) G_str() *g.String {
	return g.NewString(m.Duration().String())
}

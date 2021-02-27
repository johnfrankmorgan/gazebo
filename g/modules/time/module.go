package time

import (
	"time"

	"github.com/johnfrankmorgan/gazebo/g"
)

type TimeModule struct {
	g.Base
}

func NewTimeModule() *TimeModule {
	object := &TimeModule{}
	object.SetSelf(object)
	return object
}

func (m *TimeModule) Value() interface{} {
	return m.Name()
}

func (m *TimeModule) Name() string {
	return "time"
}

// GAZEBO TIME MODULE OBJECT METHODS

func (m *TimeModule) G_sleep(object g.Object) {
	time.Sleep(NewDurationFromObject(object).Duration())
}

func (m *TimeModule) G_now() *Time {
	return NewTime(time.Now())
}

func (m *TimeModule) G_since(value *Time) *Duration {
	return NewDuration(time.Since(value.Time()))
}

func (m *TimeModule) G_timeit(cb g.Object, args ...g.Object) *Duration {
	now := m.G_now()
	cb.G_invoke(g.NewArgs(args))
	return m.G_since(now)
}

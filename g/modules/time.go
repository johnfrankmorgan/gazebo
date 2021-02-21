package modules

import (
	"strings"
	"time"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

// Types used by the time module
var (
	TypeTime     *g.Type
	TypeDuration *g.Type
)

// ObjectTime is the underlying value of Time objects in gazebo
type ObjectTime struct {
	g.PartialObject
	value time.Time
}

// NewObjectTime creates a new ObjectTime
func NewObjectTime(value time.Time) *ObjectTime {
	object := &ObjectTime{value: value}
	object.SetType(TypeTime)
	return object
}

// Value satisfies the g.Object interface
func (m *ObjectTime) Value() interface{} {
	return m.value
}

// Call satisfies the g.Object interface
func (m *ObjectTime) Call(method string, args g.Args) g.Object {
	return m.CallMethod(m, method, args)
}

// Time returns the ObjectTime's value
func (m *ObjectTime) Time() time.Time {
	return m.value
}

// ObjectDuration is the underlying value of Duration objects in gazebo
type ObjectDuration struct {
	g.PartialObject
	value time.Duration
}

// NewObjectDuration creates a new ObjectDuration
func NewObjectDuration(value time.Duration) *ObjectDuration {
	object := &ObjectDuration{value: value}
	object.SetType(TypeDuration)
	return object
}

// Value satisfies the g.Object interface
func (m *ObjectDuration) Value() interface{} {
	return m.value
}

// Call satisfies the g.Object interface
func (m *ObjectDuration) Call(method string, args g.Args) g.Object {
	return m.CallMethod(m, method, args)
}

// Duration returns the ObjectDuration's value
func (m *ObjectDuration) Duration() time.Duration {
	return m.value
}

// EnsureTime asserts a value is a time object
func EnsureTime(value g.Object) *ObjectTime {
	errors.ErrRuntime.Expect(
		value.Type() == TypeTime,
		"expected type Time got %s",
		value.Type().Name,
	)

	return value.(*ObjectTime)
}

// EnsureDuration asserts a value is a duration object
func EnsureDuration(value g.Object) *ObjectDuration {
	errors.ErrRuntime.Expect(
		value.Type() == TypeDuration,
		"expected type Duration got %s",
		value.Type().Name,
	)

	return value.(*ObjectDuration)
}

// Time holds the definitions for the time module
var Time = &Module{
	Name: "time",
	Init: func(_ *Module) {
		TypeTime = &g.Type{
			Name:   "time.Time",
			Parent: g.TypeBase,
			Methods: g.Methods{
				protocols.String: func(self g.Object, _ g.Args) g.Object {
					return g.NewObjectString(EnsureTime(self).Time().String())
				},

				protocols.Number: func(self g.Object, _ g.Args) g.Object {
					return g.NewObjectNumber(float64(EnsureTime(self).Time().Unix()))
				},

				protocols.Add: func(self g.Object, args g.Args) g.Object {
					return NewObjectTime(
						EnsureTime(self).Time().Add(EnsureDuration(args.Self()).Duration()),
					)
				},

				protocols.Sub: func(self g.Object, args g.Args) g.Object {
					return NewObjectTime(
						EnsureTime(self).Time().Add(-EnsureDuration(args.Self()).Duration()),
					)
				},

				protocols.LessThan: func(self g.Object, args g.Args) g.Object {
					return g.NewObjectBool(
						EnsureTime(self).Time().Before(EnsureTime(args.Self()).Time()),
					)
				},

				protocols.GreaterThan: func(self g.Object, args g.Args) g.Object {
					return g.NewObjectBool(
						EnsureTime(self).Time().After(EnsureTime(args.Self()).Time()),
					)
				},

				"format": func(self g.Object, args g.Args) g.Object {
					layout := g.EnsureString(args.Self()).String()
					replacements := map[string]string{
						"%Y": "2006",
						"%m": "01",
						"%d": "02",
						"%H": "15",
						"%M": "04",
						"%S": "05",
					}
					for find, replace := range replacements {
						layout = strings.ReplaceAll(layout, find, replace)
					}
					return g.NewObjectString(EnsureTime(self).Time().Format(layout))
				},
			},
		}

		TypeDuration = &g.Type{
			Name:   "time.Duration",
			Parent: g.TypeBase,
			Methods: g.Methods{
				protocols.String: func(self g.Object, _ g.Args) g.Object {
					return g.NewObjectString(EnsureDuration(self).Duration().String())
				},

				protocols.Bool: func(self g.Object, _ g.Args) g.Object {
					return g.NewObjectBool(EnsureDuration(self).Duration() != 0)
				},

				protocols.Number: func(self g.Object, _ g.Args) g.Object {
					return g.NewObjectNumber(float64(EnsureDuration(self).Duration().Seconds()))
				},

				protocols.Add: func(self g.Object, args g.Args) g.Object {
					return NewObjectDuration(
						EnsureDuration(self).Duration() + EnsureDuration(args.Self()).Duration(),
					)
				},

				protocols.Sub: func(self g.Object, args g.Args) g.Object {
					return NewObjectDuration(
						EnsureDuration(self).Duration() - EnsureDuration(args.Self()).Duration(),
					)
				},

				protocols.LessThan: func(self g.Object, args g.Args) g.Object {
					return g.NewObjectBool(
						EnsureDuration(self).Duration() < EnsureDuration(args.Self()).Duration(),
					)
				},

				protocols.GreaterThan: func(self g.Object, args g.Args) g.Object {
					return g.NewObjectBool(
						EnsureDuration(self).Duration() > EnsureDuration(args.Self()).Duration(),
					)
				},
			},
		}
	},

	Values: map[string]g.Object{
		"new": g.NewObjectInternalFunc(func(args g.Args) g.Object {
			str := g.EnsureString(args.Self())

			value, err := time.Parse("2006-01-02 15:04:05", str.String())
			errors.ErrRuntime.ExpectNil(err, "%v", err)

			return NewObjectTime(value)
		}),

		"now": g.NewObjectInternalFunc(func(_ g.Args) g.Object {
			return NewObjectTime(time.Now())
		}),

		"sleep": g.NewObjectInternalFunc(func(args g.Args) g.Object {
			duration := time.Duration(g.EnsureNumber(args.Self()).Int())

			time.Sleep(duration * time.Millisecond)

			return g.NewObjectNil()
		}),

		"since": g.NewObjectInternalFunc(func(args g.Args) g.Object {
			duration := time.Since(EnsureTime(args.Self()).Time())
			return NewObjectDuration(duration)
		}),
	},
}

package g

import (
	"strconv"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

func initstring() {
	TypeString = &Type{
		Name:   "String",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Bool: Method(func(self Object, _ Args) Object {
				return NewObjectBool(EnsureString(self).Len() > 0)
			}),

			protocols.Number: Method(func(self Object, _ Args) Object {
				value, err := strconv.ParseFloat(EnsureString(self).String(), 64)
				errors.ErrRuntime.ExpectNil(err, err.Error())

				return NewObjectNumber(value)
			}),

			protocols.Len: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(float64(EnsureString(self).Len()))
			}),

			protocols.Index: Method(func(self Object, args Args) Object {
				index := ToInt(args.Self())
				return NewObjectString(EnsureString(self).String()[index : index+1])
			}),
		},
	}
}

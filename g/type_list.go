package g

import "github.com/johnfrankmorgan/gazebo/protocols"

func initlist() {
	TypeList = &Type{
		Name:   "List",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Bool: Method(func(self Object, _ Args) Object {
				return NewObjectBool(EnsureList(self).Len() > 0)
			}),

			protocols.Len: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(float64(EnsureList(self).Len()))
			}),

			protocols.Index: Method(func(self Object, args Args) Object {
				index := ToInt(args.Self())

				return EnsureList(self).Index(index)
			}),
		},
	}
}

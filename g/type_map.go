package g

func initmap() {
	TypeMap = &Type{
		Name:   "Map",
		Parent: TypeBase,
		Methods: Methods{
			Protocols.ToBool: Method(func(self Object, args Args) Object {
				return NewObjectBool(EnsureMap(self).Len() > 0)
			}),

			Protocols.Index: Method(func(self Object, args Args) Object {
				return EnsureMap(self).Get(args.Self())
			}),

			Protocols.Len: Method(func(self Object, args Args) Object {
				return NewObjectNumber(float64(EnsureMap(self).Len()))
			}),

			Protocols.HasAttr: Method(func(self Object, args Args) Object {
				if EnsureMap(self).Has(args.Self()) {
					return NewObjectBool(true)
				}

				return ParentCall(self, Protocols.HasAttr, args)
			}),

			Protocols.GetAttr: Method(func(self Object, args Args) Object {
				if m := EnsureMap(self); m.Has(args.Self()) {
					return m.Get(args.Self())
				}

				return ParentCall(self, Protocols.GetAttr, args)
			}),

			Protocols.SetAttr: Method(func(self Object, args Args) Object {
				args.Expects(2)

				EnsureMap(self).Set(args.Self(), args[1])

				return NewObjectNil()
			}),

			Protocols.DelAttr: Method(func(self Object, args Args) Object {
				EnsureMap(self).Delete(args.Self())

				return NewObjectNil()
			}),
		},
	}
}

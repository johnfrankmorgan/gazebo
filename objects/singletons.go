package objects

type _singletons struct {
	Null  *Null
	False *Bool
	True  *Bool
}

var Singletons = _singletons{
	Null: &Null{
		Object: ObjectInit(Types.Null),
	},

	False: &Bool{
		Integer: Integer{
			Object: ObjectInit(Types.Bool),
			value:  0,
		},
	},

	True: &Bool{
		Integer: Integer{
			Object: ObjectInit(Types.Bool),
			value:  1,
		},
	},
}

func (s _singletons) Bool(value bool) *Bool {
	if value {
		return s.True
	}

	return s.False
}

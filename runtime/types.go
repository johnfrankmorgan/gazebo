package runtime

type _types struct {
	Bool          *Type
	Exception     *Type
	Float         *Type
	Func          *Type
	Int           *Type
	List          *Type
	Map           *Type
	Nil           *Type
	Object        *Type
	String        *Type
	Tuple         *Type
	Type          *Type
	Unimplemented *Type

	Exc struct {
		Unimplemented     *Type
		InvalidType       *Type
		InvalidAttribute  *Type
		InvalidIndex      *Type
		OutOfBounds       *Type
		KeyNotFound       *Type
		UndefinedVariable *Type
	}
}

var Types _types

func init() {
	Types.Object = &Type{
		Name: "Object",

		Protocols: Protocols{
			Bool:   func(Object) Bool { return True },
			Repr:   func(self Object) String { return Stringf("%s(%p)", self.Type().Name, self) },
			String: func(self Object) String { return Objects.Repr(self) },

			Binary: BinaryProtocols{
				Equal: func(self, other Object) Object { return Objects.Is(self, other) },
			},

			Attribute: AttributeProtocols{
				Get: func(self Object, name String) Object {
					if attr, ok := self.Type().Attribute(name); ok && attr.Get != nil {
						return attr.Get(self)
					}

					panic(Exc.NewInvalidAttributeGet(name, self.Type()))
				},

				Set: func(self Object, name String, value Object) {
					if attr, ok := self.Type().Attribute(name); ok && attr.Set != nil {
						attr.Set(self, value)
						return
					}

					panic(Exc.NewInvalidAttributeSet(name, self.Type()))
				},
			},
		},

		Attributes: Attributes{
			"type": Attribute{
				Get: func(self Object) Object { return self.Type() },
			},
		},
	}

	Types.Bool = &Type{
		Name:   "Bool",
		Parent: Types.Object,
		Protocols: Protocols{
			Hash: func(self Object) uint64 { return self.(Bool).Hash() },
			Bool: func(self Object) Bool { return self.(Bool).Bool() },
			Repr: func(self Object) String { return self.(Bool).Repr() },
		},
	}

	Types.Exception = &Type{
		Name:   "Exception",
		Parent: Types.Object,
		Protocols: Protocols{
			Repr:   func(self Object) String { return self.(*Exception).Repr() },
			String: func(self Object) String { return self.(*Exception).String() },
		},
	}

	Types.Float = &Type{
		Name:   "Float",
		Parent: Types.Object,
		Protocols: Protocols{
			Hash: func(self Object) uint64 { return self.(Float).Hash() },
			Bool: func(self Object) Bool { return self.(Float).Bool() },
			Repr: func(self Object) String { return self.(Float).Repr() },

			Unary: UnaryProtocols{
				Positive: func(self Object) Object { return self.(Float) },
				Negative: func(self Object) Object { return -self.(Float) },
			},

			Binary: BinaryProtocols{
				Equal:   func(self, other Object) Object { return self.(Float).Equal(other) },
				Less:    func(self, other Object) Object { return self.(Float).Less(other) },
				Greater: func(self, other Object) Object { return self.(Float).Greater(other) },

				Add:      func(self, other Object) Object { return self.(Float).Add(other) },
				Subtract: func(self, other Object) Object { return self.(Float).Subtract(other) },
				Multiply: func(self, other Object) Object { return self.(Float).Multiply(other) },
				Divide:   func(self, other Object) Object { return self.(Float).Divide(other) },
			},
		},
	}

	Types.Func = &Type{
		Name:   "Func",
		Parent: Types.Object,

		Protocols: Protocols{
			Repr: func(self Object) String { return self.(*Func).Repr() },
			Call: func(self Object, args Tuple) Object { return self.(*Func).Call(args) },
		},

		Attributes: Attributes{
			"name": Attribute{
				Get: func(self Object) Object { return self.(*Func).Name() },
			},
		},
	}

	Types.Int = &Type{
		Name:   "Int",
		Parent: Types.Object,
		Protocols: Protocols{
			Hash: func(self Object) uint64 { return self.(Int).Hash() },
			Bool: func(self Object) Bool { return self.(Int).Bool() },
			Repr: func(self Object) String { return self.(Int).Repr() },

			Unary: UnaryProtocols{
				Positive: func(self Object) Object { return self.(Int) },
				Negative: func(self Object) Object { return -self.(Int) },
			},

			Binary: BinaryProtocols{
				Equal:   func(self, other Object) Object { return self.(Int).Equal(other) },
				Less:    func(self, other Object) Object { return self.(Int).Less(other) },
				Greater: func(self, other Object) Object { return self.(Int).Greater(other) },

				Add:      func(self, other Object) Object { return self.(Int).Add(other) },
				Subtract: func(self, other Object) Object { return self.(Int).Subtract(other) },
				Multiply: func(self, other Object) Object { return self.(Int).Multiply(other) },
				Divide:   func(self, other Object) Object { return self.(Int).Divide(other) },
				Modulo:   func(self, other Object) Object { return self.(Int).Modulo(other) },

				BitwiseAnd: func(self, other Object) Object { return self.(Int).BitwiseAnd(other) },
				BitwiseOr:  func(self, other Object) Object { return self.(Int).BitwiseOr(other) },
				BitwiseXor: func(self, other Object) Object { return self.(Int).BitwiseXor(other) },

				ShiftLeft:  func(self, other Object) Object { return self.(Int).ShiftLeft(other) },
				ShiftRight: func(self, other Object) Object { return self.(Int).ShiftRight(other) },
			},
		},
	}

	Types.List = &Type{
		Name:   "List",
		Parent: Types.Object,
		Protocols: Protocols{
			Bool: func(self Object) Bool { return self.(*List).Bool() },
			Repr: func(self Object) String { return self.(*List).Repr() },

			Binary: BinaryProtocols{
				Equal:    func(self, other Object) Object { return self.(*List).Equal(other) },
				Contains: func(self, other Object) Object { return self.(*List).Contains(other) },
				Add:      func(self, other Object) Object { return self.(*List).Add(other) },
				Multiply: func(self, other Object) Object { return self.(*List).Multiply(other) },

				Right: RBinaryProtocols{
					Multiply: func(self, other Object) Object { return self.(*List).Multiply(other) },
				},
			},

			Index: IndexProtocols{
				Get: func(self, index Object) Object { return self.(*List).GetIndex(index) },
				Set: func(self, index, value Object) { self.(*List).SetIndex(index, value) },
			},
		},
		Attributes: Attributes{
			"len": Attribute{
				Get: func(self Object) Object { return self.(*List).Len() },
			},
		},
		New: func(args Tuple) Object {
			l := NewListWithLength(len(args))

			for _, arg := range args {
				l.Append(arg)
			}

			return l
		},
	}

	Types.Map = &Type{
		Name:   "Map",
		Parent: Types.Object,
		Protocols: Protocols{
			Bool: func(self Object) Bool { return self.(*Map).Bool() },

			Binary: BinaryProtocols{
				Contains: func(self, other Object) Object { return self.(*Map).Contains(other) },
			},

			Attribute: AttributeProtocols{
				Get: func(self Object, name String) (value Object) {
					if attr, ok := self.Type().Attribute(name); ok && attr.Get != nil {
						return attr.Get(self)
					}

					return Objects.Index.Get(self, name)
				},

				Set: func(self Object, name String, value Object) {
					if attr, ok := self.Type().Attribute(name); ok && attr.Set != nil {
						attr.Set(self, value)
						return
					}

					Objects.Index.Set(self, name, value)
				},
			},

			Index: IndexProtocols{
				Get: func(self, key Object) Object {
					value, ok := self.(*Map).Get(key)
					if !ok {
						panic(Exc.NewKeyNotFound(key))
					}

					return value
				},

				Set: func(self, key, value Object) { self.(*Map).Set(key, value) },
			},
		},
		Attributes: Attributes{
			"len": Attribute{
				Get: func(self Object) Object { return self.(*Map).Len() },
			},
		},
	}

	Types.Nil = &Type{
		Name:   "Nil",
		Parent: Types.Object,
		Protocols: Protocols{
			Bool: func(self Object) Bool { return self.(Nil_).Bool() },
			Repr: func(self Object) String { return self.(Nil_).Repr() },
		},
	}

	Types.String = &Type{
		Name:   "String",
		Parent: Types.Object,
		Protocols: Protocols{
			Hash:   func(self Object) uint64 { return self.(String).Hash() },
			Bool:   func(self Object) Bool { return self.(String).Bool() },
			Repr:   func(self Object) String { return self.(String).Repr() },
			String: func(self Object) String { return self.(String).String() },

			Binary: BinaryProtocols{
				Equal:    func(self, other Object) Object { return self.(String).Equal(other) },
				Less:     func(self, other Object) Object { return self.(String).Less(other) },
				Greater:  func(self, other Object) Object { return self.(String).Greater(other) },
				Add:      func(self, other Object) Object { return self.(String).Add(other) },
				Multiply: func(self, other Object) Object { return self.(String).Multiply(other) },

				Right: RBinaryProtocols{
					Multiply: func(self, other Object) Object { return self.(String).Multiply(other) },
				},
			},

			Index: IndexProtocols{
				Get: func(self, index Object) Object { return self.(String).GetIndex(index) },
			},
		},
		Attributes: Attributes{
			"len": Attribute{
				Get: func(self Object) Object { return self.(String).Len() },
			},
		},
		New: func(args Tuple) Object {
			s := String("")

			for i, arg := range args {
				if i > 0 {
					s += " "
				}

				s += Objects.String(arg)
			}

			return s
		},
	}

	Types.Tuple = &Type{
		Name:   "Tuple",
		Parent: Types.Object,
		Protocols: Protocols{
			Hash: func(self Object) uint64 { return self.(Tuple).Hash() },
			Bool: func(self Object) Bool { return self.(Tuple).Bool() },

			Binary: BinaryProtocols{
				Equal:    func(self, other Object) Object { return self.(Tuple).Equal(other) },
				Contains: func(self, other Object) Object { return self.(Tuple).Contains(other) },
				Add:      func(self, other Object) Object { return self.(Tuple).Add(other) },
				Multiply: func(self, other Object) Object { return self.(Tuple).Multiply(other) },
			},

			Index: IndexProtocols{
				Get: func(self, index Object) Object { return self.(Tuple).GetIndex(index) },
			},
		},
		Attributes: Attributes{
			"len": Attribute{
				Get: func(self Object) Object { return self.(Tuple).Len() },
			},
		},
		New: func(args Tuple) Object {
			return append(Tuple(nil), args...)
		},
	}

	Types.Type = &Type{
		Name:   "Type",
		Parent: Types.Object,
		Protocols: Protocols{
			Repr: func(self Object) String { return self.(*Type).Repr() },
		},
		Attributes: Attributes{
			"name": Attribute{
				Get: func(self Object) Object { return self.(*Type).Name },
				Set: func(self, value Object) { self.(*Type).Name = value.(String) },
			},

			"parent": Attribute{
				Get: func(self Object) Object {
					if parent := self.(*Type).Parent; parent != nil {
						return parent
					}

					return Nil
				},
			},
		},
	}

	Types.Unimplemented = &Type{
		Name:   "Unimplemented",
		Parent: Types.Object,
	}

	Types.Exc.Unimplemented = &Type{
		Name:   "Unimplemented",
		Parent: Types.Exception,
	}

	Types.Exc.InvalidType = &Type{
		Name:   "InvalidType",
		Parent: Types.Exception,
	}

	Types.Exc.InvalidAttribute = &Type{
		Name:   "InvalidAttribute",
		Parent: Types.Exception,
	}

	Types.Exc.InvalidIndex = &Type{
		Name:   "InvalidIndex",
		Parent: Types.Exception,
	}

	Types.Exc.OutOfBounds = &Type{
		Name:   "OutOfBounds",
		Parent: Types.Exc.InvalidIndex,
	}

	Types.Exc.KeyNotFound = &Type{
		Name:   "KeyNotFound",
		Parent: Types.Exc.InvalidIndex,
	}

	Types.Exc.UndefinedVariable = &Type{
		Name:   "UndefinedVariable",
		Parent: Types.Exception,
	}
}

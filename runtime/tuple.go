package runtime

type Tuple []Object

var TupleType = &Type{
	Name:   "Tuple",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(Tuple).Hash() },
		Bool:   func(self Object) Bool { return self.(Tuple).Bool() },
		String: func(self Object) String { return self.(Tuple).String() },
	},
	Ops: TypeOps{
		Equal:    func(self, other Object) Bool { return self.(Tuple).Equal(other) },
		Contains: func(self, other Object) Bool { return self.(Tuple).Contains(other) },
		Add:      func(self, other Object) Object { return self.(Tuple).Add(other) },
		Multiply: func(self, other Object) Object { return self.(Tuple).Multiply(other) },
		GetIndex: func(self, index Object) Object { return self.(Tuple).GetIndex(index) },
	},
}

func (t Tuple) Type() *Type {
	return TupleType
}

func (t Tuple) Hash() uint64 {
	hash := uint64(0)

	for _, item := range t {
		hash ^= Hash(item)
	}

	return hash
}

func (t Tuple) Bool() Bool {
	return t.Len() != 0
}

func (t Tuple) String() String {
	panic("todo")
}

func (t Tuple) Len() Int {
	return Int(len(t))
}

func (t Tuple) Equal(other Object) Bool {
	if other, ok := other.(Tuple); ok {
		if t.Len() != other.Len() {
			return False
		}

		for i, item := range t {
			if !Equal(item, other[i]) {
				return False
			}
		}

		return True
	}

	panic(ErrUnimplemented)
}

func (t Tuple) Contains(other Object) Bool {
	for _, item := range t {
		if Equal(item, other) {
			return True
		}
	}

	return False
}

func (t Tuple) Add(other Object) Object {
	if other, ok := other.(Tuple); ok {
		result := make(Tuple, 0, t.Len()+other.Len())
		result = append(result, t...)
		result = append(result, other...)
		return result
	}

	panic(ErrUnimplemented)
}

func (t Tuple) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		result := make(Tuple, 0, len(t)*int(other))

		for i := 0; i < int(other); i++ {
			result = append(result, t...)
		}

		return result
	}

	panic(ErrUnimplemented)
}

func (t Tuple) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		return t[index]
	}

	panic(ErrUnimplemented)
}

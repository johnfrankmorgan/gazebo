package runtime

type Tuple []Object

func (t Tuple) Type() *Type {
	return Types.Tuple
}

func (t Tuple) Hash() uint64 {
	hash := uint64(0)

	for _, item := range t {
		hash ^= Objects.Hash(item)
	}

	return hash
}

func (t Tuple) Bool() Bool {
	return t.Len() != 0
}

func (t Tuple) Len() Int {
	return Int(len(t))
}

func (t Tuple) Equal(other Object) Object {
	if other, ok := other.(Tuple); ok {
		if t.Len() != other.Len() {
			return False
		}

		for i, item := range t {
			if !Objects.Binary.Equal(item, other[i]) {
				return False
			}
		}

		return True
	}

	return Unimplemented
}

func (t Tuple) Contains(other Object) Bool {
	for _, item := range t {
		if Objects.Binary.Equal(item, other) {
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

	return Unimplemented
}

func (t Tuple) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		result := make(Tuple, 0, len(t)*int(other))

		for i := 0; i < int(other); i++ {
			result = append(result, t...)
		}

		return result
	}

	return Unimplemented
}

func (t Tuple) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		return t[index]
	}

	panic(Exc.NewInvalidType(index.Type(), Types.Int))
}

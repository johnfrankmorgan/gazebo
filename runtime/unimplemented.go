package runtime

type Unimplemented_ int

const Unimplemented Unimplemented_ = 0

func (Unimplemented_) Type() *Type {
	return Types.Unimplemented
}

func (_objects) IsUnimplemented(object Object) Bool {
	return Objects.Is(object, Unimplemented)
}

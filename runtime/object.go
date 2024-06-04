package runtime

type Object interface {
	Type() *Type
}

type (
	_objects struct {
		Attribute _attribute
		Index     _index
		Unary     _unary
		Binary    _binary
	}

	_attribute struct{}
	_index     struct{}
	_unary     struct{}
	_rbinary   struct{}

	_binary struct {
		Right _rbinary
	}
)

var Objects _objects

func (_objects) Is(self, other Object) Bool {
	return self == other
}

func (_objects) IsInstance(self Object, typ *Type) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if Objects.Is(t, typ) {
			return True
		}
	}

	return False
}

func (_objects) Hash(self Object) uint64 {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Hash != nil {
			return t.Protocols.Hash(self)
		}
	}

	panic(Exc.NewUnimplemented("hash", self.Type()))
}

func (_objects) Bool(self Object) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Bool != nil {
			return t.Protocols.Bool(self)
		}
	}

	panic(Exc.NewUnimplemented("bool", self.Type()))
}

func (_objects) Repr(self Object) String {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Repr != nil {
			return t.Protocols.Repr(self)
		}
	}

	panic(Exc.NewUnimplemented("repr", self.Type()))
}

func (_objects) String(self Object) String {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.String != nil {
			return t.Protocols.String(self)
		}
	}

	panic(Exc.NewUnimplemented("string", self.Type()))
}

func (_attribute) Get(self Object, name String) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Attribute.Get != nil {
			return t.Protocols.Attribute.Get(self, name)
		}
	}

	panic(Exc.NewUnimplemented("get attribute", self.Type()))
}

func (_attribute) Set(self Object, name String, value Object) {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Attribute.Set != nil {
			t.Protocols.Attribute.Set(self, name, value)
			return
		}
	}

	panic(Exc.NewUnimplemented("set attribute", self.Type()))
}

func (_index) Get(self, index Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Index.Get != nil {
			return t.Protocols.Index.Get(self, index)
		}
	}

	panic(Exc.NewUnimplemented("get index", self.Type()))
}

func (_index) Set(self, index, value Object) {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Index.Set != nil {
			t.Protocols.Index.Set(self, index, value)
			return
		}
	}

	panic(Exc.NewUnimplemented("set index", self.Type()))
}

func (_unary) Positive(self Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Unary.Positive != nil {
			return t.Protocols.Unary.Positive(self)
		}
	}

	panic(Exc.NewUnimplementedUnary(UnaryProtocolPositive, self.Type()))
}

func (_unary) Negative(self Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Unary.Negative != nil {
			return t.Protocols.Unary.Negative(self)
		}
	}

	panic(Exc.NewUnimplementedUnary(UnaryProtocolNegative, self.Type()))
}

func (_binary) Equal(self, other Object) (result Bool) {
	if Objects.Is(self, other) {
		return True
	}

	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Equal(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Equal != nil {
			return t.Protocols.Binary.Equal(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolEqual, self.Type(), other.Type()))
}

func (_binary) NotEqual(self, other Object) Bool {
	return !Objects.Binary.Equal(self, other)
}

func (_binary) Less(self, other Object) (result Bool) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Less(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Less != nil {
			return t.Protocols.Binary.Less(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolLess, self.Type(), other.Type()))
}

func (_binary) LessOrEqual(self, other Object) Bool {
	return Objects.Binary.Less(self, other) || Objects.Binary.Equal(self, other)
}

func (_binary) Greater(self, other Object) (result Bool) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Greater(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Greater != nil {
			return t.Protocols.Binary.Greater(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolGreater, self.Type(), other.Type()))
}

func (_binary) GreaterOrEqual(self, other Object) Bool {
	return Objects.Binary.Greater(self, other) || Objects.Binary.Equal(self, other)
}

func (_binary) Contains(self, other Object) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Contains != nil {
			return t.Protocols.Binary.Contains(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolContains, self.Type(), other.Type()))
}

func (_binary) Add(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Add(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Add != nil {
			return t.Protocols.Binary.Add(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolAdd, self.Type(), other.Type()))
}

func (_binary) Subtract(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Subtract(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Subtract != nil {
			return t.Protocols.Binary.Subtract(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolSubtract, self.Type(), other.Type()))
}

func (_binary) Multiply(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Multiply(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Multiply != nil {
			return t.Protocols.Binary.Multiply(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolMultiply, self.Type(), other.Type()))
}

func (_binary) Divide(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Divide(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Divide != nil {
			return t.Protocols.Binary.Divide(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolDivide, self.Type(), other.Type()))
}

func (_binary) Modulo(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.Modulo(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Modulo != nil {
			return t.Protocols.Binary.Modulo(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolModulo, self.Type(), other.Type()))
}

func (_binary) BitwiseAnd(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.BitwiseAnd(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.BitwiseAnd != nil {
			return t.Protocols.Binary.BitwiseAnd(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseAnd, self.Type(), other.Type()))
}

func (_binary) BitwiseOr(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.BitwiseOr(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.BitwiseOr != nil {
			return t.Protocols.Binary.BitwiseOr(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseOr, self.Type(), other.Type()))
}

func (_binary) BitwiseXor(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.BitwiseXor(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.BitwiseXor != nil {
			return t.Protocols.Binary.BitwiseXor(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseXor, self.Type(), other.Type()))
}

func (_binary) ShiftLeft(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.ShiftLeft(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.ShiftLeft != nil {
			return t.Protocols.Binary.ShiftLeft(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftLeft, self.Type(), other.Type()))
}

func (_binary) ShiftRight(self, other Object) (result Object) {
	defer func() {
		if r := recover(); r != nil {
			if exc, ok := r.(*Exception); Bool(ok) && Objects.IsInstance(exc, Types.Exc.Unimplemented) {
				result = Objects.Binary.Right.ShiftRight(other, self)
			} else {
				panic(r)
			}
		}
	}()

	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.ShiftRight != nil {
			return t.Protocols.Binary.ShiftRight(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftRight, self.Type(), other.Type()))
}

func (_rbinary) Equal(self, other Object) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Equal != nil {
			return t.Protocols.Binary.Right.Equal(self, other)
		}
	}

	return False
}

func (_rbinary) Less(self, other Object) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Less != nil {
			return t.Protocols.Binary.Right.Less(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolLess, other.Type(), self.Type()))
}

func (_rbinary) Greater(self, other Object) Bool {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Greater != nil {
			return t.Protocols.Binary.Right.Greater(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolGreater, other.Type(), self.Type()))
}

func (_rbinary) Add(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Add != nil {
			return t.Protocols.Binary.Right.Add(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolAdd, other.Type(), self.Type()))
}

func (_rbinary) Subtract(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Subtract != nil {
			return t.Protocols.Binary.Right.Subtract(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolSubtract, other.Type(), self.Type()))
}

func (_rbinary) Multiply(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Multiply != nil {
			return t.Protocols.Binary.Right.Multiply(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolMultiply, other.Type(), self.Type()))
}

func (_rbinary) Divide(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Divide != nil {
			return t.Protocols.Binary.Right.Divide(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolDivide, other.Type(), self.Type()))
}

func (_rbinary) Modulo(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.Modulo != nil {
			return t.Protocols.Binary.Right.Modulo(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolModulo, other.Type(), self.Type()))
}

func (_rbinary) BitwiseAnd(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.BitwiseAnd != nil {
			return t.Protocols.Binary.Right.BitwiseAnd(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseAnd, other.Type(), self.Type()))
}

func (_rbinary) BitwiseOr(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.BitwiseOr != nil {
			return t.Protocols.Binary.Right.BitwiseOr(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseOr, other.Type(), self.Type()))
}

func (_rbinary) BitwiseXor(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.BitwiseXor != nil {
			return t.Protocols.Binary.Right.BitwiseXor(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseXor, other.Type(), self.Type()))
}

func (_rbinary) ShiftLeft(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.ShiftLeft != nil {
			return t.Protocols.Binary.Right.ShiftLeft(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftLeft, other.Type(), self.Type()))
}

func (_rbinary) ShiftRight(self, other Object) Object {
	for t := self.Type(); t != nil; t = t.Parent {
		if t.Protocols.Binary.Right.ShiftRight != nil {
			return t.Protocols.Binary.Right.ShiftRight(self, other)
		}
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftRight, other.Type(), self.Type()))
}

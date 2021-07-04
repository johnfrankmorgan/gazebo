package vm

type TypeFunction struct {
	TypeBase
}

func (m *TypeFunction) Name() String {
	return NewString("Function")
}

func (m *TypeFunction) Call(self Object, args Args) Object {
	return self.(*Function).Call(args)
}

package runtime

import "golang.org/x/exp/constraints"

type List struct {
	_items []Object
}

var ListType = &Type{
	Name:   "List",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Bool:   func(self Object) Bool { return self.(*List).Bool() },
		String: func(self Object) String { return self.(*List).String() },
	},
	Ops: TypeOps{
		Equal:    func(self, other Object) Bool { return self.(*List).Equal(other) },
		Contains: func(self, other Object) Bool { return self.(*List).Contains(other) },
		Add:      func(self, other Object) Object { return self.(*List).Add(other) },
		Multiply: func(self, other Object) Object { return self.(*List).Multiply(other) },
		GetIndex: func(self, index Object) Object { return self.(*List).GetIndex(index) },
		SetIndex: func(self, index, value Object) { self.(*List).SetIndex(index, value) },
	},
	Attributes: TypeAttributes{
		"len": Attribute{
			Get: func(self Object) Object { return self.(*List).Len() },
		},
	},
}

func NewList(items ...Object) *List {
	return &List{
		_items: items,
	}
}

func NewListWithLength[T constraints.Signed](length T) *List {
	return &List{
		_items: make([]Object, length),
	}
}

func NewListWithCapacity[T constraints.Signed](capacity T) *List {
	return &List{
		_items: make([]Object, 0, capacity),
	}
}

func (l *List) Type() *Type {
	return ListType
}

func (l *List) Bool() Bool {
	return l.Len() != 0
}

func (l *List) String() String {
	panic("todo")
}

func (l *List) Len() Int {
	return Int(len(l._items))
}

func (l *List) Append(others ...Object) {
	l._items = append(l._items, others...)
}

func (l *List) Get(index Int) Object {
	return l._items[index]
}

func (l *List) Set(index Int, value Object) {
	l._items[index] = value
}

func (l *List) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		return l.Get(index)
	}

	panic(ErrUnimplemented)
}

func (l *List) SetIndex(index, value Object) {
	if index, ok := index.(Int); ok {
		l.Set(index, value)
		return
	}
}

func (l *List) Equal(other Object) Bool {
	if other, ok := other.(*List); ok {
		if l.Len() != other.Len() {
			return False
		}

		for i := range l.Len() {
			if !Equal(l.Get(i), other.Get(i)) {
				return False
			}
		}

		return False
	}

	panic(ErrUnimplemented)
}

func (l *List) Contains(other Object) Bool {
	for _, item := range l._items {
		if Equal(item, other) {
			return True
		}
	}

	return False
}

func (l *List) Add(other Object) Object {
	if other, ok := other.(*List); ok {
		result := NewListWithCapacity(l.Len() + other.Len())
		result.Append(l._items...)
		result.Append(other._items...)
		return result
	}

	panic(ErrUnimplemented)
}

func (l *List) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		result := NewListWithCapacity(l.Len() * other)

		for i := Int(0); i < other; i++ {
			result.Append(l._items...)
		}

		return result
	}

	panic(ErrUnimplemented)
}

package runtime

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type List struct {
	items []Object
}

func NewList(items ...Object) *List {
	return &List{
		items: items,
	}
}

func NewListWithLength[T constraints.Signed](length T) *List {
	return &List{
		items: make([]Object, length),
	}
}

func NewListWithCapacity[T constraints.Signed](capacity T) *List {
	return &List{
		items: make([]Object, 0, capacity),
	}
}

func (l *List) Type() *Type {
	return Types.List
}

func (l *List) Bool() Bool {
	return l.Len() != 0
}

func (l *List) Repr() String {
	items := make([]string, l.Len())

	for i, item := range l.items {
		items[i] = string(Objects.Repr(item))
	}

	return String("[" + strings.Join(items, ", ") + "]")
}

func (l *List) Len() Int {
	return Int(len(l.items))
}

func (l *List) Append(others ...Object) {
	l.items = append(l.items, others...)
}

func (l *List) Get(index Int) Object {
	if index < 0 || index >= l.Len() {
		panic(Exc.NewOutOfBounds(index, l.Len()))
	}

	return l.items[index]
}

func (l *List) Set(index Int, value Object) {
	l.items[index] = value
}

func (l *List) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		return l.Get(index)
	}

	panic(Exc.NewInvalidType(index.Type(), Types.Int))
}

func (l *List) SetIndex(index, value Object) {
	if index, ok := index.(Int); ok {
		l.Set(index, value)
		return
	}

	panic(Exc.NewInvalidType(index.Type(), Types.Int))
}

func (l *List) Equal(other Object) Object {
	if other, ok := other.(*List); ok {
		if l.Len() != other.Len() {
			return False
		}

		for i := range l.Len() {
			if !Objects.Binary.Equal(l.Get(i), other.Get(i)) {
				return False
			}
		}

		return False
	}

	return Unimplemented
}

func (l *List) Contains(other Object) Bool {
	for _, item := range l.items {
		if Objects.Binary.Equal(item, other) {
			return True
		}
	}

	return False
}

func (l *List) Add(other Object) Object {
	if other, ok := other.(*List); ok {
		result := NewListWithCapacity(l.Len() + other.Len())
		result.Append(l.items...)
		result.Append(other.items...)
		return result
	}

	return Unimplemented
}

func (l *List) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		result := NewListWithCapacity(l.Len() * other)

		for i := Int(0); i < other; i++ {
			result.Append(l.items...)
		}

		return result
	}

	return Unimplemented
}

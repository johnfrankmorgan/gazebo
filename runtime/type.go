package runtime

type Type struct {
	Name       String
	Parent     *Type
	Protocols  Protocols
	Attributes Attributes
	New        func(Tuple) Object
}

func (t *Type) Type() *Type {
	return Types.Type
}

func (t *Type) Repr() String {
	return Stringf("%s(%s@%p)", t.Type().Name, t.Name, t)
}

func (t *Type) Attribute(name String) (Attribute, bool) {
	if name == "new" {
		return Attribute{
			Get: func(self Object) Object {
				return &Func{Stringf("%s.new", self.(*Type).Name), self.(*Type).New}
			},
		}, true
	}

	for t := t; t != nil; t = t.Parent {
		if attr, ok := t.Attributes[name]; ok {
			return attr, true
		}
	}

	return Attribute{}, false
}

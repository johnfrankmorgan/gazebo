package runtime

type Type struct {
	Name       String
	Parent     *Type
	Protocols  Protocols
	Attributes Attributes
}

func (t *Type) Type() *Type {
	return Types.Type
}

func (t *Type) Attribute(name String) (Attribute, bool) {
	for t := t; t != nil; t = t.Parent {
		if attr, ok := t.Attributes[name]; ok {
			return attr, true
		}
	}

	return Attribute{}, false
}

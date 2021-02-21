package g

import "github.com/johnfrankmorgan/gazebo/protocols"

// Type represents the type of gazebo values
type Type struct {
	Name    string
	Parent  *Type
	Methods Methods
}

// Method is the type of methods in gazebo
type Method func(Object, Args) Object

// Methods is a map of names to Funcs
type Methods map[string]Method

// Resolve resolves a method on a Type
func (m *Type) Resolve(name string) Method {
	if method, ok := m.Methods[name]; ok {
		return method
	}

	if m.Parent != nil {
		return m.Parent.Resolve(name)
	}

	return nil
}

// Implements checks if a method is implemented by a Type
func (m *Type) Implements(name string) bool {
	return m.Resolve(name) != nil
}

// Builtin types
var (
	TypeBase         *Type
	TypeNil          *Type
	TypeBool         *Type
	TypeNumber       *Type
	TypeString       *Type
	TypeList         *Type
	TypeInternalFunc *Type
	TypeFunc         *Type
	TypeInternal     *Type
)

func init() {
	for _, init := range []func(){initbase, initnil, initbool, initnumber, initstring, initlist} {
		init()
	}

	TypeInternalFunc = &Type{
		Name:   "InternalFunc",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Invoke: Method(func(self Object, args Args) Object {
				return EnsureInternalFunc(self).Func()(args)
			}),
		},
	}

	TypeFunc = &Type{
		Name:    "Func",
		Parent:  TypeBase,
		Methods: Methods{},
	}

	TypeInternal = &Type{
		Name:    "Internal",
		Parent:  TypeBase,
		Methods: Methods{},
	}
}

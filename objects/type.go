package objects

type Type interface {
	Name() string
	Parent() Type
	Is(other Type) bool
	Methods() Methods
	Object
}

var (
	TypeType   *type_type
	TypeNil    *type_nil
	TypeBool   *type_bool
	TypeString *type_string
	TypeNumber *type_number
	TypeList   *type_list
	TypeMap    *type_map
)

func init() {
	TypeType = &type_type{NewBasicType("Type")}
	TypeNil = &type_nil{NewBasicType("Nil")}
	TypeBool = &type_bool{NewBasicType("Bool")}
	TypeString = &type_string{NewBasicType("String")}
	TypeNumber = &type_number{NewBasicType("Number")}
	TypeList = &type_list{NewBasicType("List")}
	TypeMap = &type_map{NewBasicType("Map")}
}

var _ Type = &BasicType{}

type BasicType struct {
	name string
}

func NewBasicType(name string) *BasicType {
	return &BasicType{name: name}
}

/* Type methods */

func (m *BasicType) Name() string {
	return m.name
}

func (m *BasicType) Parent() Type {
	panic("BasicType.Parent()")
}

func (m *BasicType) Is(other Type) bool {
	return m.Name() == other.Name()
}

func (m *BasicType) ToString() *String {
	return NewStringf("<objects.BasicType{%s}>", m.Name())
}

/* Object methods */

func (m *BasicType) GoVal() interface{} {
	panic("BasicType.GoVal()")
}

func (m *BasicType) Type() Type {
	panic("BasicType.Type()")
}

func (m *BasicType) Hash() interface{} {
	panic("BasicType.Hash()")
}

func (m *BasicType) Methods() Methods {
	panic("BasicType.Methods()")
}

package protocols

func All() []string {
	return []string{
		Not,
		Bool,
		String,
		Number,
		Invoke,
		Inspect,
		Inverse,
		Add,
		Sub,
		Mul,
		Div,
		Equal,
		NotEqual,
		GreaterThan,
		GreaterThanEqual,
		LessThan,
		LessThanEqual,
		Index,
		Len,
		HasAttr,
		GetAttr,
		SetAttr,
		DelAttr,
	}
}

const (
	Not              = "not!"
	Bool             = "truthy?"
	String           = "str"
	Number           = "num"
	Invoke           = "invoke"
	Inspect          = "inspect"
	Inverse          = "inverse"
	Add              = "add"
	Sub              = "sub"
	Mul              = "mul"
	Div              = "div"
	Equal            = "eq?"
	NotEqual         = "neq?"
	GreaterThan      = "gt?"
	GreaterThanEqual = "gte?"
	LessThan         = "lt?"
	LessThanEqual    = "lte?"
	Index            = "index"
	Len              = "len"
	HasAttr          = "has"
	GetAttr          = "get"
	SetAttr          = "set"
	DelAttr          = "del"
)

var BinaryOperators = map[string]string{
	"+":  Add,
	"-":  Sub,
	"*":  Mul,
	"/":  Div,
	"==": Equal,
	"!=": NotEqual,
	">":  GreaterThan,
	">=": GreaterThanEqual,
	"<":  LessThan,
	"<=": LessThanEqual,
}

var UnaryOperators = map[string]string{
	"!": Not,
	"-": Inverse,
}

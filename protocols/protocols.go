package protocols

func All() []string {
	return []string{
		String,
		Number,
		Bool,
		Not,
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
	String           = "str"
	Number           = "num"
	Bool             = "bool"
	Not              = "not"
	Invoke           = "invoke"
	Inspect          = "inspect"
	Inverse          = "inverse"
	Add              = "add"
	Sub              = "sub"
	Mul              = "mul"
	Div              = "div"
	Equal            = "eq"
	NotEqual         = "neq"
	GreaterThan      = "gt"
	GreaterThanEqual = "gte"
	LessThan         = "lt"
	LessThanEqual    = "lte"
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

package protocols

func All() []string {
	return []string{
		String,
		Number,
		Bool,
		Not,
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
		Invoke,
		Inverse,
	}
}

const (
	String           = "str"
	Number           = "num"
	Bool             = "bool"
	Not              = "not"
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
	Invoke           = "invoke"
	Inverse          = "inverse"
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

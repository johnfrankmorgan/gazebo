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
		HasAttr,
		GetAttr,
		SetAttr,
		DelAttr,
		Len,
		Inverse,
	}
}

const (
	String           = "str"
	Number           = "num"
	Bool             = "bool"
	Not              = "not"
	Len              = "len"
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
	HasAttr          = "hasattr"
	GetAttr          = "getattr"
	SetAttr          = "setattr"
	DelAttr          = "delattr"
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

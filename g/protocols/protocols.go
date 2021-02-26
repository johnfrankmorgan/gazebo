package protocols

func All() []string {
	return []string{
		Represention,
		String,
		Number,
		Bool,
		Not,
		And,
		Or,
		Contains,
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
		Invoke,
	}
}

const (
	Represention     = "repr"
	String           = "str"
	Number           = "num"
	Bool             = "bool"
	Not              = "not"
	Len              = "len"
	Inverse          = "inverse"
	And              = "and"
	Or               = "or"
	Contains         = "contains"
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
	Invoke           = "invoke"
)

var BinaryOperators = map[string]string{
	"and": And,
	"or":  Or,
	"in":  Contains,
	"+":   Add,
	"-":   Sub,
	"*":   Mul,
	"/":   Div,
	"==":  Equal,
	"!=":  NotEqual,
	">":   GreaterThan,
	">=":  GreaterThanEqual,
	"<":   LessThan,
	"<=":  LessThanEqual,
}

var UnaryOperators = map[string]string{
	"!": Not,
	"-": Inverse,
}

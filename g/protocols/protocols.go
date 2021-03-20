package protocols

const (
	Bool   = "bool"
	Not    = "not"
	And    = "and"
	Or     = "or"
	String = "str"
	Number = "num"
	Invoke = "invoke"
)

var BinaryOperators = map[string]string{
	"and": And,
	"or":  Or,
}

var UnaryOperators = map[string]string{}

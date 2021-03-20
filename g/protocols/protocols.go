package protocols

const (
	Bool     = "bool"
	Not      = "not"
	And      = "and"
	Or       = "or"
	String   = "str"
	Number   = "num"
	Invoke   = "invoke"
	Len      = "len"
	Add      = "add"
	Subtract = "sub"
	Multiply = "mul"
	Divide   = "div"
)

func All() []string {
	return []string{
		Bool,
		Not,
		And,
		Or,
		String,
		Number,
		Invoke,
		Len,
		Add,
		Subtract,
		Multiply,
		Divide,
	}
}

var BinaryOperators = map[string]string{
	And: And,
	Or:  Or,
	"+": Add,
	"-": Subtract,
	"*": Multiply,
	"/": Divide,
}

var UnaryOperators = map[string]string{}

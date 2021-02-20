package g

// Protocols defines the names of magic methods that are called on an Object
var Protocols = struct {
	ToBool      string
	ToString    string
	ToNumber    string
	Invoke      string
	Inspect     string
	Equal       string
	Add         string
	Sub         string
	Mul         string
	Div         string
	GreaterThan string
	LessThan    string
	Index       string
	Len         string
	HasAttr     string
	GetAttr     string
	SetAttr     string
	DelAttr     string
	Hash        string
}{
	ToBool:      "?",
	ToString:    "str",
	ToNumber:    "num",
	Invoke:      "invoke",
	Inspect:     "inspect",
	Equal:       "=",
	Add:         "+",
	Sub:         "-",
	Mul:         "*",
	Div:         "/",
	GreaterThan: ">",
	LessThan:    "<",
	Index:       "index",
	Len:         "len",
	HasAttr:     "has",
	GetAttr:     "get",
	SetAttr:     "set",
	DelAttr:     "del",
	Hash:        "#",
}

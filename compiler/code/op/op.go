package op

// Instruction is a struct containing an Opcode
// and an optional argument
type Instruction struct {
	Opcode Opcode
	Arg    interface{}
}

// Opcode is an opcode recognised by the gazebo VM
type Opcode int

// Enumeration of available opcodes
const (
	Invalid Opcode = iota
	Placeholder
	PushValue
	LoadConst
	GetName
	SetName
	DelName
	CallFunc
	RelJump
	RelJumpIfTrue
	RelJumpIfFalse
	MakeFunc
	MakeList
	MakeMap
	GetAttr
	SetAttr
	DelAttr
	Return
	LoadModule
	NoOp
)

// Placeholder values
const (
	PlaceholderBreak    = -1
	PlaceholderContinue = -2
)

// Ins creates an Instruction for an Opcode
func (op Opcode) Ins(arg interface{}) Instruction {
	return Instruction{Opcode: op, Arg: arg}
}

func (op Opcode) String() string {
	return op.Name()
}

// Name returns an Opcode's name
func (op Opcode) Name() string {
	names := map[Opcode]string{
		Invalid:        "op.Invalid",
		Placeholder:    "op.Placeholder",
		PushValue:      "op.PushValue",
		LoadConst:      "op.LoadConst",
		GetName:        "op.GetName",
		SetName:        "op.SetName",
		DelName:        "op.DelName",
		CallFunc:       "op.CallFunc",
		RelJump:        "op.RelJump",
		RelJumpIfTrue:  "op.RelJumpIfTrue",
		RelJumpIfFalse: "op.RelJumpIfFalse",
		MakeFunc:       "op.MakeFunc",
		MakeList:       "op.MakeList",
		MakeMap:        "op.MakeMap",
		GetAttr:        "op.GetAttr",
		SetAttr:        "op.SetAttr",
		DelAttr:        "op.DelAttr",
		Return:         "op.Return",
		LoadModule:     "op.LoadModule",
		NoOp:           "op.NoOp",
	}

	if name, ok := names[op]; ok {
		return name
	}

	return "op.Unknown"
}

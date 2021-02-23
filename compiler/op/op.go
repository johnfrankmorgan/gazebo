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
	LoadConst
	StoreName
	LoadName
	RemoveName
	CallFunc
	RelJump
	RelJumpIfTrue
	RelJumpIfFalse
	PushValue
	MakeFunc
	LoadModule
	MakeList
	IndexGet
	AttributeGet
	AttributeSet
	NoOp
	Return
)

// Ins creates an Instruction for an Opcode
func (op Opcode) Ins(arg interface{}) Instruction {
	return Instruction{Opcode: op, Arg: arg}
}

// Name returns an Opcode's name
func (op Opcode) Name() string {
	names := map[Opcode]string{
		Invalid:        "op.Invalid",
		LoadConst:      "op.LoadConst",
		StoreName:      "op.StoreName",
		LoadName:       "op.LoadName",
		RemoveName:     "op.RemoveName",
		CallFunc:       "op.CallFunc",
		RelJump:        "op.RelJump",
		RelJumpIfTrue:  "op.RelJumpIfTrue",
		RelJumpIfFalse: "op.RelJumpIfFalse",
		PushValue:      "op.PushValue",
		MakeFunc:       "op.MakeFunc",
		LoadModule:     "op.LoadModule",
		MakeList:       "op.MakeList",
		IndexGet:       "op.IndexGet",
		AttributeGet:   "op.AttributeGet",
		AttributeSet:   "op.AttributeSet",
		NoOp:           "op.NoOp",
		Return:         "op.Return",
	}

	if name, ok := names[op]; ok {
		return name
	}

	return "op.Unknown"
}

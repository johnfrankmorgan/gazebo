package op

//go:generate stringer -type=Op
type Op int

const (
	_ Op = iota
	LoadConst
	StoreName
	LoadName
	Jump
	RelJumpIfTrue
	RelJumpIfFalse
	RelJump
	BinEq
	BinNEq
	BinAdd
	BinSub
	BinMul
	BinDiv
	BinLess
	BinLessEq
	BinGreater
	BinGreaterEq
	MakeFunction
	Return
)

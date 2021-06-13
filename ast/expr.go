package ast

var (
	_ Expr = &EGroup{}
	_ Expr = &EBinary{}
	_ Expr = &EUnary{}
	_ Expr = &ELiteral{}
	_ Expr = &EFuncDef{}
)

type EGroup struct {
	Expr Expr
}

func (m *EGroup) Accept(v Visitor) {
	v.VisitEGroup(m)
}

type EBinary struct {
	LHS Expr
	Op  BinOp
	RHS Expr
}

func (m *EBinary) Accept(v Visitor) {
	v.VisitEBinary(m)
}

type EUnary struct {
	Op   UnaryOp
	Expr Expr
}

func (m *EUnary) Accept(v Visitor) {
	v.VisitEUnary(m)
}

//go:generate stringer -type=LitType
type LitType int

const (
	_ LitType = iota
	LitTypeIdent
	LitTypeNumber
	LitTypeString
)

type ELiteral struct {
	Type   LitType
	Lexeme string
}

func (m *ELiteral) Accept(v Visitor) {
	v.VisitELiteral(m)
}

type EFuncDef struct {
	Args []string
	Body Stmt
}

func (m *EFuncDef) Accept(v Visitor) {
	v.VisitEFuncDef(m)
}

func (m *EFuncDef) AddArg(name string) {
	m.Args = append(m.Args, name)
}

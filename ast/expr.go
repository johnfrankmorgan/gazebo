package ast

var (
	_ Expr = &EAssign{}
	_ Expr = &EGroup{}
	_ Expr = &EBinary{}
	_ Expr = &EUnary{}
	_ Expr = &ELiteral{}
	_ Expr = &EFuncDef{}
	_ Expr = &ECall{}
	_ Expr = &EAttrGet{}
	_ Expr = &EAttrSet{}
)

type EAssign struct {
	Ident string
	Expr  Expr
}

func (m *EAssign) Accept(v Visitor) {
	v.VisitEAssign(m)
}

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

type ECall struct {
	Expr Expr
	Args []Expr
}

func (m *ECall) Accept(v Visitor) {
	v.VisitECall(m)
}

type EAttrGet struct {
	Expr Expr
	Attr string
}

func (m *EAttrGet) Accept(v Visitor) {
	v.VisitEAttrGet(m)
}

type EAttrSet struct {
	Expr  Expr
	Attr  string
	Value Expr
}

func (m *EAttrSet) Accept(v Visitor) {
	v.VisitEAttrSet(m)
}

package ast

var (
	_ Stmt = &SBlock{}
	_ Stmt = &SExpr{}
	_ Stmt = &SIf{}
	_ Stmt = &SReturn{}
)

type SBlock struct {
	Stmts []Stmt
}

func (m *SBlock) Append(stmt Stmt) {
	m.Stmts = append(m.Stmts, stmt)
}

func (m *SBlock) Accept(v Visitor) {
	v.VisitSBlock(m)
}

type SExpr struct {
	Expr Expr
}

func (m *SExpr) Accept(v Visitor) {
	v.VisitSExpr(m)
}

type SIf struct {
	Condition  Expr
	TrueBlock  Stmt
	FalseBlock Stmt
}

func (m *SIf) Accept(v Visitor) {
	v.VisitSIf(m)
}

type SWhile struct {
	Condition Expr
	Body      Stmt
}

func (m *SWhile) Accept(v Visitor) {
	v.VisitSWhile(m)
}

type SReturn struct {
	Expr Expr
}

func (m *SReturn) Accept(v Visitor) {
	v.VisitSReturn(m)
}

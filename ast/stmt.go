package ast

var (
	_ Stmt = &SBlock{}
	_ Stmt = &SAssign{}
	_ Stmt = &SExpr{}
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

type SAssign struct {
	Ident string
	Expr  Expr
}

func (m *SAssign) Accept(v Visitor) {
	v.VisitSAssign(m)
}

type SExpr struct {
	Expr Expr
}

func (m *SExpr) Accept(v Visitor) {
	v.VisitSExpr(m)
}

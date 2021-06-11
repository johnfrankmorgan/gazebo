package ast

var (
	_ Stmt = &SBlock{}
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

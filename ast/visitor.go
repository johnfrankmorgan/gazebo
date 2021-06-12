package ast

type ExprVisitor interface {
	VisitEGroup(*EGroup)
	VisitEBinary(*EBinary)
	VisitEUnary(*EUnary)
	VisitELiteral(*ELiteral)
}

type StmtVisitor interface {
	VisitSBlock(*SBlock)
	VisitSAssign(*SAssign)
	VisitSExpr(*SExpr)
}

type Visitor interface {
	ExprVisitor
	StmtVisitor
}

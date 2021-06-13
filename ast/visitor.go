package ast

type ExprVisitor interface {
	VisitEGroup(*EGroup)
	VisitEBinary(*EBinary)
	VisitEUnary(*EUnary)
	VisitELiteral(*ELiteral)
	VisitEFuncDef(*EFuncDef)
}

type StmtVisitor interface {
	VisitSBlock(*SBlock)
	VisitSAssign(*SAssign)
	VisitSExpr(*SExpr)
	VisitSIf(*SIf)
}

type Visitor interface {
	ExprVisitor
	StmtVisitor
}

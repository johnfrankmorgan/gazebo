package ast

type ExprVisitor interface {
	VisitEAssign(*EAssign)
	VisitEGroup(*EGroup)
	VisitEBinary(*EBinary)
	VisitEUnary(*EUnary)
	VisitELiteral(*ELiteral)
	VisitEFuncDef(*EFuncDef)
}

type StmtVisitor interface {
	VisitSBlock(*SBlock)
	VisitSExpr(*SExpr)
	VisitSIf(*SIf)
	VisitSWhile(*SWhile)
	VisitSReturn(*SReturn)
}

type Visitor interface {
	ExprVisitor
	StmtVisitor
}

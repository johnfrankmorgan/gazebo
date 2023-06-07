package gazebo

func compile(s *syntax) *code {
	c := &code{}
	s.compile(c)
	return c
}

func (node *syntax) compile(c *code) {
	for _, stmt := range node.Stmts {
		stmt.compile(c)
	}
}

func (node *syntaxStmt) compile(code *code) {
	switch {
	case node.Block != nil:
		node.Block.compile(code)

	case node.Print != nil:
		node.Print.compile(code)

	case node.Dump != nil:
		node.Dump.compile(code)

	case node.If != nil:
		node.If.compile(code)

	case node.While != nil:
		node.While.compile(code)

	case node.Return != nil:
		node.Return.compile(code)

	case node.Assign != nil:
		node.Assign.compile(code)

	case node.Expr != nil:
		node.Expr.compile(code)

	default:
		unreachable()
	}
}

func (node *syntaxStmtBlock) compile(c *code) {
	c = &code{
		parent: c,
	}

	for _, stmt := range node.Stmts {
		stmt.compile(c)
	}

	c.parent.emit(opExecChild, c.parent.child(c))
}

func (node *syntaxStmtPrint) compile(c *code) {
	for i := len(node.Exprs) - 1; i >= 0; i-- {
		node.Exprs[i].compile(c)
	}

	c.emit(opPrint, len(node.Exprs))
}

func (node *syntaxStmtDump) compile(c *code) {
	node.Expr.compile(c)
	c.emit(opDump)
}

func (node *syntaxStmtIf) compile(c *code) {
	node.Condition.compile(c)

	body := &code{}
	node.Body.compile(body)

	c.emit(opRelJumpIfFalse, 2)
	c.emit(opExecChild, c.child(body))

	if node.Else != nil {
		els := &code{}
		node.Else.compile(els)

		c.emit(opRelJump, 2)
		c.emit(opExecChild, c.child(els))
	}
}

func (node *syntaxStmtWhile) compile(c *code) {
	pc := c.pc()

	node.Condition.compile(c)

	body := &code{}
	node.Body.compile(body)

	c.emit(opRelJumpIfFalse, 2)
	c.emit(opExecChild, c.child(body))
	c.emit(opJump, pc)
}

func (node *syntaxStmtReturn) compile(c *code) {
	if node.Expr != nil {
		node.Expr.compile(c)
		c.emit(opReturn)
	} else {
		c.emit(opReturnNull)
	}
}

func (node *syntaxStmtAssign) compile(c *code) {
	node.Expr.compile(c)

	c.emit(opStoreName, c.name(node.Ident))
}

func (node *syntaxExpr) compile(c *code) {
	node.Logical.compile(c)
}

func (node *syntaxExprLogical) compile(c *code) {
	node.Left.compile(c)

	if node.Right != nil {
		node.Right.compile(c)

		switch *node.Op {
		case "and":
			c.emit(opBinAnd)

		case "or":
			c.emit(opBinOr)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprEquality) compile(c *code) {
	node.Left.compile(c)

	if node.Right != nil {
		node.Right.compile(c)

		switch *node.Op {
		case "==":
			c.emit(opBinEqual)

		case "!=":
			c.emit(opBinNotEqual)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprComparison) compile(c *code) {
	node.Left.compile(c)

	if node.Right != nil {
		node.Right.compile(c)

		switch *node.Op {
		case ">=":
			c.emit(opBinGreaterEqual)

		case "<=":
			c.emit(opBinLessEqual)

		case ">":
			c.emit(opBinGreater)

		case "<":
			c.emit(opBinLess)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprAddition) compile(c *code) {
	node.Left.compile(c)

	if node.Right != nil {
		node.Right.compile(c)

		switch *node.Op {
		case "+":
			c.emit(opBinAdd)

		case "-":
			c.emit(opBinSub)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprMultiplication) compile(c *code) {
	node.Left.compile(c)

	if node.Right != nil {
		node.Right.compile(c)

		switch *node.Op {
		case "*":
			c.emit(opBinMul)

		case "/":
			c.emit(opBinDiv)

		case "%":
			c.emit(opBinMod)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprUnary) compile(c *code) {
	if node.Primary != nil {
		node.Primary.compile(c)
	} else {
		node.Unary.compile(c)

		switch *node.Op {
		case "!":
			c.emit(opUnNot)

		case "-":
			c.emit(opUnNegate)

		default:
			unreachable()
		}
	}
}

func (node *syntaxExprPrimary) compile(c *code) {
	switch {
	case node.Group != nil:
		node.Group.compile(c)

	case node.Literal != nil:
		node.Literal.compile(c)

	case node.Ident != nil:
		switch *node.Ident {
		case "true":
			c.emit(opLoadTrue)

		case "false":
			c.emit(opLoadFalse)

		case "null":
			c.emit(opLoadNull)

		default:
			c.emit(opLoadName, c.name(*node.Ident))
		}

	default:
		unreachable()
	}
}

func (node *syntaxExprLiteral) compile(c *code) {
	switch {
	case node.Int != nil:
		c.emit(opLoadConst, c.constant(NewIntObject(*node.Int).AsObject()))

	case node.String != nil:
		c.emit(opLoadConst, c.constant(NewStringObject(*node.String).AsObject()))

	default:
		unreachable()
	}
}

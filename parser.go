package gazebo

func init() {
	yyErrorVerbose = true
}

func Parse(source string) (*AST, error) {
	l := NewLexer(source, &AST{})

	yyParse(l)

	return l.ast, l.err
}

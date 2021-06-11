package main

import (
	"flag"

	"github.com/johnfrankmorgan/gazebo/parser"
	"github.com/niemeyer/pretty"
)

var config struct {
	dump struct {
		ast bool
	}
}

func main() {
	flag.BoolVar(&config.dump.ast, "A", true, "Dump out the AST")
	flag.Parse()

	program := flag.Args()[0]
	parser := parser.New(parser.Tokenize(program))

	tree := parser.Parse()

	if config.dump.ast {
		pretty.Printf("%# v", tree)
	}
}

package main

import "flag"

var config struct {
	dump struct {
		ast bool
	}
}

func main() {
	flag.BoolVar(&config.dump.ast, "A", true, "Dump out the AST")
	flag.Parse()
}

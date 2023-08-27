package main

import (
	"github.com/alecthomas/kong"
)

func main() {
	cli := struct {
		Parse   Parse   `cmd:"" help:"Parses the specified file(s) and prints the resulting AST(s)."`
		Compile Compile `cmd:"" help:"Compiles the specified file(s) and prints the resulting bytecode."`
	}{}

	ctx := kong.Parse(&cli)

	ctx.FatalIfErrorf(ctx.Run())
}

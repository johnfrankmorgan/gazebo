package main

import (
	"fmt"
	"gazebo/util/must"
	"os"
	"runtime/pprof"

	"github.com/alecthomas/kong"
)

func main() {
	cli := struct {
		Run     Run     `cmd:"" default:"withargs" help:"Executes the specified file."`
		Parse   Parse   `cmd:"" help:"Parses the specified file(s) and prints the resulting AST(s)."`
		Compile Compile `cmd:"" help:"Compiles the specified file(s) and prints the resulting bytecode."`
		Profile string  `help:"Write profiling information here."`
	}{}

	ctx := kong.Parse(&cli)

	if cli.Profile != "" {
		f, err := os.Create(cli.Profile)
		ctx.FatalIfErrorf(err)

		defer must.Close(f)

		ctx.FatalIfErrorf(pprof.StartCPUProfile(f))

		defer pprof.StopCPUProfile()

		fmt.Println("writing profiling information to", f.Name())
	}

	ctx.FatalIfErrorf(ctx.Run())
}

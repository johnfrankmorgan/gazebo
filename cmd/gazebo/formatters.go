package main

import (
	"encoding/json"
	"io"

	"github.com/alecthomas/repr"
	"gopkg.in/yaml.v3"
)

var Formatters = map[string]func(io.Writer, any) error{
	"go": func(out io.Writer, program any) error {
		repr.New(out).Println(program)
		return nil
	},

	"json": func(out io.Writer, program any) error {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(program)
	},

	"yaml": func(out io.Writer, program any) error {
		enc := yaml.NewEncoder(out)
		enc.SetIndent(2)
		return enc.Encode(program)
	},
}

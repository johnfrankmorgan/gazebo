package main

import (
	"fmt"
	"os"

	"github.com/johnfrankmorgan/gazebo"
)

func main() {
	if err := gazebo.Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

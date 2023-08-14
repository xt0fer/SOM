package main

import (
	"os"

	"github.com/xt0fer/som/smog"
)

func main() {
	u := &smog.Universe{}
	args2 := os.Args[1:]
	u.Interpret(args2)
	u.Exit(0)
}

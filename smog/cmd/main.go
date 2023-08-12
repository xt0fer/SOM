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

// Main = (
// 	run: args = (
// 	  | u args2 |
// 	  u := Universe new.
// 	  args2 := args copyFrom: 2.
// 	  u interpret: args2.
// 	  u exit: 0.
// 	)
//   )

package main

import (
	"os"

	"labs.lesiw.io/ops/golib"
	"lesiw.io/ops"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "check")
	}
	ops.Handle(golib.Ops{})
}

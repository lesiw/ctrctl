package main

import (
	"bufio"
	"os"
	"strings"

	"labs.lesiw.io/ops/golib"
	"lesiw.io/cmdio/sys"
	"lesiw.io/ops"
)

type Ops struct {
	Bump func()
	golib.Ops
}

var Runner = sys.Runner()

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "check")
	}
	op := Ops{}
	op.Bump = op.Ops.Bump
	ops.Handle(op)
}

func (op Ops) Update() {
	if _, err := Runner.Get("which", "goimports"); err != nil {
		Runner.MustRun("go", "install",
			"golang.org/x/tools/cmd/goimports@latest")
	}
	Runner.MustRun("go", "generate", "./...")
	if Runner.MustGet("git", "status", "-s").Out == "" {
		return // Nothing changed.
	}
	if gitconfig("user.name") == "" {
		Runner.MustRun("git", "config", "--global", "user.name", "llbot")
		Runner.MustRun("git", "config", "--global", "user.email",
			"llbot@noreply.lesiwlabs.com")
	}
	Runner.MustRun("git", "commit", "-am", "Automated update")
	Runner.MustRun("op") // Validate changes.
	op.Bump()
	Runner.MustRun("git", "push")
}

func gitconfig(key string) string {
	cmd := Runner.Command("git", "config", "-l")
	scanner := bufio.NewScanner(cmd)
	for scanner.Scan() {
		k, v, ok := strings.Cut(scanner.Text(), "=")
		if ok && k == key {
			return v
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ""
}

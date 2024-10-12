package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"labs.lesiw.io/ops/git"
	"labs.lesiw.io/ops/golang"
	"lesiw.io/cmdio"
)

var uniq = make(map[string]int)

var bump func()

func (Ops) Bump() { bump() }

type testCdr struct {
	cmd        [][]string
	gitChanged bool
	gitConfig  string
}

func (c *testCdr) Command(
	_ context.Context, env map[string]string, args ...string,
) io.ReadWriter {
	c.cmd = append(c.cmd, args)
	s := fmt.Sprintf("%v", args)
	uniq[s]++
	if uniq[s] > 1 {
		s = fmt.Sprintf("%s[%d]", s, uniq[s])
	}
	if !c.gitChanged && slices.Equal(args, []string{"git", "status", "-s"}) {
		return bytes.NewBufferString("")
	} else if slices.Equal(args, []string{"git", "config", "-l"}) {
		return bytes.NewBufferString(c.gitConfig)
	} else {
		return bytes.NewBufferString(s + "\n")
	}
}

func TestUpdate(t *testing.T) {
	t.Cleanup(func() { clear(uniq) })
	cdr := new(testCdr)
	cdr.gitChanged = true
	Runner = cmdio.NewRunner(context.Background(), nil, cdr)
	golang.Runner = func() *cmdio.Runner { return Runner }
	git.Runner = Runner

	bump = func() { Runner.MustRun("bump") }
	Ops{}.Update()

	expected := [][]string{
		{"which", "goimports"},
		{"go", "generate", "./..."},
		{"git", "status", "-s"},
		{"git", "config", "-l"},
		{"git", "config", "--global", "user.name", "llbot"},
		{"git", "config", "--global", "user.email",
			"llbot@noreply.lesiwlabs.com"},
		{"git", "commit", "-am", "Automated update"},
		{"op"},
		{"bump"},
		{"git", "push"},
	}
	if got, want := cdr.cmd, expected; !cmp.Equal(got, want) {
		t.Errorf("commands -want +got:\n%s", cmp.Diff(got, want))
	}
}

func TestNoChange(t *testing.T) {
	t.Cleanup(func() { clear(uniq) })
	cdr := new(testCdr)
	Runner = cmdio.NewRunner(context.Background(), nil, cdr)
	golang.Runner = func() *cmdio.Runner { return Runner }
	git.Runner = Runner

	bump = func() { Runner.MustRun("bump") }
	Ops{}.Update()

	expected := [][]string{
		{"which", "goimports"},
		{"go", "generate", "./..."},
		{"git", "status", "-s"},
	}
	if got, want := cdr.cmd, expected; !cmp.Equal(got, want) {
		t.Errorf("commands -want +got:\n%s", cmp.Diff(got, want))
	}
}

func TestGitUserExists(t *testing.T) {
	t.Cleanup(func() { clear(uniq) })
	cdr := new(testCdr)
	cdr.gitChanged = true
	cdr.gitConfig = "user.name=example"
	Runner = cmdio.NewRunner(context.Background(), nil, cdr)
	golang.Runner = func() *cmdio.Runner { return Runner }
	git.Runner = Runner

	bump = func() { Runner.MustRun("bump") }
	Ops{}.Update()

	expected := [][]string{
		{"which", "goimports"},
		{"go", "generate", "./..."},
		{"git", "status", "-s"},
		{"git", "config", "-l"},
		{"git", "commit", "-am", "Automated update"},
		{"op"},
		{"bump"},
		{"git", "push"},
	}
	if got, want := cdr.cmd, expected; !cmp.Equal(got, want) {
		t.Errorf("commands -want +got:\n%s", cmp.Diff(got, want))
	}
}

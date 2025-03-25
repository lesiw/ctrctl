//go:generate go run internal/generate.go

// Package ctrctl wraps container CLIs.
package ctrctl

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

// Cli is the command prefix.
var Cli []string

var clis = [...][]string{
	{"docker"},
	{"podman"},
	{"nerdctl"},
	{"lima", "nerdctl"},
}

var findCli = sync.OnceValue(func() error {
	if Cli != nil {
		return nil
	}
	for _, cli := range clis {
		if _, err := exec.LookPath(cli[0]); err == nil {
			Cli = cli
			return nil
		}
	}
	return fmt.Errorf("no container cli found. searched: %+v", clis)
})

// Verbose mode prints the commands being run and streams their output to the
// terminal.
var Verbose bool

var shUnsafe = regexp.MustCompile(`[^\w@%+=:,./-]`)

type CliError struct {
	*os.ProcessState

	Stderr string
}

func (e *CliError) Error() string {
	if e.Stderr != "" {
		return e.Stderr
	} else {
		return e.String()
	}
}

func runCtrCmd(
	subcommand []string, args []string, opts any, optpos int,
) (string, error) {
	var strout strings.Builder
	var strerr strings.Builder
	var cmd *exec.Cmd

	cmdargs := append(Cli, subcommand...)
	for i := 0; i < len(args); i++ {
		if optpos > -1 && i == optpos {
			if !reflect.ValueOf(opts).IsNil() {
				cmdargs = append(cmdargs, optsToArgs(opts)...)
			}
			optpos = -1
			i--
		} else if !reflect.ValueOf(args[i]).IsZero() {
			cmdargs = append(cmdargs, args[i])
		}
	}

	if !reflect.ValueOf(opts).IsNil() {
		var ok bool
		cmd, ok = reflect.ValueOf(opts).Elem().FieldByName("Cmd").
			Interface().(*exec.Cmd)
		if !ok {
			panic("exec.Cmd failed type assertion")
		}
	}

	if cmd == nil {
		cmd = &exec.Cmd{}
	}
	cmd.Path = cmdargs[0]
	cmd.Args = cmdargs
	if filepath.Base(cmd.Path) == cmd.Path {
		lp, err := exec.LookPath(cmd.Path)
		if lp != "" {
			cmd.Path = lp
		}
		if err != nil {
			cmd.Err = err
		}
	}

	prepareStreams(cmd, &strout, &strerr)

	if Verbose {
		fmt.Fprintf(os.Stderr, "+ %s\n", shJoin(cmdargs))
	}

	err := cmd.Run()
	if ee, ok := err.(*exec.ExitError); ok {
		err = &CliError{
			ProcessState: ee.ProcessState,
			Stderr:       strings.TrimSpace(strerr.String()),
		}
	}

	return strings.TrimSpace(strout.String()), err
}

func optsToArgs(opts interface{}) []string {
	result := []string{}
	val := reflect.ValueOf(opts).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		value := val.Field(i)
		field := typ.Field(i)

		if field.Name == "Cmd" {
			continue
		}

		if field.Type.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			val = val.Elem()
		}

		if field.Type.Kind() == reflect.String && value.IsZero() {
			continue
		}
		if field.Type.Kind() == reflect.Bool && !value.Bool() {
			continue
		}
		if field.Type.Kind() == reflect.Slice {
			for j := 0; j < value.Len(); j++ {
				str := value.Index(j)
				result = append(result, fieldToFlag(field.Name), str.String())
			}
		} else {
			result = append(result, fieldToFlag(field.Name))
			switch field.Type.Kind() {
			case reflect.Int:
				result = append(result, strconv.FormatInt(value.Int(), 10))
			case reflect.String:
				result = append(result, value.String())
			}
		}
	}

	return result
}

func fieldToFlag(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '-')
			}
		}
		result = append(result, unicode.ToLower(r))
	}
	if len(result) == 1 {
		return fmt.Sprintf("-%s", string(result))
	} else {
		return fmt.Sprintf("--%s", string(result))
	}
}

func shQuote(s string) string {
	if s == "" {
		return `''`
	}
	if !shUnsafe.MatchString(s) {
		return s
	}
	return `'` + strings.ReplaceAll(s, `'`, `\'`) + `'`
}

func shJoin(parts []string) string {
	quotedParts := make([]string, len(parts))
	for i, part := range parts {
		quotedParts[i] = shQuote(part)
	}
	return strings.Join(quotedParts, " ")
}

func prepareStreams(cmd *exec.Cmd, out io.Writer, err io.Writer) {
	if cmd.Stdout == nil {
		cmd.Stdout = out
	}
	if Verbose && cmd.Stdout != os.Stdout {
		cmd.Stdout = io.MultiWriter(cmd.Stdout, os.Stdout)
	}
	if cmd.Stderr == nil {
		cmd.Stderr = err
	}
	if Verbose && cmd.Stderr != os.Stderr {
		cmd.Stderr = io.MultiWriter(cmd.Stderr, os.Stderr)
	}
}

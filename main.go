//go:generate go run internal/generate.go

// Package ctrctl provides a wrapper around container CLI commands.
package ctrctl

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Cli prefixes every command.
var Cli = []string{"docker"}

// Verbose mode prints the commands being run and streams output to the terminal.
var Verbose bool

var shUnsafe = regexp.MustCompile(`[^\w@%+=:,./-]`)

func runCtrCmd(subcommand []string, args []string, opts interface{}, optpos int) (
	stdout string, stderr string, err error) {

	var strout strings.Builder
	var strerr strings.Builder
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

	cmd := exec.Command(cmdargs[0], cmdargs[1:]...)
	if Verbose {
		fmt.Fprintf(os.Stderr, "+ %s\n", shJoin(cmdargs))
		cmd.Stdout = io.MultiWriter(&strout, os.Stdout)
		cmd.Stderr = io.MultiWriter(&strout, os.Stderr)
	} else {
		cmd.Stdout = &strout
		cmd.Stderr = &strerr
	}

	err = cmd.Run()

	stdout = strings.TrimSpace(strout.String())
	stderr = strings.TrimSpace(strerr.String())

	return
}

func optsToArgs(opts interface{}) []string {
	result := []string{}
	val := reflect.ValueOf(opts).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		value := val.Field(i)
		field := typ.Field(i)

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
		result = append(result, fieldToFlag(field.Name))

		switch field.Type.Kind() {
		case reflect.Int:
			result = append(result, strconv.FormatInt(value.Int(), 10))
		case reflect.String:
			result = append(result, value.String())
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

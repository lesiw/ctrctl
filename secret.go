package ctrctl

import (
	"fmt"
	"os/exec"
)

type SecretCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Secret driver.
	Driver string

	// Print usage.
	Help bool

	// Secret labels.
	Label string

	// Template driver.
	TemplateDriver string
}

// Create a secret from a file or STDIN as content.
func SecretCreate(opts *SecretCreateOpts, secret string, file string) (string, error) {
	return runCtrCmd(
		[]string{"secret", "create"},
		[]string{secret, file},
		opts,
		0,
	)
}

type SecretInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more secrets.
func SecretInspect(opts *SecretInspectOpts, secret ...string) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("secret must have at least one value")
	}
	return runCtrCmd(
		[]string{"secret", "inspect"},
		secret,
		opts,
		0,
	)
}

type SecretLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Only display IDs.
	Quiet bool
}

// List secrets.
func SecretLs(opts *SecretLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"secret", "ls"},
		[]string{},
		opts,
		0,
	)
}

type SecretRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Remove one or more secrets.
func SecretRm(opts *SecretRmOpts, secret ...string) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("secret must have at least one value")
	}
	return runCtrCmd(
		[]string{"secret", "rm"},
		secret,
		opts,
		-1,
	)
}

package ctrctl

import (
	"fmt"
	"os/exec"
)

type ContextCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Description of the context.
	Description string

	// set the docker endpoint.
	Docker string

	// create context from a named context.
	From string

	// Print usage.
	Help bool
}

// Create a context.
func ContextCreate(opts *ContextCreateOpts, context string) (string, error) {
	return runCtrCmd(
		[]string{"context", "create"},
		[]string{context},
		opts,
		0,
	)
}

type ContextExportOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Export a context to a tar archive FILE or a tar stream on STDOUT.
func ContextExport(opts *ContextExportOpts, context string, file string) (string, error) {
	return runCtrCmd(
		[]string{"context", "export"},
		[]string{context, file},
		opts,
		0,
	)
}

type ContextImportOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Import a context from a tar or zip file.
func ContextImport(opts *ContextImportOpts, context string, file string) (string, error) {
	return runCtrCmd(
		[]string{"context", "import"},
		[]string{context, file},
		opts,
		-1,
	)
}

type ContextInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool
}

// Display detailed information on one or more contexts.
func ContextInspect(opts *ContextInspectOpts, context ...string) (string, error) {
	if len(context) == 0 {
		return "", fmt.Errorf("context must have at least one value")
	}
	return runCtrCmd(
		[]string{"context", "inspect"},
		context,
		opts,
		0,
	)
}

type ContextLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Only show context names.
	Quiet bool
}

// List contexts.
func ContextLs(opts *ContextLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"context", "ls"},
		[]string{},
		opts,
		0,
	)
}

type ContextRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force the removal of a context in use.
	Force bool

	// Print usage.
	Help bool
}

// Remove one or more contexts.
func ContextRm(opts *ContextRmOpts, context ...string) (string, error) {
	if len(context) == 0 {
		return "", fmt.Errorf("context must have at least one value")
	}
	return runCtrCmd(
		[]string{"context", "rm"},
		context,
		opts,
		-1,
	)
}

type ContextShowOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Print the name of the current context.
func ContextShow(opts *ContextShowOpts) (string, error) {
	return runCtrCmd(
		[]string{"context", "show"},
		[]string{},
		opts,
		-1,
	)
}

type ContextUpdateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Description of the context.
	Description string

	// set the docker endpoint.
	Docker string

	// Print usage.
	Help bool
}

// Update a context.
func ContextUpdate(opts *ContextUpdateOpts, context string) (string, error) {
	return runCtrCmd(
		[]string{"context", "update"},
		[]string{context},
		opts,
		0,
	)
}

type ContextUseOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Set the current docker context.
func ContextUse(opts *ContextUseOpts, context string) (string, error) {
	return runCtrCmd(
		[]string{"context", "use"},
		[]string{context},
		opts,
		-1,
	)
}

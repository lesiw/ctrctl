package ctrctl

type ContextLsOpts struct {
	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Only show context names.
	Quiet string
}

// List contexts.
func ContextLs(opts *ContextLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "context", "ls" },
		[]string{},
		opts,
		0,
	)
}

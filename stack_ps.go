package ctrctl

type StackPsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Do not map IDs to Names.
	NoResolve string

	// Do not truncate output.
	NoTrunc string

	// Only display task IDs.
	Quiet string
}

// List the tasks in the stack.
func StackPs(opts *StackPsOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack", "ps" },
		[]string{ stack },
		opts,
		0,
	)
}

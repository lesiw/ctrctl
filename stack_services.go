package ctrctl

type StackServicesOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Only display IDs.
	Quiet bool
}

// List the services in the stack.
func StackServices(opts *StackServicesOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack", "services" },
		[]string{ stack },
		opts,
		0,
	)
}

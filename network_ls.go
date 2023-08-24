package ctrctl

type NetworkLsOpts struct {
	// Provide filter values (e.g. `driver=bridge`).
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Do not truncate the output.
	NoTrunc bool

	// Only display network IDs.
	Quiet bool
}

// List networks.
func NetworkLs(opts *NetworkLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network", "ls" },
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type NetworkInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Verbose output for diagnostics.
	Verbose bool
}

// Display detailed information on one or more networks.
func NetworkInspect(opts *NetworkInspectOpts, network string, extraNetwork ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network", "inspect" },
		append([]string{ network }, extraNetwork...),
		opts,
		0,
	)
}

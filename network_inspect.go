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
func NetworkInspect(opts *NetworkInspectOpts, network ...string) (
	stdout string, stderr string, err error) {
	if len(network) == 0 {
		return "", "", fmt.Errorf("network must have at least one value")
	}
	return runCtrCmd(
		[]string{"network", "inspect"},
		network,
		opts,
		0,
	)
}

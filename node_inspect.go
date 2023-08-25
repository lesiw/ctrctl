package ctrctl

type NodeInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more nodes.
func NodeInspect(opts *NodeInspectOpts, selfNode string, node ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"node", "inspect"},
		append([]string{selfNode}, node...),
		opts,
		0,
	)
}

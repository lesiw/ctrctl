package ctrctl

type NodePsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Do not map IDs to Names.
	NoResolve string

	// Do not truncate output.
	NoTrunc string

	// Only display task IDs.
	Quiet string
}

// List tasks running on one or more nodes, defaults to current node.
func NodePs(opts *NodePsOpts, node ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "node", "ps" },
		[]string{},
		opts,
		0,
	)
}

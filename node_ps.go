package ctrctl

type NodePsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Only display task IDs.
	Quiet bool
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

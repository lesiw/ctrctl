package ctrctl

type NodeRmOpts struct {
	// Force remove a node from the swarm.
	Force string
}

// Remove one or more nodes from the swarm.
func NodeRm(opts *NodeRmOpts, node string, extraNode ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "node", "rm" },
		append([]string{ node }, extraNode...),
		opts,
		0,
	)
}

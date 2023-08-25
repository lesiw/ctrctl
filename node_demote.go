package ctrctl

type NodeDemoteOpts struct {
}

// Demote one or more nodes from manager in the swarm.
func NodeDemote(opts *NodeDemoteOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "demote"},
		node,
		opts,
		-1,
	)
}

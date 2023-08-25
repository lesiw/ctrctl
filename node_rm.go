package ctrctl

import "fmt"

type NodeRmOpts struct {
	// Force remove a node from the swarm.
	Force bool
}

// Remove one or more nodes from the swarm.
func NodeRm(opts *NodeRmOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{ "node", "rm" },
		node,
		opts,
		0,
	)
}

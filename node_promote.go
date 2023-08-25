package ctrctl

import "fmt"

type NodePromoteOpts struct {

}

// Promote one or more nodes to manager in the swarm.
func NodePromote(opts *NodePromoteOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{ "node", "promote" },
		node,
		opts,
		-1,
	)
}

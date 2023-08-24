package ctrctl

type NodePromoteOpts struct {

}

// Promote one or more nodes to manager in the swarm.
func NodePromote(opts *NodePromoteOpts, node string, extraNode ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "node", "promote" },
		append([]string{ node }, extraNode...),
		opts,
		-1,
	)
}

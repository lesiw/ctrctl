package ctrctl

type NodeDemoteOpts struct {

}

// Demote one or more nodes from manager in the swarm.
func NodeDemote(opts *NodeDemoteOpts, node string, extraNode ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "node", "demote" },
		append([]string{ node }, extraNode...),
		opts,
		-1,
	)
}

package ctrctl

type NodeOpts struct {

}

// Manage Swarm nodes.
func Node(opts *NodeOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "node" },
		[]string{},
		opts,
		-1,
	)
}

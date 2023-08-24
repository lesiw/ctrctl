package ctrctl

type SwarmLeaveOpts struct {
	// Force this node to leave the swarm, ignoring warnings.
	Force bool
}

// Leave the swarm.
func SwarmLeave(opts *SwarmLeaveOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "swarm", "leave" },
		[]string{},
		opts,
		0,
	)
}

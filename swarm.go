package ctrctl

type SwarmOpts struct {
}

// Manage Swarm.
func Swarm(opts *SwarmOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm"},
		[]string{},
		opts,
		-1,
	)
}

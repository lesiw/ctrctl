package ctrctl

type NetworkOpts struct {

}

// Manage networks.
func Network(opts *NetworkOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network" },
		[]string{},
		opts,
		-1,
	)
}

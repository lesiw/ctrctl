package ctrctl

type NetworkRmOpts struct {
	// Do not error if the network does not exist.
	Force bool
}

// Remove one or more networks.
func NetworkRm(opts *NetworkRmOpts, network string, extraNetwork ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network", "rm" },
		append([]string{ network }, extraNetwork...),
		opts,
		-1,
	)
}

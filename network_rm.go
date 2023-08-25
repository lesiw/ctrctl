package ctrctl

type NetworkRmOpts struct {
	// Do not error if the network does not exist.
	Force bool
}

// Remove one or more networks.
func NetworkRm(opts *NetworkRmOpts, network ...string) (
	stdout string, stderr string, err error) {
	if len(network) == 0 {
		return "", "", fmt.Errorf("network must have at least one value")
	}
	return runCtrCmd(
		[]string{"network", "rm"},
		network,
		opts,
		-1,
	)
}

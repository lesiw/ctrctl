package ctrctl

type NetworkDisconnectOpts struct {
	// Force the container to disconnect from a network.
	Force bool
}

// Disconnect a container from a network.
func NetworkDisconnect(opts *NetworkDisconnectOpts, network string, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"network", "disconnect"},
		[]string{network, container},
		opts,
		0,
	)
}

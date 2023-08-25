package ctrctl

type NetworkConnectOpts struct {
	// Add network-scoped alias for the container.
	Alias string

	// driver options for the network.
	DriverOpt string

	// IPv4 address (e.g., `172.30.100.104`).
	Ip string

	// IPv6 address (e.g., `2001:db8::33`).
	Ip6 string

	// Add link to another container.
	Link string

	// Add a link-local address for the container.
	LinkLocalIp string
}

// Connect a container to a network.
func NetworkConnect(opts *NetworkConnectOpts, network string, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"network", "connect"},
		[]string{network, container},
		opts,
		0,
	)
}

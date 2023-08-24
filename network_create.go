package ctrctl

type NetworkCreateOpts struct {
	// Enable manual container attachment.
	Attachable string

	// Auxiliary IPv4 or IPv6 addresses used by Network driver.
	AuxAddress string

	// The network from which to copy the configuration.
	ConfigFrom string

	// Create a configuration only network.
	ConfigOnly string

	// Driver to manage the Network.
	Driver string

	// IPv4 or IPv6 Gateway for the master subnet.
	Gateway string

	// Create swarm routing-mesh network.
	Ingress string

	// Restrict external access to the network.
	Internal string

	// Allocate container ip from a sub-range.
	IpRange string

	// IP Address Management Driver.
	IpamDriver string

	// Set IPAM driver specific options.
	IpamOpt string

	// Enable IPv6 networking.
	Ipv6 string

	// Set metadata on a network.
	Label string

	// Set driver specific options.
	Opt string

	// Control the network's scope.
	Scope string

	// Subnet in CIDR format that represents a network segment.
	Subnet string
}

// Create a network.
func NetworkCreate(opts *NetworkCreateOpts, network string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network", "create" },
		[]string{ network },
		opts,
		0,
	)
}

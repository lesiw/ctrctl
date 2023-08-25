package ctrctl

import "fmt"

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

type NetworkCreateOpts struct {
	// Enable manual container attachment.
	Attachable bool

	// Auxiliary IPv4 or IPv6 addresses used by Network driver.
	AuxAddress string

	// The network from which to copy the configuration.
	ConfigFrom string

	// Create a configuration only network.
	ConfigOnly bool

	// Driver to manage the Network.
	Driver string

	// IPv4 or IPv6 Gateway for the master subnet.
	Gateway string

	// Create swarm routing-mesh network.
	Ingress bool

	// Restrict external access to the network.
	Internal bool

	// Allocate container ip from a sub-range.
	IpRange string

	// IP Address Management Driver.
	IpamDriver string

	// Set IPAM driver specific options.
	IpamOpt string

	// Enable IPv6 networking.
	Ipv6 bool

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
		[]string{"network", "create"},
		[]string{network},
		opts,
		0,
	)
}

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

type NetworkInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Verbose output for diagnostics.
	Verbose bool
}

// Display detailed information on one or more networks.
func NetworkInspect(opts *NetworkInspectOpts, network ...string) (
	stdout string, stderr string, err error) {
	if len(network) == 0 {
		return "", "", fmt.Errorf("network must have at least one value")
	}
	return runCtrCmd(
		[]string{"network", "inspect"},
		network,
		opts,
		0,
	)
}

type NetworkLsOpts struct {
	// Provide filter values (e.g. `driver=bridge`).
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Do not truncate the output.
	NoTrunc bool

	// Only display network IDs.
	Quiet bool
}

// List networks.
func NetworkLs(opts *NetworkLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"network", "ls"},
		[]string{},
		opts,
		0,
	)
}

type NetworkPruneOpts struct {
	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool
}

// Remove all unused networks.
func NetworkPrune(opts *NetworkPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"network", "prune"},
		[]string{},
		opts,
		0,
	)
}

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

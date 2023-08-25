package ctrctl

type SwarmJoinOpts struct {
	// Advertised address (format: `<ip|interface>[:port]`).
	AdvertiseAddr string

	// Availability of the node (`active`, `pause`, `drain`).
	Availability string

	// Address or interface to use for data path traffic (format: `<ip|interface>`).
	DataPathAddr string

	// Listen address (format: `<ip|interface>[:port]`).
	ListenAddr string

	// Token for entry into the swarm.
	Token string
}

// Join a swarm as a node and/or manager.
func SwarmJoin(opts *SwarmJoinOpts, hostPort string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "join"},
		[]string{hostPort},
		opts,
		0,
	)
}

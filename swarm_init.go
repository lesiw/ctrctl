package ctrctl

type SwarmInitOpts struct {
	// Advertised address (format: `<ip|interface>[:port]`).
	AdvertiseAddr string

	// Enable manager autolocking (requiring an unlock key to start a stopped manager).
	Autolock string

	// Availability of the node (`active`, `pause`, `drain`).
	Availability string

	// Validity period for node certificates (ns|us|ms|s|m|h).
	CertExpiry string

	// Address or interface to use for data path traffic (format: `<ip|interface>`).
	DataPathAddr string

	// Port number to use for data path traffic (1024 - 49151). If no value is set or is set to 0, the default port (4789) is used.
	DataPathPort string

	// default address pool in CIDR format.
	DefaultAddrPool string

	// default address pool subnet mask length.
	DefaultAddrPoolMaskLength string

	// Dispatcher heartbeat period (ns|us|ms|s|m|h).
	DispatcherHeartbeat string

	// Specifications of one or more certificate signing endpoints.
	ExternalCa string

	// Force create a new cluster from current state.
	ForceNewCluster string

	// Listen address (format: `<ip|interface>[:port]`).
	ListenAddr string

	// Number of additional Raft snapshots to retain.
	MaxSnapshots string

	// Number of log entries between Raft snapshots.
	SnapshotInterval string

	// Task history retention limit.
	TaskHistoryLimit string
}

// Initialize a swarm.
func SwarmInit(opts *SwarmInitOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "swarm", "init" },
		[]string{},
		opts,
		0,
	)
}

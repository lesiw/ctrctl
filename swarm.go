package ctrctl

type SwarmCaOpts struct {
	// Path to the PEM-formatted root CA certificate to use for the new cluster.
	CaCert string

	// Path to the PEM-formatted root CA key to use for the new cluster.
	CaKey string

	// Validity period for node certificates (ns|us|ms|s|m|h).
	CertExpiry string

	// Exit immediately instead of waiting for the root rotation to converge.
	Detach bool

	// Specifications of one or more certificate signing endpoints.
	ExternalCa string

	// Suppress progress output.
	Quiet bool

	// Rotate the swarm CA - if no certificate or key are provided, new ones will be generated.
	Rotate bool
}

// Display and rotate the root CA.
func SwarmCa(opts *SwarmCaOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "ca"},
		[]string{},
		opts,
		0,
	)
}

type SwarmInitOpts struct {
	// Advertised address (format: `<ip|interface>[:port]`).
	AdvertiseAddr string

	// Enable manager autolocking (requiring an unlock key to start a stopped manager).
	Autolock bool

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
	ForceNewCluster bool

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
		[]string{"swarm", "init"},
		[]string{},
		opts,
		0,
	)
}

type SwarmJoinTokenOpts struct {
	// Only display token.
	Quiet bool

	// Rotate join token.
	Rotate bool
}

// Manage join tokens.
func SwarmJoinToken(opts *SwarmJoinTokenOpts, WorkerManager string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "join-token"},
		[]string{WorkerManager},
		opts,
		0,
	)
}

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

type SwarmLeaveOpts struct {
	// Force this node to leave the swarm, ignoring warnings.
	Force bool
}

// Leave the swarm.
func SwarmLeave(opts *SwarmLeaveOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "leave"},
		[]string{},
		opts,
		0,
	)
}

type SwarmUnlockKeyOpts struct {
	// Only display token.
	Quiet bool

	// Rotate unlock key.
	Rotate bool
}

// Manage the unlock key.
func SwarmUnlockKey(opts *SwarmUnlockKeyOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "unlock-key"},
		[]string{},
		opts,
		0,
	)
}

type SwarmUnlockOpts struct {
}

// Unlock swarm.
func SwarmUnlock(opts *SwarmUnlockOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "unlock"},
		[]string{},
		opts,
		-1,
	)
}

type SwarmUpdateOpts struct {
	// Change manager autolocking setting (true|false).
	Autolock bool

	// Validity period for node certificates (ns|us|ms|s|m|h).
	CertExpiry string

	// Dispatcher heartbeat period (ns|us|ms|s|m|h).
	DispatcherHeartbeat string

	// Specifications of one or more certificate signing endpoints.
	ExternalCa string

	// Number of additional Raft snapshots to retain.
	MaxSnapshots string

	// Number of log entries between Raft snapshots.
	SnapshotInterval string

	// Task history retention limit.
	TaskHistoryLimit string
}

// Update the swarm.
func SwarmUpdate(opts *SwarmUpdateOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"swarm", "update"},
		[]string{},
		opts,
		0,
	)
}

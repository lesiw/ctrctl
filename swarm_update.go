package ctrctl

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

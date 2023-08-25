package ctrctl

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

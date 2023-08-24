package ctrctl

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
		[]string{ "swarm", "unlock-key" },
		[]string{},
		opts,
		0,
	)
}

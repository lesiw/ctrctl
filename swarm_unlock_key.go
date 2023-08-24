package ctrctl

type SwarmUnlockKeyOpts struct {
	// Only display token.
	Quiet string

	// Rotate unlock key.
	Rotate string
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

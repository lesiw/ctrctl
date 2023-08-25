package ctrctl

type CheckpointLsOpts struct {
	// Use a custom checkpoint storage directory.
	CheckpointDir string
}

// List checkpoints for a container.
func CheckpointLs(opts *CheckpointLsOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"checkpoint", "ls"},
		[]string{container},
		opts,
		0,
	)
}

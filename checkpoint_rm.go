package ctrctl

type CheckpointRmOpts struct {
	// Use a custom checkpoint storage directory.
	CheckpointDir string
}

// Remove a checkpoint.
func CheckpointRm(opts *CheckpointRmOpts, container string, checkpoint string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "checkpoint", "rm" },
		[]string{ container, checkpoint },
		opts,
		0,
	)
}

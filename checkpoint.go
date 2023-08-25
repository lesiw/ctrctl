package ctrctl

type CheckpointOpts struct {
}

// Manage checkpoints.
func Checkpoint(opts *CheckpointOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"checkpoint"},
		[]string{},
		opts,
		-1,
	)
}

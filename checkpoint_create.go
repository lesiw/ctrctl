package ctrctl

type CheckpointCreateOpts struct {
	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Leave the container running after checkpoint.
	LeaveRunning bool
}

// Create a checkpoint from a running container.
func CheckpointCreate(opts *CheckpointCreateOpts, container string, checkpoint string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "checkpoint", "create" },
		[]string{ container, checkpoint },
		opts,
		0,
	)
}

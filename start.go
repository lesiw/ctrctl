package ctrctl

type StartOpts struct {
	// Attach STDOUT/STDERR and forward signals.
	Attach bool

	// Restore from this checkpoint.
	Checkpoint string

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Attach container's STDIN.
	Interactive bool
}

// Start one or more stopped containers.
func Start(opts *StartOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"start"},
		container,
		opts,
		0,
	)
}

package ctrctl

type StartOpts struct {
	// Attach STDOUT/STDERR and forward signals.
	Attach string

	// Restore from this checkpoint.
	Checkpoint string

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Attach container's STDIN.
	Interactive string
}

// Start one or more stopped containers.
func Start(opts *StartOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "start" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

package ctrctl

type ContainerStartOpts struct {
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
func ContainerStart(opts *ContainerStartOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "start" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

package ctrctl

type ContainerRestartOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Restart one or more containers.
func ContainerRestart(opts *ContainerRestartOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "restart"},
		container,
		opts,
		0,
	)
}

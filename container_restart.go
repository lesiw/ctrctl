package ctrctl

type ContainerRestartOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Restart one or more containers.
func ContainerRestart(opts *ContainerRestartOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "restart" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

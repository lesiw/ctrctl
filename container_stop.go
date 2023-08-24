package ctrctl

type ContainerStopOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Stop one or more running containers.
func ContainerStop(opts *ContainerStopOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "stop" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

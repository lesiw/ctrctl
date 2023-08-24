package ctrctl

type StopOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time string
}

// Stop one or more running containers.
func Stop(opts *StopOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stop" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

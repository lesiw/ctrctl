package ctrctl

type RestartOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time string
}

// Restart one or more containers.
func Restart(opts *RestartOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "restart" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

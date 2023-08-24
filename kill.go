package ctrctl

type KillOpts struct {
	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func Kill(opts *KillOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "kill" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

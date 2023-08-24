package ctrctl

type ContainerKillOpts struct {
	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func ContainerKill(opts *ContainerKillOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "kill" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

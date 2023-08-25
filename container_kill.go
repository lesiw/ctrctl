package ctrctl

type ContainerKillOpts struct {
	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func ContainerKill(opts *ContainerKillOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "kill"},
		container,
		opts,
		0,
	)
}

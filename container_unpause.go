package ctrctl

type ContainerUnpauseOpts struct {
}

// Unpause all processes within one or more containers.
func ContainerUnpause(opts *ContainerUnpauseOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "unpause"},
		container,
		opts,
		-1,
	)
}

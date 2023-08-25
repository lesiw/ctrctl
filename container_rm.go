package ctrctl

type ContainerRmOpts struct {
	// Force the removal of a running container (uses SIGKILL).
	Force bool

	// Remove the specified link.
	Link bool

	// Remove anonymous volumes associated with the container.
	Volumes bool
}

// Remove one or more containers.
func ContainerRm(opts *ContainerRmOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "rm"},
		container,
		opts,
		0,
	)
}

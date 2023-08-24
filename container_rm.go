package ctrctl

type ContainerRmOpts struct {
	// Force the removal of a running container (uses SIGKILL).
	Force string

	// Remove the specified link.
	Link string

	// Remove anonymous volumes associated with the container.
	Volumes string
}

// Remove one or more containers.
func ContainerRm(opts *ContainerRmOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "rm" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

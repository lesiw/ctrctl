package ctrctl

type RmOpts struct {
	// Force the removal of a running container (uses SIGKILL).
	Force string

	// Remove the specified link.
	Link string

	// Remove anonymous volumes associated with the container.
	Volumes string
}

// Remove one or more containers.
func Rm(opts *RmOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "rm" },
		append([]string{ container }, extraContainer...),
		opts,
		0,
	)
}

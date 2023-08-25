package ctrctl

type UnpauseOpts struct {
}

// Unpause all processes within one or more containers.
func Unpause(opts *UnpauseOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"unpause"},
		container,
		opts,
		-1,
	)
}

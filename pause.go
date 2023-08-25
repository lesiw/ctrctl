package ctrctl

type PauseOpts struct {
}

// Pause all processes within one or more containers.
func Pause(opts *PauseOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"pause"},
		container,
		opts,
		-1,
	)
}

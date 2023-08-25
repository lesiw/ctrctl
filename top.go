package ctrctl

type TopOpts struct {
}

// Display the running processes of a container.
func Top(opts *TopOpts, container string, psOptions string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"top"},
		[]string{container, psOptions},
		opts,
		-1,
	)
}

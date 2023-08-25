package ctrctl

type SystemOpts struct {
}

// Manage Docker.
func System(opts *SystemOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system"},
		[]string{},
		opts,
		-1,
	)
}

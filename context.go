package ctrctl

type ContextOpts struct {
}

// Manage contexts.
func Context(opts *ContextOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"context"},
		[]string{},
		opts,
		-1,
	)
}

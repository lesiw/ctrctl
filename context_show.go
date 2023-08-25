package ctrctl

type ContextShowOpts struct {
}

// Print the name of the current context.
func ContextShow(opts *ContextShowOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"context", "show"},
		[]string{},
		opts,
		-1,
	)
}

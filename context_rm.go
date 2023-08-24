package ctrctl

type ContextRmOpts struct {
	// Force the removal of a context in use.
	Force string
}

// Remove one or more contexts.
func ContextRm(opts *ContextRmOpts, context string, extraContext ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "context", "rm" },
		append([]string{ context }, extraContext...),
		opts,
		-1,
	)
}

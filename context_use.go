package ctrctl

type ContextUseOpts struct {

}

// Set the current docker context.
func ContextUse(opts *ContextUseOpts, context string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "context", "use" },
		[]string{ context },
		opts,
		-1,
	)
}

package ctrctl

import "fmt"

type ContextRmOpts struct {
	// Force the removal of a context in use.
	Force bool
}

// Remove one or more contexts.
func ContextRm(opts *ContextRmOpts, context ...string) (
	stdout string, stderr string, err error) {
	if len(context) == 0 {
		return "", "", fmt.Errorf("context must have at least one value")
	}
	return runCtrCmd(
		[]string{ "context", "rm" },
		context,
		opts,
		-1,
	)
}

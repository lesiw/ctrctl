package ctrctl

import "fmt"

type ContextInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more contexts.
func ContextInspect(opts *ContextInspectOpts, context ...string) (
	stdout string, stderr string, err error) {
	if len(context) == 0 {
		return "", "", fmt.Errorf("context must have at least one value")
	}
	return runCtrCmd(
		[]string{ "context", "inspect" },
		context,
		opts,
		0,
	)
}

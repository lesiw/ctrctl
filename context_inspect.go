package ctrctl

type ContextInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more contexts.
func ContextInspect(opts *ContextInspectOpts, context string, extraContext ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "context", "inspect" },
		append([]string{ context }, extraContext...),
		opts,
		0,
	)
}

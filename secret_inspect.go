package ctrctl

type SecretInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more secrets.
func SecretInspect(opts *SecretInspectOpts, secret string, extraSecret ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "secret", "inspect" },
		append([]string{ secret }, extraSecret...),
		opts,
		0,
	)
}

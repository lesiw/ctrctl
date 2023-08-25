package ctrctl

import "fmt"

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
func SecretInspect(opts *SecretInspectOpts, secret ...string) (
	stdout string, stderr string, err error) {
	if len(secret) == 0 {
		return "", "", fmt.Errorf("secret must have at least one value")
	}
	return runCtrCmd(
		[]string{ "secret", "inspect" },
		secret,
		opts,
		0,
	)
}

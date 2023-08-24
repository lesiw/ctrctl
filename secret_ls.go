package ctrctl

type SecretLsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Only display IDs.
	Quiet string
}

// List secrets.
func SecretLs(opts *SecretLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "secret", "ls" },
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type VersionOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Show the Docker version information.
func Version(opts *VersionOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"version"},
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type SystemInfoOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display system-wide information.
func SystemInfo(opts *SystemInfoOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "info"},
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type ConfigInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty string
}

// Display detailed information on one or more configs.
func ConfigInspect(opts *ConfigInspectOpts, config string, extraConfig ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "config", "inspect" },
		append([]string{ config }, extraConfig...),
		opts,
		0,
	)
}

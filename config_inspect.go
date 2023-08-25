package ctrctl

import "fmt"

type ConfigInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more configs.
func ConfigInspect(opts *ConfigInspectOpts, config ...string) (
	stdout string, stderr string, err error) {
	if len(config) == 0 {
		return "", "", fmt.Errorf("config must have at least one value")
	}
	return runCtrCmd(
		[]string{ "config", "inspect" },
		config,
		opts,
		0,
	)
}

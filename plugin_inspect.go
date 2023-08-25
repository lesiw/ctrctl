package ctrctl

type PluginInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more plugins.
func PluginInspect(opts *PluginInspectOpts, plugin ...string) (
	stdout string, stderr string, err error) {
	if len(plugin) == 0 {
		return "", "", fmt.Errorf("plugin must have at least one value")
	}
	return runCtrCmd(
		[]string{"plugin", "inspect"},
		plugin,
		opts,
		0,
	)
}

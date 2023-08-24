package ctrctl

type PluginInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more plugins.
func PluginInspect(opts *PluginInspectOpts, plugin string, extraPlugin ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "inspect" },
		append([]string{ plugin }, extraPlugin...),
		opts,
		0,
	)
}

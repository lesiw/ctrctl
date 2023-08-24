package ctrctl

type PluginRmOpts struct {
	// Force the removal of an active plugin.
	Force bool
}

// Remove one or more plugins.
func PluginRm(opts *PluginRmOpts, plugin string, extraPlugin ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "rm" },
		append([]string{ plugin }, extraPlugin...),
		opts,
		0,
	)
}

package ctrctl

type PluginEnableOpts struct {
	// HTTP client timeout (in seconds).
	Timeout string
}

// Enable a plugin.
func PluginEnable(opts *PluginEnableOpts, plugin string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "enable" },
		[]string{ plugin },
		opts,
		0,
	)
}

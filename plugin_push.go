package ctrctl

type PluginPushOpts struct {
	// Skip image signing.
	DisableContentTrust bool
}

// Push a plugin to a registry.
func PluginPush(opts *PluginPushOpts, pluginTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"plugin", "push"},
		[]string{pluginTag},
		opts,
		0,
	)
}

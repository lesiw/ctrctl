package ctrctl

type PluginDisableOpts struct {
	// Force the disable of an active plugin.
	Force bool
}

// Disable a plugin.
func PluginDisable(opts *PluginDisableOpts, plugin string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"plugin", "disable"},
		[]string{plugin},
		opts,
		0,
	)
}

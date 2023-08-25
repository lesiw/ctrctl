package ctrctl

import "fmt"

type PluginRmOpts struct {
	// Force the removal of an active plugin.
	Force bool
}

// Remove one or more plugins.
func PluginRm(opts *PluginRmOpts, plugin ...string) (
	stdout string, stderr string, err error) {
	if len(plugin) == 0 {
		return "", "", fmt.Errorf("plugin must have at least one value")
	}
	return runCtrCmd(
		[]string{ "plugin", "rm" },
		plugin,
		opts,
		0,
	)
}

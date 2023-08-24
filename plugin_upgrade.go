package ctrctl

type PluginUpgradeOpts struct {
	// Skip image verification.
	DisableContentTrust bool

	// Grant all permissions necessary to run the plugin.
	GrantAllPermissions bool

	// Do not check if specified remote plugin matches existing plugin image.
	SkipRemoteCheck bool
}

// Upgrade an existing plugin.
func PluginUpgrade(opts *PluginUpgradeOpts, plugin string, remote string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "upgrade" },
		[]string{ plugin, remote },
		opts,
		0,
	)
}

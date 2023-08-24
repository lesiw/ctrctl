package ctrctl

type PluginInstallOpts struct {
	// Local name for plugin.
	Alias string

	// Do not enable the plugin on install.
	Disable bool

	// Skip image verification.
	DisableContentTrust bool

	// Grant all permissions necessary to run the plugin.
	GrantAllPermissions bool
}

// Install a plugin.
func PluginInstall(opts *PluginInstallOpts, plugin string, keyValue ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "install" },
		append([]string{ plugin }, keyValue...),
		opts,
		0,
	)
}

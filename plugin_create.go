package ctrctl

type PluginCreateOpts struct {
	// Compress the context using gzip.
	Compress string
}

// Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory.
func PluginCreate(opts *PluginCreateOpts, plugin string, pluginDataDir string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "create" },
		[]string{ plugin, pluginDataDir },
		opts,
		0,
	)
}

package ctrctl

type PluginOpts struct {

}

// Manage plugins.
func Plugin(opts *PluginOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin" },
		[]string{},
		opts,
		-1,
	)
}

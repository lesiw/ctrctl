package ctrctl

type PluginSetOpts struct {

}

// Change settings for a plugin.
func PluginSet(opts *PluginSetOpts, plugin string, keyValue string, extraKeyValue ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "set" },
		append([]string{ plugin,keyValue }, extraKeyValue...),
		opts,
		-1,
	)
}

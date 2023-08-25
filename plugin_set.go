package ctrctl

import "fmt"

type PluginSetOpts struct {

}

// Change settings for a plugin.
func PluginSet(opts *PluginSetOpts, plugin string, keyValue ...string) (
	stdout string, stderr string, err error) {
	if len(keyValue) == 0 {
		return "", "", fmt.Errorf("keyValue must have at least one value")
	}
	return runCtrCmd(
		[]string{ "plugin", "set" },
		append([]string{ plugin }, keyValue...),
		opts,
		-1,
	)
}

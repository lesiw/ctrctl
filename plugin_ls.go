package ctrctl

type PluginLsOpts struct {
	// Provide filter values (e.g. `enabled=true`).
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Don't truncate output.
	NoTrunc bool

	// Only display plugin IDs.
	Quiet bool
}

// List plugins.
func PluginLs(opts *PluginLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "plugin", "ls" },
		[]string{},
		opts,
		0,
	)
}

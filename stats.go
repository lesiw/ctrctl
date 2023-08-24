package ctrctl

type StatsOpts struct {
	// Show all containers (default shows just running).
	All string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Disable streaming stats and only pull the first result.
	NoStream string

	// Do not truncate output.
	NoTrunc string
}

// Display a live stream of container(s) resource usage statistics.
func Stats(opts *StatsOpts, container ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stats" },
		[]string{},
		opts,
		0,
	)
}

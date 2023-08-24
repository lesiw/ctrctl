package ctrctl

type HistoryOpts struct {
	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print sizes and dates in human readable format.
	Human bool

	// Don't truncate output.
	NoTrunc bool

	// Only show image IDs.
	Quiet bool
}

// Show the history of an image.
func History(opts *HistoryOpts, image string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "history" },
		[]string{ image },
		opts,
		0,
	)
}

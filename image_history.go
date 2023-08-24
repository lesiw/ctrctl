package ctrctl

type ImageHistoryOpts struct {
	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print sizes and dates in human readable format.
	Human string

	// Don't truncate output.
	NoTrunc string

	// Only show image IDs.
	Quiet string
}

// Show the history of an image.
func ImageHistory(opts *ImageHistoryOpts, image string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "history" },
		[]string{ image },
		opts,
		0,
	)
}

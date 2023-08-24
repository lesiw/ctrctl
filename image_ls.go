package ctrctl

type ImageLsOpts struct {
	// Show all images (default hides intermediate images).
	All bool

	// Show digests.
	Digests bool

	// Filter output based on conditions provided.
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

	// Only show image IDs.
	Quiet bool
}

// List images.
func ImageLs(opts *ImageLsOpts, repositoryTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "ls" },
		[]string{ repositoryTag },
		opts,
		0,
	)
}

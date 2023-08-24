package ctrctl

type SearchOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print search using a Go template.
	Format string

	// Max number of search results.
	Limit *int

	// Don't truncate output.
	NoTrunc bool
}

// Search Docker Hub for images.
func Search(opts *SearchOpts, term string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "search" },
		[]string{ term },
		opts,
		0,
	)
}

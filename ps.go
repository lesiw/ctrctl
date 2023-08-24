package ctrctl

type PsOpts struct {
	// Show all containers (default shows just running).
	All string

	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Show n last created containers (includes all states).
	Last string

	// Show the latest created container (includes all states).
	Latest string

	// Don't truncate output.
	NoTrunc string

	// Only display container IDs.
	Quiet string

	// Display total file sizes.
	Size string
}

// List containers.
func Ps(opts *PsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "ps" },
		[]string{},
		opts,
		0,
	)
}

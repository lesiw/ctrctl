package ctrctl

type ContainerLsOpts struct {
	// Show all containers (default shows just running).
	All bool

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
	Last *int

	// Show the latest created container (includes all states).
	Latest bool

	// Don't truncate output.
	NoTrunc bool

	// Only display container IDs.
	Quiet bool

	// Display total file sizes.
	Size bool
}

// List containers.
func ContainerLs(opts *ContainerLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "ls"},
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type VolumeLsOpts struct {
	// Display only cluster volumes, and use cluster volume list formatting.
	Cluster bool

	// Provide filter values (e.g. `dangling=true`).
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Only display volume names.
	Quiet bool
}

// List volumes.
func VolumeLs(opts *VolumeLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "ls"},
		[]string{},
		opts,
		0,
	)
}

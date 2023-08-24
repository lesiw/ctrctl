package ctrctl

type ServicePsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Do not map IDs to Names.
	NoResolve string

	// Do not truncate output.
	NoTrunc string

	// Only display task IDs.
	Quiet string
}

// List the tasks of one or more services.
func ServicePs(opts *ServicePsOpts, service string, extraService ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "ps" },
		append([]string{ service }, extraService...),
		opts,
		0,
	)
}

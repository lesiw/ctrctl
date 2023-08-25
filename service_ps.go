package ctrctl

import "fmt"

type ServicePsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Only display task IDs.
	Quiet bool
}

// List the tasks of one or more services.
func ServicePs(opts *ServicePsOpts, service ...string) (
	stdout string, stderr string, err error) {
	if len(service) == 0 {
		return "", "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{ "service", "ps" },
		service,
		opts,
		0,
	)
}

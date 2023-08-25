package ctrctl

import "fmt"

type ServiceInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more services.
func ServiceInspect(opts *ServiceInspectOpts, service ...string) (
	stdout string, stderr string, err error) {
	if len(service) == 0 {
		return "", "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{ "service", "inspect" },
		service,
		opts,
		0,
	)
}

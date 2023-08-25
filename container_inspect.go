package ctrctl

type ContainerInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Display total file sizes.
	Size bool
}

// Display detailed information on one or more containers.
func ContainerInspect(opts *ContainerInspectOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "inspect"},
		container,
		opts,
		0,
	)
}

package ctrctl

type InspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Display total file sizes if the type is container.
	Size bool

	// Return JSON for specified type.
	Type string
}

// Return low-level information on Docker objects.
func Inspect(opts *InspectOpts, nameId ...string) (
	stdout string, stderr string, err error) {
	if len(nameId) == 0 {
		return "", "", fmt.Errorf("nameId must have at least one value")
	}
	return runCtrCmd(
		[]string{"inspect"},
		nameId,
		opts,
		0,
	)
}

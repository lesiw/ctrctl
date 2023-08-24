package ctrctl

type InspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Display total file sizes if the type is container.
	Size string

	// Return JSON for specified type.
	Type string
}

// Return low-level information on Docker objects.
func Inspect(opts *InspectOpts, nameId string, extraNameId ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "inspect" },
		append([]string{ nameId }, extraNameId...),
		opts,
		0,
	)
}

package ctrctl

type VolumeInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more volumes.
func VolumeInspect(opts *VolumeInspectOpts, volume string, extraVolume ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "volume", "inspect" },
		append([]string{ volume }, extraVolume...),
		opts,
		0,
	)
}

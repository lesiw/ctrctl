package ctrctl

type ImageInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more images.
func ImageInspect(opts *ImageInspectOpts, image string, extraImage ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "inspect" },
		append([]string{ image }, extraImage...),
		opts,
		0,
	)
}

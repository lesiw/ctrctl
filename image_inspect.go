package ctrctl

import "fmt"

type ImageInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display detailed information on one or more images.
func ImageInspect(opts *ImageInspectOpts, image ...string) (
	stdout string, stderr string, err error) {
	if len(image) == 0 {
		return "", "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{ "image", "inspect" },
		image,
		opts,
		0,
	)
}

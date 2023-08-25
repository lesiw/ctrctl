package ctrctl

type ImageTagOpts struct {
}

// Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
func ImageTag(opts *ImageTagOpts, sourceImageTag string, targetImageTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"image", "tag"},
		[]string{sourceImageTag, targetImageTag},
		opts,
		-1,
	)
}

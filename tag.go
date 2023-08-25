package ctrctl

type TagOpts struct {
}

// Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
func Tag(opts *TagOpts, sourceImageTag string, targetImageTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"tag"},
		[]string{sourceImageTag, targetImageTag},
		opts,
		-1,
	)
}

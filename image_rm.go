package ctrctl

type ImageRmOpts struct {
	// Force removal of the image.
	Force string

	// Do not delete untagged parents.
	NoPrune string
}

// Remove one or more images.
func ImageRm(opts *ImageRmOpts, image string, extraImage ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "rm" },
		append([]string{ image }, extraImage...),
		opts,
		0,
	)
}

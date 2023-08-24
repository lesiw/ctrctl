package ctrctl

type ImageRmOpts struct {
	// Force removal of the image.
	Force bool

	// Do not delete untagged parents.
	NoPrune bool
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

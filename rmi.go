package ctrctl

type RmiOpts struct {
	// Force removal of the image.
	Force string

	// Do not delete untagged parents.
	NoPrune string
}

// Remove one or more images.
func Rmi(opts *RmiOpts, image string, extraImage ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "rmi" },
		append([]string{ image }, extraImage...),
		opts,
		0,
	)
}

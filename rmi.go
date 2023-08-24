package ctrctl

type RmiOpts struct {
	// Force removal of the image.
	Force bool

	// Do not delete untagged parents.
	NoPrune bool
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

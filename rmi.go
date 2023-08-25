package ctrctl

type RmiOpts struct {
	// Force removal of the image.
	Force bool

	// Do not delete untagged parents.
	NoPrune bool
}

// Remove one or more images.
func Rmi(opts *RmiOpts, image ...string) (
	stdout string, stderr string, err error) {
	if len(image) == 0 {
		return "", "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"rmi"},
		image,
		opts,
		0,
	)
}

package ctrctl

type ImageOpts struct {
}

// Manage images.
func Image(opts *ImageOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"image"},
		[]string{},
		opts,
		-1,
	)
}

package ctrctl

type ImageLoadOpts struct {
	// Read from tar archive file, instead of STDIN.
	Input string

	// Suppress the load output.
	Quiet bool
}

// Load an image from a tar archive or STDIN.
func ImageLoad(opts *ImageLoadOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"image", "load"},
		[]string{},
		opts,
		0,
	)
}

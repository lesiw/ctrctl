package ctrctl

type LoadOpts struct {
	// Read from tar archive file, instead of STDIN.
	Input string

	// Suppress the load output.
	Quiet bool
}

// Load an image from a tar archive or STDIN.
func Load(opts *LoadOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "load" },
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type ExportOpts struct {
	// Write to a file, instead of STDOUT.
	Output string
}

// Export a container's filesystem as a tar archive.
func Export(opts *ExportOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"export"},
		[]string{container},
		opts,
		0,
	)
}

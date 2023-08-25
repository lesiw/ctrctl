package ctrctl

type ContainerExportOpts struct {
	// Write to a file, instead of STDOUT.
	Output string
}

// Export a container's filesystem as a tar archive.
func ContainerExport(opts *ContainerExportOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "export"},
		[]string{container},
		opts,
		0,
	)
}

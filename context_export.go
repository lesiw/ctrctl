package ctrctl

type ContextExportOpts struct {
	// Export as a kubeconfig file.
	Kubeconfig bool
}

// Export a context to a tar archive FILE or a tar stream on STDOUT.
func ContextExport(opts *ContextExportOpts, context string, file string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"context", "export"},
		[]string{context, file},
		opts,
		0,
	)
}

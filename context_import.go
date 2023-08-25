package ctrctl

type ContextImportOpts struct {
}

// Import a context from a tar or zip file.
func ContextImport(opts *ContextImportOpts, context string, file string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"context", "import"},
		[]string{context, file},
		opts,
		-1,
	)
}

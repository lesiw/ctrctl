package ctrctl

type ImageImportOpts struct {
	// Apply Dockerfile instruction to the created image.
	Change string

	// Set commit message for imported image.
	Message string

	// Set platform if server is multi-platform capable.
	Platform string
}

// Import the contents from a tarball to create a filesystem image.
func ImageImport(opts *ImageImportOpts, fileUrl string, RepositoryTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"image", "import"},
		[]string{fileUrl, RepositoryTag},
		opts,
		0,
	)
}

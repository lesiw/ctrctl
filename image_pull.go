package ctrctl

type ImagePullOpts struct {
	// Download all tagged images in the repository.
	AllTags string

	// Skip image verification.
	DisableContentTrust string

	// Set platform if server is multi-platform capable.
	Platform string

	// Suppress verbose output.
	Quiet string
}

// Download an image from a registry.
func ImagePull(opts *ImagePullOpts, nameTagDigest string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "pull" },
		[]string{ nameTagDigest },
		opts,
		0,
	)
}

package ctrctl

type ImagePushOpts struct {
	// Push all tags of an image to the repository.
	AllTags bool

	// Skip image signing.
	DisableContentTrust bool

	// Suppress verbose output.
	Quiet bool
}

// Upload an image to a registry.
func ImagePush(opts *ImagePushOpts, nameTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"image", "push"},
		[]string{nameTag},
		opts,
		0,
	)
}

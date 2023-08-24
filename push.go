package ctrctl

type PushOpts struct {
	// Push all tags of an image to the repository.
	AllTags bool

	// Skip image signing.
	DisableContentTrust bool

	// Suppress verbose output.
	Quiet bool
}

// Upload an image to a registry.
func Push(opts *PushOpts, nameTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "push" },
		[]string{ nameTag },
		opts,
		0,
	)
}

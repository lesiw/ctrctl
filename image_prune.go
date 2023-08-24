package ctrctl

type ImagePruneOpts struct {
	// Remove all unused images, not just dangling ones.
	All string

	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force string
}

// Remove unused images.
func ImagePrune(opts *ImagePruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "prune" },
		[]string{},
		opts,
		0,
	)
}

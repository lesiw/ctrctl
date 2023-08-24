package ctrctl

type ManifestCreateOpts struct {
	// Amend an existing manifest list.
	Amend string

	// Allow communication with an insecure registry.
	Insecure string
}

// Create a local manifest list for annotating and pushing to a registry.
func ManifestCreate(opts *ManifestCreateOpts, manifestList string, manifest string, extraManifest ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest", "create" },
		append([]string{ manifestList,manifest }, extraManifest...),
		opts,
		-1,
	)
}

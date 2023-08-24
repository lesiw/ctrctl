package ctrctl

type ManifestInspectOpts struct {
	// Allow communication with an insecure registry.
	Insecure string

	// Output additional info including layers and platform.
	Verbose string
}

// Display an image manifest, or manifest list.
func ManifestInspect(opts *ManifestInspectOpts, manifestList string, manifest string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest", "inspect" },
		[]string{ manifestList, manifest },
		opts,
		0,
	)
}

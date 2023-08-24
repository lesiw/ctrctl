package ctrctl

type ManifestAnnotateOpts struct {
	// Set architecture.
	Arch string

	// Set operating system.
	Os string

	// Set operating system feature.
	OsFeatures string

	// Set operating system version.
	OsVersion string

	// Set architecture variant.
	Variant string
}

// Add additional information to a local image manifest.
func ManifestAnnotate(opts *ManifestAnnotateOpts, manifestList string, manifest string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest", "annotate" },
		[]string{ manifestList, manifest },
		opts,
		0,
	)
}

package ctrctl

type ManifestPushOpts struct {
	// Allow push to an insecure registry.
	Insecure string

	// Remove the local manifest list after push.
	Purge string
}

// Push a manifest list to a repository.
func ManifestPush(opts *ManifestPushOpts, manifestList string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest", "push" },
		[]string{ manifestList },
		opts,
		0,
	)
}

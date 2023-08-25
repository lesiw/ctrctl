package ctrctl

import "fmt"

type ManifestCreateOpts struct {
	// Amend an existing manifest list.
	Amend bool

	// Allow communication with an insecure registry.
	Insecure bool
}

// Create a local manifest list for annotating and pushing to a registry.
func ManifestCreate(opts *ManifestCreateOpts, manifestList string, manifest ...string) (
	stdout string, stderr string, err error) {
	if len(manifest) == 0 {
		return "", "", fmt.Errorf("manifest must have at least one value")
	}
	return runCtrCmd(
		[]string{ "manifest", "create" },
		append([]string{ manifestList }, manifest...),
		opts,
		-1,
	)
}

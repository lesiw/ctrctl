package ctrctl

import "fmt"

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
		[]string{"manifest", "annotate"},
		[]string{manifestList, manifest},
		opts,
		0,
	)
}

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
		[]string{"manifest", "create"},
		append([]string{manifestList}, manifest...),
		opts,
		-1,
	)
}

type ManifestInspectOpts struct {
	// Allow communication with an insecure registry.
	Insecure bool

	// Output additional info including layers and platform.
	Verbose bool
}

// Display an image manifest, or manifest list.
func ManifestInspect(opts *ManifestInspectOpts, manifestList string, manifest string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"manifest", "inspect"},
		[]string{manifestList, manifest},
		opts,
		0,
	)
}

type ManifestPushOpts struct {
	// Allow push to an insecure registry.
	Insecure bool

	// Remove the local manifest list after push.
	Purge bool
}

// Push a manifest list to a repository.
func ManifestPush(opts *ManifestPushOpts, manifestList string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"manifest", "push"},
		[]string{manifestList},
		opts,
		0,
	)
}

type ManifestRmOpts struct {
}

// Delete one or more manifest lists from local storage.
func ManifestRm(opts *ManifestRmOpts, manifestList ...string) (
	stdout string, stderr string, err error) {
	if len(manifestList) == 0 {
		return "", "", fmt.Errorf("manifestList must have at least one value")
	}
	return runCtrCmd(
		[]string{"manifest", "rm"},
		manifestList,
		opts,
		-1,
	)
}

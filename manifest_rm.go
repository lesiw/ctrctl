package ctrctl

type ManifestRmOpts struct {

}

// Delete one or more manifest lists from local storage.
func ManifestRm(opts *ManifestRmOpts, manifestList string, extraManifestList ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest", "rm" },
		append([]string{ manifestList }, extraManifestList...),
		opts,
		-1,
	)
}

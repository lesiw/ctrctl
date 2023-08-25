package ctrctl

import "fmt"

type ManifestRmOpts struct {

}

// Delete one or more manifest lists from local storage.
func ManifestRm(opts *ManifestRmOpts, manifestList ...string) (
	stdout string, stderr string, err error) {
	if len(manifestList) == 0 {
		return "", "", fmt.Errorf("manifestList must have at least one value")
	}
	return runCtrCmd(
		[]string{ "manifest", "rm" },
		manifestList,
		opts,
		-1,
	)
}

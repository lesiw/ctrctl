package ctrctl

import "fmt"

type TrustSignerRemoveOpts struct {
	// Do not prompt for confirmation before removing the most recent signer.
	Force bool
}

// Remove a signer.
func TrustSignerRemove(opts *TrustSignerRemoveOpts, name string, repository ...string) (
	stdout string, stderr string, err error) {
	if len(repository) == 0 {
		return "", "", fmt.Errorf("repository must have at least one value")
	}
	return runCtrCmd(
		[]string{ "trust", "signer", "remove" },
		append([]string{ name }, repository...),
		opts,
		0,
	)
}

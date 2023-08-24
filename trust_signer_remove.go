package ctrctl

type TrustSignerRemoveOpts struct {
	// Do not prompt for confirmation before removing the most recent signer.
	Force string
}

// Remove a signer.
func TrustSignerRemove(opts *TrustSignerRemoveOpts, name string, repository string, extraRepository ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "signer", "remove" },
		append([]string{ name,repository }, extraRepository...),
		opts,
		0,
	)
}

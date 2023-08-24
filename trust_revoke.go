package ctrctl

type TrustRevokeOpts struct {
	// Do not prompt for confirmation.
	Yes string
}

// Remove trust for an image.
func TrustRevoke(opts *TrustRevokeOpts, imageTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "revoke" },
		[]string{ imageTag },
		opts,
		0,
	)
}

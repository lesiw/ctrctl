package ctrctl

type TrustSignerOpts struct {

}

// Manage entities who can sign Docker images.
func TrustSigner(opts *TrustSignerOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "signer" },
		[]string{},
		opts,
		-1,
	)
}

package ctrctl

type TrustSignerAddOpts struct {
	// Path to the signer's public key file.
	Key string
}

// Add a signer.
func TrustSignerAdd(opts *TrustSignerAddOpts, name string, repository string, extraRepository ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "signer", "add" },
		append([]string{ name,repository }, extraRepository...),
		opts,
		0,
	)
}

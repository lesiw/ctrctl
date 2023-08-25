package ctrctl

type TrustSignerAddOpts struct {
	// Path to the signer's public key file.
	Key string
}

// Add a signer.
func TrustSignerAdd(opts *TrustSignerAddOpts, name string, repository ...string) (
	stdout string, stderr string, err error) {
	if len(repository) == 0 {
		return "", "", fmt.Errorf("repository must have at least one value")
	}
	return runCtrCmd(
		[]string{"trust", "signer", "add"},
		append([]string{name}, repository...),
		opts,
		0,
	)
}

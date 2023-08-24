package ctrctl

type TrustKeyGenerateOpts struct {
	// Directory to generate key in, defaults to current directory.
	Dir string
}

// Generate and load a signing key-pair.
func TrustKeyGenerate(opts *TrustKeyGenerateOpts, name string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "key", "generate" },
		[]string{ name },
		opts,
		-1,
	)
}

package ctrctl

type TrustKeyLoadOpts struct {
	// Name for the loaded key.
	Name string
}

// Load a private key file for signing.
func TrustKeyLoad(opts *TrustKeyLoadOpts, keyfile string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"trust", "key", "load"},
		[]string{keyfile},
		opts,
		0,
	)
}

package ctrctl

type TrustKeyOpts struct {
}

// Manage keys for signing Docker images.
func TrustKey(opts *TrustKeyOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"trust", "key"},
		[]string{},
		opts,
		-1,
	)
}

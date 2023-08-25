package ctrctl

type TrustOpts struct {
}

// Manage trust on Docker images.
func Trust(opts *TrustOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"trust"},
		[]string{},
		opts,
		-1,
	)
}

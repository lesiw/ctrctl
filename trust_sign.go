package ctrctl

type TrustSignOpts struct {
	// Sign a locally tagged image.
	Local bool
}

// Sign an image.
func TrustSign(opts *TrustSignOpts, imageTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"trust", "sign"},
		[]string{imageTag},
		opts,
		-1,
	)
}

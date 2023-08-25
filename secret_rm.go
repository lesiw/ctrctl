package ctrctl

type SecretRmOpts struct {
}

// Remove one or more secrets.
func SecretRm(opts *SecretRmOpts, secret ...string) (
	stdout string, stderr string, err error) {
	if len(secret) == 0 {
		return "", "", fmt.Errorf("secret must have at least one value")
	}
	return runCtrCmd(
		[]string{"secret", "rm"},
		secret,
		opts,
		-1,
	)
}

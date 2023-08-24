package ctrctl

type SecretRmOpts struct {

}

// Remove one or more secrets.
func SecretRm(opts *SecretRmOpts, secret string, extraSecret ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "secret", "rm" },
		append([]string{ secret }, extraSecret...),
		opts,
		-1,
	)
}

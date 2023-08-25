package ctrctl

type SecretCreateOpts struct {
	// Secret driver.
	Driver string

	// Secret labels.
	Label string

	// Template driver.
	TemplateDriver string
}

// Create a secret from a file or STDIN as content.
func SecretCreate(opts *SecretCreateOpts, secret string, file string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"secret", "create"},
		[]string{secret, file},
		opts,
		0,
	)
}

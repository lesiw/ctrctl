package ctrctl

type SystemDialStdioOpts struct {
}

// Proxy the stdio stream to the daemon connection. Should not be invoked manually.
func SystemDialStdio(opts *SystemDialStdioOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "dial-stdio"},
		[]string{},
		opts,
		-1,
	)
}

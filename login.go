package ctrctl

type LoginOpts struct {
	// Password.
	Password string

	// Take the password from stdin.
	PasswordStdin bool

	// Username.
	Username string
}

// Log in to a registry.
func Login(opts *LoginOpts, server string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "login" },
		[]string{ server },
		opts,
		0,
	)
}

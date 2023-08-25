package ctrctl

type LogoutOpts struct {
}

// Log out from a registry.
func Logout(opts *LogoutOpts, server string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"logout"},
		[]string{server},
		opts,
		-1,
	)
}

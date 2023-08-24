package ctrctl

type SecretOpts struct {

}

// Manage Swarm secrets.
func Secret(opts *SecretOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "secret" },
		[]string{},
		opts,
		-1,
	)
}

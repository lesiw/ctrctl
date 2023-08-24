package ctrctl

type ConfigOpts struct {

}

// Manage Swarm configs.
func Config(opts *ConfigOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "config" },
		[]string{},
		opts,
		-1,
	)
}

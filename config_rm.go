package ctrctl

type ConfigRmOpts struct {
}

// Remove one or more configs.
func ConfigRm(opts *ConfigRmOpts, config ...string) (
	stdout string, stderr string, err error) {
	if len(config) == 0 {
		return "", "", fmt.Errorf("config must have at least one value")
	}
	return runCtrCmd(
		[]string{"config", "rm"},
		config,
		opts,
		-1,
	)
}

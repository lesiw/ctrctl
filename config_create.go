package ctrctl

type ConfigCreateOpts struct {
	// Config labels.
	Label string

	// Template driver.
	TemplateDriver string
}

// Create a config from a file or STDIN.
func ConfigCreate(opts *ConfigCreateOpts, config string, file string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"config", "create"},
		[]string{config, file},
		opts,
		0,
	)
}

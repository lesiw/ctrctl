package ctrctl

type StackConfigOpts struct {
	// Path to a Compose file, or `-` to read from stdin.
	ComposeFile string

	// Skip interpolation and output only merged config.
	SkipInterpolation bool
}

// Outputs the final config file, after doing merges and interpolations.
func StackConfig(opts *StackConfigOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack", "config" },
		[]string{},
		opts,
		0,
	)
}

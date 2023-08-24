package ctrctl

type ConfigRmOpts struct {

}

// Remove one or more configs.
func ConfigRm(opts *ConfigRmOpts, config string, extraConfig ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "config", "rm" },
		append([]string{ config }, extraConfig...),
		opts,
		-1,
	)
}

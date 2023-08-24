package ctrctl

type VolumeOpts struct {

}

// Manage volumes.
func Volume(opts *VolumeOpts, command string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "volume" },
		[]string{ command },
		opts,
		-1,
	)
}

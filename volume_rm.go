package ctrctl

type VolumeRmOpts struct {
	// Force the removal of one or more volumes.
	Force bool
}

// Remove one or more volumes.
func VolumeRm(opts *VolumeRmOpts, volume string, extraVolume ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "volume", "rm" },
		append([]string{ volume }, extraVolume...),
		opts,
		0,
	)
}

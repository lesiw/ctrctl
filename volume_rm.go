package ctrctl

type VolumeRmOpts struct {
	// Force the removal of one or more volumes.
	Force bool
}

// Remove one or more volumes.
func VolumeRm(opts *VolumeRmOpts, volume ...string) (
	stdout string, stderr string, err error) {
	if len(volume) == 0 {
		return "", "", fmt.Errorf("volume must have at least one value")
	}
	return runCtrCmd(
		[]string{"volume", "rm"},
		volume,
		opts,
		0,
	)
}

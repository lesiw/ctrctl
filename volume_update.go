package ctrctl

type VolumeUpdateOpts struct {
	// Cluster Volume availability (`active`, `pause`, `drain`).
	Availability string
}

// Update a volume (cluster volumes only).
func VolumeUpdate(opts *VolumeUpdateOpts, volume string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "volume", "update" },
		[]string{ volume },
		opts,
		0,
	)
}

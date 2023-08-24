package ctrctl

type VolumePruneOpts struct {
	// Remove all unused volumes, not just anonymous ones.
	All string

	// Provide filter values (e.g. `label=<label>`).
	Filter string

	// Do not prompt for confirmation.
	Force string
}

// Remove all unused local volumes.
func VolumePrune(opts *VolumePruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "volume", "prune" },
		[]string{},
		opts,
		0,
	)
}

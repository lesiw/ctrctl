package ctrctl

type VolumePruneOpts struct {
	// Remove all unused volumes, not just anonymous ones.
	All bool

	// Provide filter values (e.g. `label=<label>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool
}

// Remove all unused local volumes.
func VolumePrune(opts *VolumePruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "prune"},
		[]string{},
		opts,
		0,
	)
}

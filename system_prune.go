package ctrctl

type SystemPruneOpts struct {
	// Remove all unused images not just dangling ones.
	All bool

	// Provide filter values (e.g. `label=<key>=<value>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Prune volumes.
	Volumes bool
}

// Remove unused data.
func SystemPrune(opts *SystemPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "prune"},
		[]string{},
		opts,
		0,
	)
}

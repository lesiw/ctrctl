package ctrctl

type SystemPruneOpts struct {
	// Remove all unused images not just dangling ones.
	All string

	// Provide filter values (e.g. `label=<key>=<value>`).
	Filter string

	// Do not prompt for confirmation.
	Force string

	// Prune volumes.
	Volumes string
}

// Remove unused data.
func SystemPrune(opts *SystemPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "system", "prune" },
		[]string{},
		opts,
		0,
	)
}

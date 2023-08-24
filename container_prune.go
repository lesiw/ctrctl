package ctrctl

type ContainerPruneOpts struct {
	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool
}

// Remove all stopped containers.
func ContainerPrune(opts *ContainerPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "prune" },
		[]string{},
		opts,
		0,
	)
}

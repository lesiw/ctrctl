package ctrctl

type NetworkPruneOpts struct {
	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force string
}

// Remove all unused networks.
func NetworkPrune(opts *NetworkPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "network", "prune" },
		[]string{},
		opts,
		0,
	)
}

package ctrctl

type BuilderPruneOpts struct {
	// Remove all unused build cache, not just dangling ones.
	All string

	// Provide filter values (e.g. `until=24h`).
	Filter string

	// Do not prompt for confirmation.
	Force string

	// Amount of disk space to keep for cache.
	KeepStorage string
}

// Remove build cache.
func BuilderPrune(opts *BuilderPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "builder", "prune" },
		[]string{},
		opts,
		-1,
	)
}

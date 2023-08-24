package ctrctl

type StackDeployOpts struct {
	// Path to a Compose file, or `-` to read from stdin.
	ComposeFile string

	// Prune services that are no longer referenced.
	Prune bool

	// Query the registry to resolve image digest and supported platforms (`always`, `changed`, `never`).
	ResolveImage string

	// Send registry authentication details to Swarm agents.
	WithRegistryAuth bool
}

// Deploy a new stack or update an existing stack.
func StackDeploy(opts *StackDeployOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack", "deploy" },
		[]string{ stack },
		opts,
		0,
	)
}

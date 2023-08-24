package ctrctl

type StackOpts struct {
	// Orchestrator to use (swarm|all).
	Orchestrator string
}

// Manage Swarm stacks.
func Stack(opts *StackOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack" },
		[]string{},
		opts,
		0,
	)
}

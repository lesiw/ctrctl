package ctrctl

type ContextUpdateOpts struct {
	// Default orchestrator for stack operations to use with this context (swarm|kubernetes|all).
	DefaultStackOrchestrator string

	// Description of the context.
	Description string

	// set the docker endpoint.
	Docker string

	// set the kubernetes endpoint.
	Kubernetes string
}

// Update a context.
func ContextUpdate(opts *ContextUpdateOpts, context string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "context", "update" },
		[]string{ context },
		opts,
		0,
	)
}

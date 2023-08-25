package ctrctl

type ContextCreateOpts struct {
	// Default orchestrator for stack operations to use with this context (`swarm`, `kubernetes`, `all`).
	DefaultStackOrchestrator string

	// Description of the context.
	Description string

	// set the docker endpoint.
	Docker string

	// create context from a named context.
	From string

	// set the kubernetes endpoint.
	Kubernetes string
}

// Create a context.
func ContextCreate(opts *ContextCreateOpts, context string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"context", "create"},
		[]string{context},
		opts,
		0,
	)
}

package ctrctl

type ServiceRollbackOpts struct {
	// Exit immediately instead of waiting for the service to converge.
	Detach string

	// Suppress progress output.
	Quiet string
}

// Revert changes to a service's configuration.
func ServiceRollback(opts *ServiceRollbackOpts, service string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "rollback" },
		[]string{ service },
		opts,
		0,
	)
}

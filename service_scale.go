package ctrctl

type ServiceScaleOpts struct {
	// Exit immediately instead of waiting for the service to converge.
	Detach bool
}

// Scale one or multiple replicated services.
func ServiceScale(opts *ServiceScaleOpts, serviceReplicas ...string) (
	stdout string, stderr string, err error) {
	if len(serviceReplicas) == 0 {
		return "", "", fmt.Errorf("serviceReplicas must have at least one value")
	}
	return runCtrCmd(
		[]string{"service", "scale"},
		serviceReplicas,
		opts,
		-1,
	)
}

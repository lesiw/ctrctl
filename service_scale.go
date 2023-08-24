package ctrctl

type ServiceScaleOpts struct {
	// Exit immediately instead of waiting for the service to converge.
	Detach string
}

// Scale one or multiple replicated services.
func ServiceScale(opts *ServiceScaleOpts, serviceReplicas string, extraServiceReplicas ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "scale" },
		append([]string{ serviceReplicas }, extraServiceReplicas...),
		opts,
		-1,
	)
}

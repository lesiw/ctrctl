package ctrctl

type ServiceOpts struct {
}

// Manage Swarm services.
func Service(opts *ServiceOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"service"},
		[]string{},
		opts,
		-1,
	)
}

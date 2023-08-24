package ctrctl

type ServiceRmOpts struct {

}

// Remove one or more services.
func ServiceRm(opts *ServiceRmOpts, service string, extraService ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "rm" },
		append([]string{ service }, extraService...),
		opts,
		-1,
	)
}

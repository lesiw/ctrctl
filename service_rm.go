package ctrctl

import "fmt"

type ServiceRmOpts struct {

}

// Remove one or more services.
func ServiceRm(opts *ServiceRmOpts, service ...string) (
	stdout string, stderr string, err error) {
	if len(service) == 0 {
		return "", "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{ "service", "rm" },
		service,
		opts,
		-1,
	)
}

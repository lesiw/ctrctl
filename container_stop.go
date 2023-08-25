package ctrctl

import "fmt"

type ContainerStopOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Stop one or more running containers.
func ContainerStop(opts *ContainerStopOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{ "container", "stop" },
		container,
		opts,
		0,
	)
}

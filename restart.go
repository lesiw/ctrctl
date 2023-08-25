package ctrctl

import "fmt"

type RestartOpts struct {
	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Restart one or more containers.
func Restart(opts *RestartOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{ "restart" },
		container,
		opts,
		0,
	)
}

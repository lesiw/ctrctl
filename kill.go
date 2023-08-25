package ctrctl

import "fmt"

type KillOpts struct {
	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func Kill(opts *KillOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{ "kill" },
		container,
		opts,
		0,
	)
}

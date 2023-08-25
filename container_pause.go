package ctrctl

import "fmt"

type ContainerPauseOpts struct {

}

// Pause all processes within one or more containers.
func ContainerPause(opts *ContainerPauseOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{ "container", "pause" },
		container,
		opts,
		-1,
	)
}

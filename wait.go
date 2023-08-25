package ctrctl

import "fmt"

type WaitOpts struct {

}

// Block until one or more containers stop, then print their exit codes.
func Wait(opts *WaitOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{ "wait" },
		container,
		opts,
		-1,
	)
}

package ctrctl

type ContainerWaitOpts struct {

}

// Block until one or more containers stop, then print their exit codes.
func ContainerWait(opts *ContainerWaitOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "wait" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

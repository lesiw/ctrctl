package ctrctl

type ContainerUnpauseOpts struct {

}

// Unpause all processes within one or more containers.
func ContainerUnpause(opts *ContainerUnpauseOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "unpause" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

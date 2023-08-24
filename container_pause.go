package ctrctl

type ContainerPauseOpts struct {

}

// Pause all processes within one or more containers.
func ContainerPause(opts *ContainerPauseOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "pause" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

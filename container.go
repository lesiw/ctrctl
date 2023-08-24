package ctrctl

type ContainerOpts struct {

}

// Manage containers.
func Container(opts *ContainerOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container" },
		[]string{},
		opts,
		-1,
	)
}

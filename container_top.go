package ctrctl

type ContainerTopOpts struct {
}

// Display the running processes of a container.
func ContainerTop(opts *ContainerTopOpts, container string, psOptions string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "top"},
		[]string{container, psOptions},
		opts,
		-1,
	)
}

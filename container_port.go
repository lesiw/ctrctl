package ctrctl

type ContainerPortOpts struct {
}

// List port mappings or a specific mapping for the container.
func ContainerPort(opts *ContainerPortOpts, container string, privatePortProto string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "port"},
		[]string{container, privatePortProto},
		opts,
		-1,
	)
}

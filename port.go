package ctrctl

type PortOpts struct {
}

// List port mappings or a specific mapping for the container.
func Port(opts *PortOpts, container string, privatePortProto string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"port"},
		[]string{container, privatePortProto},
		opts,
		-1,
	)
}

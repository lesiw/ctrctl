package ctrctl

type ContainerAttachOpts struct {
	// Override the key sequence for detaching a container.
	DetachKeys string

	// Do not attach STDIN.
	NoStdin bool

	// Proxy all received signals to the process.
	SigProxy bool
}

// Attach local standard input, output, and error streams to a running container.
func ContainerAttach(opts *ContainerAttachOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "attach" },
		[]string{ container },
		opts,
		0,
	)
}

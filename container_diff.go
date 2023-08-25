package ctrctl

type ContainerDiffOpts struct {
}

// Inspect changes to files or directories on a container's filesystem.
func ContainerDiff(opts *ContainerDiffOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "diff"},
		[]string{container},
		opts,
		-1,
	)
}

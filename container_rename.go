package ctrctl

type ContainerRenameOpts struct {
}

// Rename a container.
func ContainerRename(opts *ContainerRenameOpts, container string, newName string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"container", "rename"},
		[]string{container, newName},
		opts,
		-1,
	)
}

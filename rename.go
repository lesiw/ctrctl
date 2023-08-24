package ctrctl

type RenameOpts struct {

}

// Rename a container.
func Rename(opts *RenameOpts, container string, newName string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "rename" },
		[]string{ container, newName },
		opts,
		-1,
	)
}

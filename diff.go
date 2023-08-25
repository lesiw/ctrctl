package ctrctl

type DiffOpts struct {
}

// Inspect changes to files or directories on a container's filesystem.
func Diff(opts *DiffOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"diff"},
		[]string{container},
		opts,
		-1,
	)
}

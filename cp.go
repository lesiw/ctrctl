package ctrctl

type CpOpts struct {
	// Archive mode (copy all uid/gid information).
	Archive string

	// Always follow symbol link in SRC_PATH.
	FollowLink string

	// Suppress progress output during copy. Progress output is automatically suppressed if no terminal is attached.
	Quiet string
}

// Copy files/folders between a container and the local filesystem.
func Cp(opts *CpOpts, srcpath string, dstpath string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "cp" },
		[]string{ srcpath, dstpath },
		opts,
		-1,
	)
}

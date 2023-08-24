package ctrctl

type WaitOpts struct {

}

// Block until one or more containers stop, then print their exit codes.
func Wait(opts *WaitOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "wait" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

package ctrctl

type UnpauseOpts struct {

}

// Unpause all processes within one or more containers.
func Unpause(opts *UnpauseOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "unpause" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

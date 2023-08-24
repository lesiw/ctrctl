package ctrctl

type PauseOpts struct {

}

// Pause all processes within one or more containers.
func Pause(opts *PauseOpts, container string, extraContainer ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "pause" },
		append([]string{ container }, extraContainer...),
		opts,
		-1,
	)
}

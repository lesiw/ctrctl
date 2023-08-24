package ctrctl

type BuilderOpts struct {

}

// Manage builds.
func Builder(opts *BuilderOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "builder" },
		[]string{},
		opts,
		-1,
	)
}

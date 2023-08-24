package ctrctl

type StackRmOpts struct {

}

// Remove one or more stacks.
func StackRm(opts *StackRmOpts, stack string, extraStack ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "stack", "rm" },
		append([]string{ stack }, extraStack...),
		opts,
		0,
	)
}

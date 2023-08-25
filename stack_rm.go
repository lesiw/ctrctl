package ctrctl

type StackRmOpts struct {
}

// Remove one or more stacks.
func StackRm(opts *StackRmOpts, stack ...string) (
	stdout string, stderr string, err error) {
	if len(stack) == 0 {
		return "", "", fmt.Errorf("stack must have at least one value")
	}
	return runCtrCmd(
		[]string{"stack", "rm"},
		stack,
		opts,
		0,
	)
}

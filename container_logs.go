package ctrctl

type ContainerLogsOpts struct {
	// Show extra details provided to logs.
	Details string

	// Follow log output.
	Follow string

	// Show logs since timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Since string

	// Number of lines to show from the end of the logs.
	Tail string

	// Show timestamps.
	Timestamps string

	// Show logs before a timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Until string
}

// Fetch the logs of a container.
func ContainerLogs(opts *ContainerLogsOpts, container string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "logs" },
		[]string{ container },
		opts,
		0,
	)
}

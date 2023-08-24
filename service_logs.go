package ctrctl

type ServiceLogsOpts struct {
	// Show extra details provided to logs.
	Details string

	// Follow log output.
	Follow string

	// Do not map IDs to Names in output.
	NoResolve string

	// Do not include task IDs in output.
	NoTaskIds string

	// Do not truncate output.
	NoTrunc string

	// Do not neatly format logs.
	Raw string

	// Show logs since timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Since string

	// Number of lines to show from the end of the logs.
	Tail string

	// Show timestamps.
	Timestamps string
}

// Fetch the logs of a service or task.
func ServiceLogs(opts *ServiceLogsOpts, serviceTask string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "logs" },
		[]string{ serviceTask },
		opts,
		0,
	)
}

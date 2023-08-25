package ctrctl

type SystemDfOpts struct {
	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Show detailed information on space usage.
	Verbose bool
}

// Show docker disk usage.
func SystemDf(opts *SystemDfOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "df"},
		[]string{},
		opts,
		0,
	)
}

type SystemDialStdioOpts struct {
}

// Proxy the stdio stream to the daemon connection. Should not be invoked manually.
func SystemDialStdio(opts *SystemDialStdioOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "dial-stdio"},
		[]string{},
		opts,
		-1,
	)
}

type SystemEventsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Format the output using the given Go template.
	Format string

	// Show all events created since timestamp.
	Since string

	// Stream events until this timestamp.
	Until string
}

// Get real time events from the server.
func SystemEvents(opts *SystemEventsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "events"},
		[]string{},
		opts,
		0,
	)
}

type SystemInfoOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string
}

// Display system-wide information.
func SystemInfo(opts *SystemInfoOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "info"},
		[]string{},
		opts,
		0,
	)
}

type SystemPruneOpts struct {
	// Remove all unused images not just dangling ones.
	All bool

	// Provide filter values (e.g. `label=<key>=<value>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Prune volumes.
	Volumes bool
}

// Remove unused data.
func SystemPrune(opts *SystemPruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"system", "prune"},
		[]string{},
		opts,
		0,
	)
}

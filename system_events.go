package ctrctl

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

package ctrctl

type SwarmJoinTokenOpts struct {
	// Only display token.
	Quiet string

	// Rotate join token.
	Rotate string
}

// Manage join tokens.
func SwarmJoinToken(opts *SwarmJoinTokenOpts, WorkerManager string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "swarm", "join-token" },
		[]string{ WorkerManager },
		opts,
		0,
	)
}

package ctrctl

type ContainerCommitOpts struct {
	// Author (e.g., `John Hannibal Smith <hannibal@a-team.com>`).
	Author string

	// Apply Dockerfile instruction to the created image.
	Change string

	// Commit message.
	Message string

	// Pause container during commit.
	Pause bool
}

// Create a new image from a container's changes.
func ContainerCommit(opts *ContainerCommitOpts, container string, repositoryTag string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "commit" },
		[]string{ container, repositoryTag },
		opts,
		0,
	)
}

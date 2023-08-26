package ctrctl

import "os/exec"

type CheckpointCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool

	// Leave the container running after checkpoint.
	LeaveRunning bool
}

// Create a checkpoint from a running container.
func CheckpointCreate(opts *CheckpointCreateOpts, container string, checkpoint string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "create"},
		[]string{container, checkpoint},
		opts,
		0,
	)
}

type CheckpointLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool
}

// List checkpoints for a container.
func CheckpointLs(opts *CheckpointLsOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "ls"},
		[]string{container},
		opts,
		0,
	)
}

type CheckpointRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool
}

// Remove a checkpoint.
func CheckpointRm(opts *CheckpointRmOpts, container string, checkpoint string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "rm"},
		[]string{container, checkpoint},
		opts,
		0,
	)
}

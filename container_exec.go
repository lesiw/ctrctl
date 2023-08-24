package ctrctl

type ContainerExecOpts struct {
	// Detached mode: run command in the background.
	Detach string

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

	// Keep STDIN open even if not attached.
	Interactive string

	// Give extended privileges to the command.
	Privileged string

	// Allocate a pseudo-TTY.
	Tty string

	// Username or UID (format: `<name|uid>[:<group|gid>]`).
	User string

	// Working directory inside the container.
	Workdir string
}

// Execute a command in a running container.
func ContainerExec(opts *ContainerExecOpts, container string, command string, arg ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "exec" },
		append([]string{ container,command }, arg...),
		opts,
		0,
	)
}

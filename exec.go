package ctrctl

type ExecOpts struct {
	// Detached mode: run command in the background.
	Detach bool

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

	// Keep STDIN open even if not attached.
	Interactive bool

	// Give extended privileges to the command.
	Privileged bool

	// Allocate a pseudo-TTY.
	Tty bool

	// Username or UID (format: `<name|uid>[:<group|gid>]`).
	User string

	// Working directory inside the container.
	Workdir string
}

// Execute a command in a running container.
func Exec(opts *ExecOpts, container string, command string, arg ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "exec" },
		append([]string{ container,command }, arg...),
		opts,
		0,
	)
}

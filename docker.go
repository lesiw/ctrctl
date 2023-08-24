package ctrctl

type DockerOpts struct {
	// Location of client config files.
	Config string

	// Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with `docker context use`).
	Context string

	// Enable debug mode.
	Debug bool

	// Print usage.
	Help bool

	// Daemon socket to connect to.
	Host string

	// Set the logging level (`debug`, `info`, `warn`, `error`, `fatal`).
	LogLevel string

	// Use TLS; implied by --tlsverify.
	Tls bool

	// Trust certs signed only by this CA.
	Tlscacert string

	// Path to TLS certificate file.
	Tlscert string

	// Path to TLS key file.
	Tlskey string

	// Use TLS and verify the remote.
	Tlsverify bool
}

// The base command for the Docker CLI.
func Docker(opts *DockerOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{  },
		[]string{},
		opts,
		-1,
	)
}

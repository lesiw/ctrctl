package ctrctl

type ServiceCreateOpts struct {
	// Add Linux capabilities.
	CapAdd string

	// Drop Linux capabilities.
	CapDrop string

	// Specify configurations to expose to the service.
	Config string

	// Placement constraints.
	Constraint string

	// Container labels.
	ContainerLabel string

	// Credential spec for managed service account (Windows only).
	CredentialSpec string

	// Exit immediately instead of waiting for the service to converge.
	Detach bool

	// Set custom DNS servers.
	Dns string

	// Set DNS options.
	DnsOption string

	// Set custom DNS search domains.
	DnsSearch string

	// Endpoint mode (vip or dnsrr).
	EndpointMode string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

	// User defined resources.
	GenericResource string

	// Set one or more supplementary user groups for the container.
	Group string

	// Command to run to check health.
	HealthCmd string

	// Time between running the check (ms|s|m|h).
	HealthInterval string

	// Consecutive failures needed to report unhealthy.
	HealthRetries *int

	// Start period for the container to initialize before counting retries towards unstable (ms|s|m|h).
	HealthStartPeriod string

	// Maximum time to allow one check to run (ms|s|m|h).
	HealthTimeout string

	// Set one or more custom host-to-IP mappings (host:ip).
	Host string

	// Container hostname.
	Hostname string

	// Use an init inside each service container to forward signals and reap processes.
	Init bool

	// Service container isolation mode.
	Isolation string

	// Service labels.
	Label string

	// Limit CPUs.
	LimitCpu string

	// Limit Memory.
	LimitMemory string

	// Limit maximum number of processes (default 0 = unlimited).
	LimitPids string

	// Logging driver for service.
	LogDriver string

	// Logging driver options.
	LogOpt string

	// Number of job tasks to run concurrently (default equal to --replicas).
	MaxConcurrent string

	// Service mode (`replicated`, `global`, `replicated-job`, `global-job`).
	Mode string

	// Attach a filesystem mount to the service.
	Mount string

	// Service name.
	Name string

	// Network attachments.
	Network string

	// Disable any container-specified HEALTHCHECK.
	NoHealthcheck bool

	// Do not query the registry to resolve image digest and supported platforms.
	NoResolveImage bool

	// Add a placement preference.
	PlacementPref string

	// Publish a port as a node port.
	Publish string

	// Suppress progress output.
	Quiet bool

	// Mount the container's root filesystem as read only.
	ReadOnly bool

	// Number of tasks.
	Replicas string

	// Maximum number of tasks per node (default 0 = unlimited).
	ReplicasMaxPerNode string

	// Reserve CPUs.
	ReserveCpu string

	// Reserve Memory.
	ReserveMemory string

	// Restart when condition is met (`none`, `on-failure`, `any`) (default `any`).
	RestartCondition string

	// Delay between restart attempts (ns|us|ms|s|m|h) (default 5s).
	RestartDelay string

	// Maximum number of restarts before giving up.
	RestartMaxAttempts string

	// Window used to evaluate the restart policy (ns|us|ms|s|m|h).
	RestartWindow string

	// Delay between task rollbacks (ns|us|ms|s|m|h) (default 0s).
	RollbackDelay string

	// Action on rollback failure (`pause`, `continue`) (default `pause`).
	RollbackFailureAction string

	// Failure rate to tolerate during a rollback (default 0).
	RollbackMaxFailureRatio string

	// Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h) (default 5s).
	RollbackMonitor string

	// Rollback order (`start-first`, `stop-first`) (default `stop-first`).
	RollbackOrder string

	// Maximum number of tasks rolled back simultaneously (0 to roll back all at once).
	RollbackParallelism string

	// Specify secrets to expose to the service.
	Secret string

	// Time to wait before force killing a container (ns|us|ms|s|m|h) (default 10s).
	StopGracePeriod string

	// Signal to stop the container.
	StopSignal string

	// Sysctl options.
	Sysctl string

	// Allocate a pseudo-TTY.
	Tty bool

	// Ulimit options.
	Ulimit string

	// Delay between updates (ns|us|ms|s|m|h) (default 0s).
	UpdateDelay string

	// Action on update failure (`pause`, `continue`, `rollback`) (default `pause`).
	UpdateFailureAction string

	// Failure rate to tolerate during an update (default 0).
	UpdateMaxFailureRatio string

	// Duration after each task update to monitor for failure (ns|us|ms|s|m|h) (default 5s).
	UpdateMonitor string

	// Update order (`start-first`, `stop-first`) (default `stop-first`).
	UpdateOrder string

	// Maximum number of tasks updated simultaneously (0 to update all at once).
	UpdateParallelism string

	// Username or UID (format: <name|uid>[:<group|gid>]).
	User string

	// Send registry authentication details to swarm agents.
	WithRegistryAuth bool

	// Working directory inside the container.
	Workdir string
}

// Create a new service.
func ServiceCreate(opts *ServiceCreateOpts, image string, command string, arg ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "service", "create" },
		append([]string{ image,command }, arg...),
		opts,
		0,
	)
}

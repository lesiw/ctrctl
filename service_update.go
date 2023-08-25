package ctrctl

type ServiceUpdateOpts struct {
	// Service command args.
	Args string

	// Add Linux capabilities.
	CapAdd string

	// Drop Linux capabilities.
	CapDrop string

	// Add or update a config file on a service.
	ConfigAdd string

	// Remove a configuration file.
	ConfigRm string

	// Add or update a placement constraint.
	ConstraintAdd string

	// Remove a constraint.
	ConstraintRm string

	// Add or update a container label.
	ContainerLabelAdd string

	// Remove a container label by its key.
	ContainerLabelRm string

	// Credential spec for managed service account (Windows only).
	CredentialSpec string

	// Exit immediately instead of waiting for the service to converge.
	Detach bool

	// Add or update a custom DNS server.
	DnsAdd string

	// Add or update a DNS option.
	DnsOptionAdd string

	// Remove a DNS option.
	DnsOptionRm string

	// Remove a custom DNS server.
	DnsRm string

	// Add or update a custom DNS search domain.
	DnsSearchAdd string

	// Remove a DNS search domain.
	DnsSearchRm string

	// Endpoint mode (vip or dnsrr).
	EndpointMode string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Add or update an environment variable.
	EnvAdd string

	// Remove an environment variable.
	EnvRm string

	// Force update even if no changes require it.
	Force bool

	// Add a Generic resource.
	GenericResourceAdd string

	// Remove a Generic resource.
	GenericResourceRm string

	// Add an additional supplementary user group to the container.
	GroupAdd string

	// Remove a previously added supplementary user group from the container.
	GroupRm string

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

	// Add a custom host-to-IP mapping (`host:ip`).
	HostAdd string

	// Remove a custom host-to-IP mapping (`host:ip`).
	HostRm string

	// Container hostname.
	Hostname string

	// Service image tag.
	Image string

	// Use an init inside each service container to forward signals and reap processes.
	Init bool

	// Service container isolation mode.
	Isolation string

	// Add or update a service label.
	LabelAdd string

	// Remove a label by its key.
	LabelRm string

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

	// Add or update a mount on a service.
	MountAdd string

	// Remove a mount by its target path.
	MountRm string

	// Add a network.
	NetworkAdd string

	// Remove a network.
	NetworkRm string

	// Disable any container-specified HEALTHCHECK.
	NoHealthcheck bool

	// Do not query the registry to resolve image digest and supported platforms.
	NoResolveImage bool

	// Add a placement preference.
	PlacementPrefAdd string

	// Remove a placement preference.
	PlacementPrefRm string

	// Add or update a published port.
	PublishAdd string

	// Remove a published port by its target port.
	PublishRm string

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

	// Restart when condition is met (`none`, `on-failure`, `any`).
	RestartCondition string

	// Delay between restart attempts (ns|us|ms|s|m|h).
	RestartDelay string

	// Maximum number of restarts before giving up.
	RestartMaxAttempts string

	// Window used to evaluate the restart policy (ns|us|ms|s|m|h).
	RestartWindow string

	// Rollback to previous specification.
	Rollback bool

	// Delay between task rollbacks (ns|us|ms|s|m|h).
	RollbackDelay string

	// Action on rollback failure (`pause`, `continue`).
	RollbackFailureAction string

	// Failure rate to tolerate during a rollback.
	RollbackMaxFailureRatio string

	// Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h).
	RollbackMonitor string

	// Rollback order (`start-first`, `stop-first`).
	RollbackOrder string

	// Maximum number of tasks rolled back simultaneously (0 to roll back all at once).
	RollbackParallelism string

	// Add or update a secret on a service.
	SecretAdd string

	// Remove a secret.
	SecretRm string

	// Time to wait before force killing a container (ns|us|ms|s|m|h).
	StopGracePeriod string

	// Signal to stop the container.
	StopSignal string

	// Add or update a Sysctl option.
	SysctlAdd string

	// Remove a Sysctl option.
	SysctlRm string

	// Allocate a pseudo-TTY.
	Tty bool

	// Add or update a ulimit option.
	UlimitAdd string

	// Remove a ulimit option.
	UlimitRm string

	// Delay between updates (ns|us|ms|s|m|h).
	UpdateDelay string

	// Action on update failure (`pause`, `continue`, `rollback`).
	UpdateFailureAction string

	// Failure rate to tolerate during an update.
	UpdateMaxFailureRatio string

	// Duration after each task update to monitor for failure (ns|us|ms|s|m|h).
	UpdateMonitor string

	// Update order (`start-first`, `stop-first`).
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

// Update a service.
func ServiceUpdate(opts *ServiceUpdateOpts, service string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"service", "update"},
		[]string{service},
		opts,
		0,
	)
}

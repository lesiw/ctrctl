package ctrctl

import (
	"fmt"
	"os/exec"
)

type ServiceCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

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

	// Print usage.
	Help bool

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
func ServiceCreate(opts *ServiceCreateOpts, image string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"service", "create"},
		append([]string{image, command}, arg...),
		opts,
		0,
	)
}

type ServiceInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more services.
func ServiceInspect(opts *ServiceInspectOpts, service ...string) (string, error) {
	if len(service) == 0 {
		return "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{"service", "inspect"},
		service,
		opts,
		0,
	)
}

type ServiceLogsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show extra details provided to logs.
	Details bool

	// Follow log output.
	Follow bool

	// Print usage.
	Help bool

	// Do not map IDs to Names in output.
	NoResolve bool

	// Do not include task IDs in output.
	NoTaskIds bool

	// Do not truncate output.
	NoTrunc bool

	// Do not neatly format logs.
	Raw bool

	// Show logs since timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Since string

	// Number of lines to show from the end of the logs.
	Tail string

	// Show timestamps.
	Timestamps bool
}

// Fetch the logs of a service or task.
func ServiceLogs(opts *ServiceLogsOpts, serviceTask string) (string, error) {
	return runCtrCmd(
		[]string{"service", "logs"},
		[]string{serviceTask},
		opts,
		0,
	)
}

type ServiceLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Only display IDs.
	Quiet bool
}

// List services.
func ServiceLs(opts *ServiceLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"service", "ls"},
		[]string{},
		opts,
		0,
	)
}

type ServicePsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Print usage.
	Help bool

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Only display task IDs.
	Quiet bool
}

// List the tasks of one or more services.
func ServicePs(opts *ServicePsOpts, service ...string) (string, error) {
	if len(service) == 0 {
		return "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{"service", "ps"},
		service,
		opts,
		0,
	)
}

type ServiceRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Remove one or more services.
func ServiceRm(opts *ServiceRmOpts, service ...string) (string, error) {
	if len(service) == 0 {
		return "", fmt.Errorf("service must have at least one value")
	}
	return runCtrCmd(
		[]string{"service", "rm"},
		service,
		opts,
		-1,
	)
}

type ServiceRollbackOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Exit immediately instead of waiting for the service to converge.
	Detach bool

	// Print usage.
	Help bool

	// Suppress progress output.
	Quiet bool
}

// Revert changes to a service's configuration.
func ServiceRollback(opts *ServiceRollbackOpts, service string) (string, error) {
	return runCtrCmd(
		[]string{"service", "rollback"},
		[]string{service},
		opts,
		0,
	)
}

type ServiceScaleOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Exit immediately instead of waiting for the service to converge.
	Detach bool

	// Print usage.
	Help bool
}

// Scale one or multiple replicated services.
func ServiceScale(opts *ServiceScaleOpts, serviceReplicas ...string) (string, error) {
	if len(serviceReplicas) == 0 {
		return "", fmt.Errorf("serviceReplicas must have at least one value")
	}
	return runCtrCmd(
		[]string{"service", "scale"},
		serviceReplicas,
		opts,
		-1,
	)
}

type ServiceUpdateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

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

	// Print usage.
	Help bool

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
func ServiceUpdate(opts *ServiceUpdateOpts, service string) (string, error) {
	return runCtrCmd(
		[]string{"service", "update"},
		[]string{service},
		opts,
		0,
	)
}

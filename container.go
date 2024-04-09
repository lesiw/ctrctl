package ctrctl

import (
	"fmt"
	"os/exec"
)

type ContainerAttachOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Print usage.
	Help bool

	// Do not attach STDIN.
	NoStdin bool

	// Proxy all received signals to the process.
	SigProxy bool
}

// Attach local standard input, output, and error streams to a running container.
func ContainerAttach(opts *ContainerAttachOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"container", "attach"},
		[]string{container},
		opts,
		0,
	)
}

type ContainerCommitOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Author (e.g., `John Hannibal Smith <hannibal@a-team.com>`).
	Author string

	// Apply Dockerfile instruction to the created image.
	Change []string

	// Print usage.
	Help bool

	// Commit message.
	Message string

	// Pause container during commit.
	Pause bool
}

// Create a new image from a container's changes.
func ContainerCommit(opts *ContainerCommitOpts, container string, repositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"container", "commit"},
		[]string{container, repositoryTag},
		opts,
		0,
	)
}

type ContainerCpOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Archive mode (copy all uid/gid information).
	Archive bool

	// Always follow symbol link in SRC_PATH.
	FollowLink bool

	// Print usage.
	Help bool

	// Suppress progress output during copy. Progress output is automatically suppressed if no terminal is attached.
	Quiet bool
}

// Copy files/folders between a container and the local filesystem.
func ContainerCp(opts *ContainerCpOpts, srcpath string, dstpath string) (string, error) {
	return runCtrCmd(
		[]string{"container", "cp"},
		[]string{srcpath, dstpath},
		opts,
		-1,
	)
}

type ContainerCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Add a custom host-to-IP mapping (host:ip).
	AddHost []string

	// Add an annotation to the container (passed through to the OCI runtime).
	Annotation string

	// Attach to STDIN, STDOUT or STDERR.
	Attach []string

	// Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0).
	BlkioWeight string

	// Block IO weight (relative device weight).
	BlkioWeightDevice []string

	// Add Linux capabilities.
	CapAdd []string

	// Drop Linux capabilities.
	CapDrop []string

	// Optional parent cgroup for the container.
	CgroupParent string

	// Cgroup namespace to use (host|private).
	// 'host':    Run the container in the Docker host's cgroup namespace.
	// 'private': Run the container in its own private cgroup namespace.
	// '':        Use the cgroup namespace as configured by the.
	// default-cgroupns-mode option on the daemon (default).
	Cgroupns string

	// Write the container ID to the file.
	Cidfile string

	// CPU count (Windows only).
	CpuCount string

	// CPU percent (Windows only).
	CpuPercent string

	// Limit CPU CFS (Completely Fair Scheduler) period.
	CpuPeriod string

	// Limit CPU CFS (Completely Fair Scheduler) quota.
	CpuQuota string

	// Limit CPU real-time period in microseconds.
	CpuRtPeriod string

	// Limit CPU real-time runtime in microseconds.
	CpuRtRuntime string

	// CPU shares (relative weight).
	CpuShares string

	// Number of CPUs.
	Cpus string

	// CPUs in which to allow execution (0-3, 0,1).
	CpusetCpus string

	// MEMs in which to allow execution (0-3, 0,1).
	CpusetMems string

	// Add a host device to the container.
	Device []string

	// Add a rule to the cgroup allowed devices list.
	DeviceCgroupRule []string

	// Limit read rate (bytes per second) from a device.
	DeviceReadBps []string

	// Limit read rate (IO per second) from a device.
	DeviceReadIops []string

	// Limit write rate (bytes per second) to a device.
	DeviceWriteBps []string

	// Limit write rate (IO per second) to a device.
	DeviceWriteIops []string

	// Skip image verification.
	DisableContentTrust bool

	// Set custom DNS servers.
	Dns []string

	// Set DNS options.
	DnsOpt []string

	// Set DNS options.
	DnsOption []string

	// Set custom DNS search domains.
	DnsSearch []string

	// Container NIS domain name.
	Domainname string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Set environment variables.
	Env []string

	// Read in a file of environment variables.
	EnvFile []string

	// Expose a port or a range of ports.
	Expose []string

	// GPU devices to add to the container ('all' to pass all GPUs).
	Gpus string

	// Add additional groups to join.
	GroupAdd []string

	// Command to run to check health.
	HealthCmd string

	// Time between running the check (ms|s|m|h) (default 0s).
	HealthInterval string

	// Consecutive failures needed to report unhealthy.
	HealthRetries *int

	// Time between running the check during the start period (ms|s|m|h) (default 0s).
	HealthStartInterval string

	// Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s).
	HealthStartPeriod string

	// Maximum time to allow one check to run (ms|s|m|h) (default 0s).
	HealthTimeout string

	// Print usage.
	Help bool

	// Container host name.
	Hostname string

	// Run an init inside the container that forwards signals and reaps processes.
	Init bool

	// Keep STDIN open even if not attached.
	Interactive bool

	// Maximum IO bandwidth limit for the system drive (Windows only).
	IoMaxbandwidth string

	// Maximum IOps limit for the system drive (Windows only).
	IoMaxiops string

	// IPv4 address (e.g., 172.30.100.104).
	Ip string

	// IPv6 address (e.g., 2001:db8::33).
	Ip6 string

	// IPC mode to use.
	Ipc string

	// Container isolation technology.
	Isolation string

	// Kernel memory limit.
	KernelMemory string

	// Set meta data on a container.
	Label []string

	// Read in a line delimited file of labels.
	LabelFile []string

	// Add link to another container.
	Link []string

	// Container IPv4/IPv6 link-local addresses.
	LinkLocalIp []string

	// Logging driver for the container.
	LogDriver string

	// Log driver options.
	LogOpt []string

	// Container MAC address (e.g., 92:d0:c6:0a:29:33).
	MacAddress string

	// Memory limit.
	Memory string

	// Memory soft limit.
	MemoryReservation string

	// Swap limit equal to memory plus swap: '-1' to enable unlimited swap.
	MemorySwap string

	// Tune container memory swappiness (0 to 100).
	MemorySwappiness string

	// Attach a filesystem mount to the container.
	Mount string

	// Assign a name to the container.
	Name string

	// Connect a container to a network.
	Net string

	// Add network-scoped alias for the container.
	NetAlias []string

	// Connect a container to a network.
	Network string

	// Add network-scoped alias for the container.
	NetworkAlias []string

	// Disable any container-specified HEALTHCHECK.
	NoHealthcheck bool

	// Disable OOM Killer.
	OomKillDisable bool

	// Tune host's OOM preferences (-1000 to 1000).
	OomScoreAdj *int

	// PID namespace to use.
	Pid string

	// Tune container pids limit (set -1 for unlimited).
	PidsLimit string

	// Set platform if server is multi-platform capable.
	Platform string

	// Give extended privileges to this container.
	Privileged bool

	// Publish a container's port(s) to the host.
	Publish []string

	// Publish all exposed ports to random ports.
	PublishAll bool

	// Pull image before creating (`always`, `|missing`, `never`).
	Pull string

	// Suppress the pull output.
	Quiet bool

	// Mount the container's root filesystem as read only.
	ReadOnly bool

	// Restart policy to apply when a container exits.
	Restart string

	// Automatically remove the container when it exits.
	Rm bool

	// Runtime to use for this container.
	Runtime string

	// Security Options.
	SecurityOpt []string

	// Size of /dev/shm.
	ShmSize string

	// Signal to stop the container.
	StopSignal string

	// Timeout (in seconds) to stop a container.
	StopTimeout *int

	// Storage driver options for the container.
	StorageOpt []string

	// Sysctl options.
	Sysctl string

	// Mount a tmpfs directory.
	Tmpfs []string

	// Allocate a pseudo-TTY.
	Tty bool

	// Ulimit options.
	Ulimit string

	// Username or UID (format: <name|uid>[:<group|gid>]).
	User string

	// User namespace to use.
	Userns string

	// UTS namespace to use.
	Uts string

	// Bind mount a volume.
	Volume []string

	// Optional volume driver for the container.
	VolumeDriver string

	// Mount volumes from the specified container(s).
	VolumesFrom []string

	// Working directory inside the container.
	Workdir string
}

// Create a new container.
func ContainerCreate(opts *ContainerCreateOpts, image string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"container", "create"},
		append([]string{image, command}, arg...),
		opts,
		0,
	)
}

type ContainerDiffOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Inspect changes to files or directories on a container's filesystem.
func ContainerDiff(opts *ContainerDiffOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"container", "diff"},
		[]string{container},
		opts,
		-1,
	)
}

type ContainerExecOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Detached mode: run command in the background.
	Detach bool

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Set environment variables.
	Env []string

	// Read in a file of environment variables.
	EnvFile []string

	// Print usage.
	Help bool

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
func ContainerExec(opts *ContainerExecOpts, container string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"container", "exec"},
		append([]string{container, command}, arg...),
		opts,
		0,
	)
}

type ContainerExportOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Write to a file, instead of STDOUT.
	Output string
}

// Export a container's filesystem as a tar archive.
func ContainerExport(opts *ContainerExportOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"container", "export"},
		[]string{container},
		opts,
		0,
	)
}

type ContainerInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Display total file sizes.
	Size bool
}

// Display detailed information on one or more containers.
func ContainerInspect(opts *ContainerInspectOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "inspect"},
		container,
		opts,
		0,
	)
}

type ContainerKillOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func ContainerKill(opts *ContainerKillOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "kill"},
		container,
		opts,
		0,
	)
}

type ContainerLogsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show extra details provided to logs.
	Details bool

	// Follow log output.
	Follow bool

	// Print usage.
	Help bool

	// Show logs since timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Since string

	// Number of lines to show from the end of the logs.
	Tail string

	// Show timestamps.
	Timestamps bool

	// Show logs before a timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes).
	Until string
}

// Fetch the logs of a container.
func ContainerLogs(opts *ContainerLogsOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"container", "logs"},
		[]string{container},
		opts,
		0,
	)
}

type ContainerLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show all containers (default shows just running).
	All bool

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

	// Show n last created containers (includes all states).
	Last *int

	// Show the latest created container (includes all states).
	Latest bool

	// Don't truncate output.
	NoTrunc bool

	// Only display container IDs.
	Quiet bool

	// Display total file sizes.
	Size bool
}

// List containers.
func ContainerLs(opts *ContainerLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"container", "ls"},
		[]string{},
		opts,
		0,
	)
}

type ContainerPauseOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Pause all processes within one or more containers.
func ContainerPause(opts *ContainerPauseOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "pause"},
		container,
		opts,
		-1,
	)
}

type ContainerPortOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// List port mappings or a specific mapping for the container.
func ContainerPort(opts *ContainerPortOpts, container string, privatePortProto string) (string, error) {
	return runCtrCmd(
		[]string{"container", "port"},
		[]string{container, privatePortProto},
		opts,
		-1,
	)
}

type ContainerPruneOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Print usage.
	Help bool
}

// Remove all stopped containers.
func ContainerPrune(opts *ContainerPruneOpts) (string, error) {
	return runCtrCmd(
		[]string{"container", "prune"},
		[]string{},
		opts,
		0,
	)
}

type ContainerRenameOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Rename a container.
func ContainerRename(opts *ContainerRenameOpts, container string, newName string) (string, error) {
	return runCtrCmd(
		[]string{"container", "rename"},
		[]string{container, newName},
		opts,
		-1,
	)
}

type ContainerRestartOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Restart one or more containers.
func ContainerRestart(opts *ContainerRestartOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "restart"},
		container,
		opts,
		0,
	)
}

type ContainerRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force the removal of a running container (uses SIGKILL).
	Force bool

	// Print usage.
	Help bool

	// Remove the specified link.
	Link bool

	// Remove anonymous volumes associated with the container.
	Volumes bool
}

// Remove one or more containers.
func ContainerRm(opts *ContainerRmOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "rm"},
		container,
		opts,
		0,
	)
}

type ContainerRunOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Add a custom host-to-IP mapping (host:ip).
	AddHost []string

	// Add an annotation to the container (passed through to the OCI runtime).
	Annotation string

	// Attach to STDIN, STDOUT or STDERR.
	Attach []string

	// Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0).
	BlkioWeight string

	// Block IO weight (relative device weight).
	BlkioWeightDevice []string

	// Add Linux capabilities.
	CapAdd []string

	// Drop Linux capabilities.
	CapDrop []string

	// Optional parent cgroup for the container.
	CgroupParent string

	// Cgroup namespace to use (host|private).
	// 'host':    Run the container in the Docker host's cgroup namespace.
	// 'private': Run the container in its own private cgroup namespace.
	// '':        Use the cgroup namespace as configured by the.
	// default-cgroupns-mode option on the daemon (default).
	Cgroupns string

	// Write the container ID to the file.
	Cidfile string

	// CPU count (Windows only).
	CpuCount string

	// CPU percent (Windows only).
	CpuPercent string

	// Limit CPU CFS (Completely Fair Scheduler) period.
	CpuPeriod string

	// Limit CPU CFS (Completely Fair Scheduler) quota.
	CpuQuota string

	// Limit CPU real-time period in microseconds.
	CpuRtPeriod string

	// Limit CPU real-time runtime in microseconds.
	CpuRtRuntime string

	// CPU shares (relative weight).
	CpuShares string

	// Number of CPUs.
	Cpus string

	// CPUs in which to allow execution (0-3, 0,1).
	CpusetCpus string

	// MEMs in which to allow execution (0-3, 0,1).
	CpusetMems string

	// Run container in background and print container ID.
	Detach bool

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Add a host device to the container.
	Device []string

	// Add a rule to the cgroup allowed devices list.
	DeviceCgroupRule []string

	// Limit read rate (bytes per second) from a device.
	DeviceReadBps []string

	// Limit read rate (IO per second) from a device.
	DeviceReadIops []string

	// Limit write rate (bytes per second) to a device.
	DeviceWriteBps []string

	// Limit write rate (IO per second) to a device.
	DeviceWriteIops []string

	// Skip image verification.
	DisableContentTrust bool

	// Set custom DNS servers.
	Dns []string

	// Set DNS options.
	DnsOpt []string

	// Set DNS options.
	DnsOption []string

	// Set custom DNS search domains.
	DnsSearch []string

	// Container NIS domain name.
	Domainname string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Set environment variables.
	Env []string

	// Read in a file of environment variables.
	EnvFile []string

	// Expose a port or a range of ports.
	Expose []string

	// GPU devices to add to the container ('all' to pass all GPUs).
	Gpus string

	// Add additional groups to join.
	GroupAdd []string

	// Command to run to check health.
	HealthCmd string

	// Time between running the check (ms|s|m|h) (default 0s).
	HealthInterval string

	// Consecutive failures needed to report unhealthy.
	HealthRetries *int

	// Time between running the check during the start period (ms|s|m|h) (default 0s).
	HealthStartInterval string

	// Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s).
	HealthStartPeriod string

	// Maximum time to allow one check to run (ms|s|m|h) (default 0s).
	HealthTimeout string

	// Print usage.
	Help bool

	// Container host name.
	Hostname string

	// Run an init inside the container that forwards signals and reaps processes.
	Init bool

	// Keep STDIN open even if not attached.
	Interactive bool

	// Maximum IO bandwidth limit for the system drive (Windows only).
	IoMaxbandwidth string

	// Maximum IOps limit for the system drive (Windows only).
	IoMaxiops string

	// IPv4 address (e.g., 172.30.100.104).
	Ip string

	// IPv6 address (e.g., 2001:db8::33).
	Ip6 string

	// IPC mode to use.
	Ipc string

	// Container isolation technology.
	Isolation string

	// Kernel memory limit.
	KernelMemory string

	// Set meta data on a container.
	Label []string

	// Read in a line delimited file of labels.
	LabelFile []string

	// Add link to another container.
	Link []string

	// Container IPv4/IPv6 link-local addresses.
	LinkLocalIp []string

	// Logging driver for the container.
	LogDriver string

	// Log driver options.
	LogOpt []string

	// Container MAC address (e.g., 92:d0:c6:0a:29:33).
	MacAddress string

	// Memory limit.
	Memory string

	// Memory soft limit.
	MemoryReservation string

	// Swap limit equal to memory plus swap: '-1' to enable unlimited swap.
	MemorySwap string

	// Tune container memory swappiness (0 to 100).
	MemorySwappiness string

	// Attach a filesystem mount to the container.
	Mount string

	// Assign a name to the container.
	Name string

	// Connect a container to a network.
	Net string

	// Add network-scoped alias for the container.
	NetAlias []string

	// Connect a container to a network.
	Network string

	// Add network-scoped alias for the container.
	NetworkAlias []string

	// Disable any container-specified HEALTHCHECK.
	NoHealthcheck bool

	// Disable OOM Killer.
	OomKillDisable bool

	// Tune host's OOM preferences (-1000 to 1000).
	OomScoreAdj *int

	// PID namespace to use.
	Pid string

	// Tune container pids limit (set -1 for unlimited).
	PidsLimit string

	// Set platform if server is multi-platform capable.
	Platform string

	// Give extended privileges to this container.
	Privileged bool

	// Publish a container's port(s) to the host.
	Publish []string

	// Publish all exposed ports to random ports.
	PublishAll bool

	// Pull image before running (`always`, `missing`, `never`).
	Pull string

	// Suppress the pull output.
	Quiet bool

	// Mount the container's root filesystem as read only.
	ReadOnly bool

	// Restart policy to apply when a container exits.
	Restart string

	// Automatically remove the container when it exits.
	Rm bool

	// Runtime to use for this container.
	Runtime string

	// Security Options.
	SecurityOpt []string

	// Size of /dev/shm.
	ShmSize string

	// Proxy received signals to the process.
	SigProxy bool

	// Signal to stop the container.
	StopSignal string

	// Timeout (in seconds) to stop a container.
	StopTimeout *int

	// Storage driver options for the container.
	StorageOpt []string

	// Sysctl options.
	Sysctl string

	// Mount a tmpfs directory.
	Tmpfs []string

	// Allocate a pseudo-TTY.
	Tty bool

	// Ulimit options.
	Ulimit string

	// Username or UID (format: <name|uid>[:<group|gid>]).
	User string

	// User namespace to use.
	Userns string

	// UTS namespace to use.
	Uts string

	// Bind mount a volume.
	Volume []string

	// Optional volume driver for the container.
	VolumeDriver string

	// Mount volumes from the specified container(s).
	VolumesFrom []string

	// Working directory inside the container.
	Workdir string
}

// Create and run a new container from an image.
func ContainerRun(opts *ContainerRunOpts, image string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"container", "run"},
		append([]string{image, command}, arg...),
		opts,
		0,
	)
}

type ContainerStartOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Attach STDOUT/STDERR and forward signals.
	Attach bool

	// Restore from this checkpoint.
	Checkpoint string

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Print usage.
	Help bool

	// Attach container's STDIN.
	Interactive bool
}

// Start one or more stopped containers.
func ContainerStart(opts *ContainerStartOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "start"},
		container,
		opts,
		0,
	)
}

type ContainerStatsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show all containers (default shows just running).
	All bool

	// Format output using a custom template:
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Disable streaming stats and only pull the first result.
	NoStream bool

	// Do not truncate output.
	NoTrunc bool
}

// Display a live stream of container(s) resource usage statistics.
func ContainerStats(opts *ContainerStatsOpts, container ...string) (string, error) {
	return runCtrCmd(
		[]string{"container", "stats"},
		container,
		opts,
		0,
	)
}

type ContainerStopOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Signal to send to the container.
	Signal string

	// Seconds to wait before killing the container.
	Time *int
}

// Stop one or more running containers.
func ContainerStop(opts *ContainerStopOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "stop"},
		container,
		opts,
		0,
	)
}

type ContainerTopOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Display the running processes of a container.
func ContainerTop(opts *ContainerTopOpts, container string, psOptions string) (string, error) {
	return runCtrCmd(
		[]string{"container", "top"},
		[]string{container, psOptions},
		opts,
		-1,
	)
}

type ContainerUnpauseOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Unpause all processes within one or more containers.
func ContainerUnpause(opts *ContainerUnpauseOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "unpause"},
		container,
		opts,
		-1,
	)
}

type ContainerUpdateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0).
	BlkioWeight string

	// Limit CPU CFS (Completely Fair Scheduler) period.
	CpuPeriod string

	// Limit CPU CFS (Completely Fair Scheduler) quota.
	CpuQuota string

	// Limit the CPU real-time period in microseconds.
	CpuRtPeriod string

	// Limit the CPU real-time runtime in microseconds.
	CpuRtRuntime string

	// CPU shares (relative weight).
	CpuShares string

	// Number of CPUs.
	Cpus string

	// CPUs in which to allow execution (0-3, 0,1).
	CpusetCpus string

	// MEMs in which to allow execution (0-3, 0,1).
	CpusetMems string

	// Print usage.
	Help bool

	// Kernel memory limit (deprecated).
	KernelMemory string

	// Memory limit.
	Memory string

	// Memory soft limit.
	MemoryReservation string

	// Swap limit equal to memory plus swap: -1 to enable unlimited swap.
	MemorySwap string

	// Tune container pids limit (set -1 for unlimited).
	PidsLimit string

	// Restart policy to apply when a container exits.
	Restart string
}

// Update configuration of one or more containers.
func ContainerUpdate(opts *ContainerUpdateOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "update"},
		container,
		opts,
		0,
	)
}

type ContainerWaitOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Block until one or more containers stop, then print their exit codes.
func ContainerWait(opts *ContainerWaitOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "wait"},
		container,
		opts,
		-1,
	)
}

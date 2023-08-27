package ctrctl

import (
	"fmt"
	"os/exec"
)

type DockerOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

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
func Docker(opts *DockerOpts) (string, error) {
	return runCtrCmd(
		[]string{},
		[]string{},
		opts,
		-1,
	)
}

type AttachOpts struct {
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
func Attach(opts *AttachOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"attach"},
		[]string{container},
		opts,
		0,
	)
}

type BuildOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Add a custom host-to-IP mapping (`host:ip`).
	AddHost string

	// Set build-time variables.
	BuildArg string

	// Images to consider as cache sources.
	CacheFrom string

	// Optional parent cgroup for the container.
	CgroupParent string

	// Compress the build context using gzip.
	Compress bool

	// Limit the CPU CFS (Completely Fair Scheduler) period.
	CpuPeriod string

	// Limit the CPU CFS (Completely Fair Scheduler) quota.
	CpuQuota string

	// CPU shares (relative weight).
	CpuShares string

	// CPUs in which to allow execution (0-3, 0,1).
	CpusetCpus string

	// MEMs in which to allow execution (0-3, 0,1).
	CpusetMems string

	// Skip image verification.
	DisableContentTrust bool

	// Name of the Dockerfile (Default is `PATH/Dockerfile`).
	File string

	// Always remove intermediate containers.
	ForceRm bool

	// Print usage.
	Help bool

	// Write the image ID to the file.
	Iidfile string

	// Container isolation technology.
	Isolation string

	// Set metadata for an image.
	Label string

	// Memory limit.
	Memory string

	// Swap limit equal to memory plus swap: -1 to enable unlimited swap.
	MemorySwap string

	// Set the networking mode for the RUN instructions during build.
	Network string

	// Do not use cache when building the image.
	NoCache bool

	// Set platform if server is multi-platform capable.
	Platform string

	// Always attempt to pull a newer version of the image.
	Pull bool

	// Suppress the build output and print image ID on success.
	Quiet bool

	// Remove intermediate containers after a successful build.
	Rm bool

	// Security options.
	SecurityOpt string

	// Size of `/dev/shm`.
	ShmSize string

	// Squash newly built layers into a single new layer.
	Squash bool

	// Name and optionally a tag in the `name:tag` format.
	Tag string

	// Set the target build stage to build.
	Target string

	// Ulimit options.
	Ulimit string
}

// Build an image from a Dockerfile.
func Build(opts *BuildOpts, path string, url string) (string, error) {
	return runCtrCmd(
		[]string{"build"},
		[]string{path, url},
		opts,
		0,
	)
}

type BuilderOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage builds.
func Builder(opts *BuilderOpts) (string, error) {
	return runCtrCmd(
		[]string{"builder"},
		[]string{},
		opts,
		-1,
	)
}

type CheckpointOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage checkpoints.
func Checkpoint(opts *CheckpointOpts) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint"},
		[]string{},
		opts,
		-1,
	)
}

type CommitOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Author (e.g., `John Hannibal Smith <hannibal@a-team.com>`).
	Author string

	// Apply Dockerfile instruction to the created image.
	Change string

	// Print usage.
	Help bool

	// Commit message.
	Message string

	// Pause container during commit.
	Pause bool
}

// Create a new image from a container's changes.
func Commit(opts *CommitOpts, container string, repositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"commit"},
		[]string{container, repositoryTag},
		opts,
		0,
	)
}

type ConfigOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Swarm configs.
func Config(opts *ConfigOpts) (string, error) {
	return runCtrCmd(
		[]string{"config"},
		[]string{},
		opts,
		-1,
	)
}

type ContainerOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage containers.
func Container(opts *ContainerOpts) (string, error) {
	return runCtrCmd(
		[]string{"container"},
		[]string{},
		opts,
		-1,
	)
}

type ContextOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage contexts.
func Context(opts *ContextOpts) (string, error) {
	return runCtrCmd(
		[]string{"context"},
		[]string{},
		opts,
		-1,
	)
}

type CpOpts struct {
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
func Cp(opts *CpOpts, srcpath string, dstpath string) (string, error) {
	return runCtrCmd(
		[]string{"cp"},
		[]string{srcpath, dstpath},
		opts,
		-1,
	)
}

type CreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Add a custom host-to-IP mapping (host:ip).
	AddHost string

	// Add an annotation to the container (passed through to the OCI runtime).
	Annotation string

	// Attach to STDIN, STDOUT or STDERR.
	Attach string

	// Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0).
	BlkioWeight string

	// Block IO weight (relative device weight).
	BlkioWeightDevice string

	// Add Linux capabilities.
	CapAdd string

	// Drop Linux capabilities.
	CapDrop string

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
	Device string

	// Add a rule to the cgroup allowed devices list.
	DeviceCgroupRule string

	// Limit read rate (bytes per second) from a device.
	DeviceReadBps string

	// Limit read rate (IO per second) from a device.
	DeviceReadIops string

	// Limit write rate (bytes per second) to a device.
	DeviceWriteBps string

	// Limit write rate (IO per second) to a device.
	DeviceWriteIops string

	// Skip image verification.
	DisableContentTrust bool

	// Set custom DNS servers.
	Dns string

	// Set DNS options.
	DnsOpt string

	// Set DNS options.
	DnsOption string

	// Set custom DNS search domains.
	DnsSearch string

	// Container NIS domain name.
	Domainname string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

	// Expose a port or a range of ports.
	Expose string

	// GPU devices to add to the container ('all' to pass all GPUs).
	Gpus string

	// Add additional groups to join.
	GroupAdd string

	// Command to run to check health.
	HealthCmd string

	// Time between running the check (ms|s|m|h) (default 0s).
	HealthInterval string

	// Consecutive failures needed to report unhealthy.
	HealthRetries *int

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
	Label string

	// Read in a line delimited file of labels.
	LabelFile string

	// Add link to another container.
	Link string

	// Container IPv4/IPv6 link-local addresses.
	LinkLocalIp string

	// Logging driver for the container.
	LogDriver string

	// Log driver options.
	LogOpt string

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
	NetAlias string

	// Connect a container to a network.
	Network string

	// Add network-scoped alias for the container.
	NetworkAlias string

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
	Publish string

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
	SecurityOpt string

	// Size of /dev/shm.
	ShmSize string

	// Signal to stop the container.
	StopSignal string

	// Timeout (in seconds) to stop a container.
	StopTimeout *int

	// Storage driver options for the container.
	StorageOpt string

	// Sysctl options.
	Sysctl string

	// Mount a tmpfs directory.
	Tmpfs string

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
	Volume string

	// Optional volume driver for the container.
	VolumeDriver string

	// Mount volumes from the specified container(s).
	VolumesFrom string

	// Working directory inside the container.
	Workdir string
}

// Create a new container.
func Create(opts *CreateOpts, image string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"create"},
		append([]string{image, command}, arg...),
		opts,
		0,
	)
}

type DiffOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Inspect changes to files or directories on a container's filesystem.
func Diff(opts *DiffOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"diff"},
		[]string{container},
		opts,
		-1,
	)
}

type EventsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Format the output using the given Go template.
	Format string

	// Print usage.
	Help bool

	// Show all events created since timestamp.
	Since string

	// Stream events until this timestamp.
	Until string
}

// Get real time events from the server.
func Events(opts *EventsOpts) (string, error) {
	return runCtrCmd(
		[]string{"events"},
		[]string{},
		opts,
		0,
	)
}

type ExecOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Detached mode: run command in the background.
	Detach bool

	// Override the key sequence for detaching a container.
	DetachKeys string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

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
func Exec(opts *ExecOpts, container string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"exec"},
		append([]string{container, command}, arg...),
		opts,
		0,
	)
}

type ExportOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Write to a file, instead of STDOUT.
	Output string
}

// Export a container's filesystem as a tar archive.
func Export(opts *ExportOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"export"},
		[]string{container},
		opts,
		0,
	)
}

type HistoryOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Print sizes and dates in human readable format.
	Human bool

	// Don't truncate output.
	NoTrunc bool

	// Only show image IDs.
	Quiet bool
}

// Show the history of an image.
func History(opts *HistoryOpts, image string) (string, error) {
	return runCtrCmd(
		[]string{"history"},
		[]string{image},
		opts,
		0,
	)
}

type ImageOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage images.
func Image(opts *ImageOpts) (string, error) {
	return runCtrCmd(
		[]string{"image"},
		[]string{},
		opts,
		-1,
	)
}

type ImagesOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show all images (default hides intermediate images).
	All bool

	// Show digests.
	Digests bool

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

	// Don't truncate output.
	NoTrunc bool

	// Only show image IDs.
	Quiet bool
}

// List images.
func Images(opts *ImagesOpts, repositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"images"},
		[]string{repositoryTag},
		opts,
		0,
	)
}

type ImportOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Apply Dockerfile instruction to the created image.
	Change string

	// Print usage.
	Help bool

	// Set commit message for imported image.
	Message string

	// Set platform if server is multi-platform capable.
	Platform string
}

// Import the contents from a tarball to create a filesystem image.
func Import(opts *ImportOpts, fileUrl string, RepositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"import"},
		[]string{fileUrl, RepositoryTag},
		opts,
		0,
	)
}

type InfoOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool
}

// Display system-wide information.
func Info(opts *InfoOpts) (string, error) {
	return runCtrCmd(
		[]string{"info"},
		[]string{},
		opts,
		0,
	)
}

type InspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Display total file sizes if the type is container.
	Size bool

	// Return JSON for specified type.
	Type string
}

// Return low-level information on Docker objects.
func Inspect(opts *InspectOpts, nameId ...string) (string, error) {
	if len(nameId) == 0 {
		return "", fmt.Errorf("nameId must have at least one value")
	}
	return runCtrCmd(
		[]string{"inspect"},
		nameId,
		opts,
		0,
	)
}

type KillOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Signal to send to the container.
	Signal string
}

// Kill one or more running containers.
func Kill(opts *KillOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"kill"},
		container,
		opts,
		0,
	)
}

type LoadOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Read from tar archive file, instead of STDIN.
	Input string

	// Suppress the load output.
	Quiet bool
}

// Load an image from a tar archive or STDIN.
func Load(opts *LoadOpts) (string, error) {
	return runCtrCmd(
		[]string{"load"},
		[]string{},
		opts,
		0,
	)
}

type LoginOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Password.
	Password string

	// Take the password from stdin.
	PasswordStdin bool

	// Username.
	Username string
}

// Log in to a registry.
func Login(opts *LoginOpts, server string) (string, error) {
	return runCtrCmd(
		[]string{"login"},
		[]string{server},
		opts,
		0,
	)
}

type LogoutOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Log out from a registry.
func Logout(opts *LogoutOpts, server string) (string, error) {
	return runCtrCmd(
		[]string{"logout"},
		[]string{server},
		opts,
		-1,
	)
}

type LogsOpts struct {
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
func Logs(opts *LogsOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"logs"},
		[]string{container},
		opts,
		0,
	)
}

type ManifestOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Docker image manifests and manifest lists.
func Manifest(opts *ManifestOpts, command string) (string, error) {
	return runCtrCmd(
		[]string{"manifest"},
		[]string{command},
		opts,
		-1,
	)
}

type NetworkOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage networks.
func Network(opts *NetworkOpts) (string, error) {
	return runCtrCmd(
		[]string{"network"},
		[]string{},
		opts,
		-1,
	)
}

type NodeOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Swarm nodes.
func Node(opts *NodeOpts) (string, error) {
	return runCtrCmd(
		[]string{"node"},
		[]string{},
		opts,
		-1,
	)
}

type PauseOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Pause all processes within one or more containers.
func Pause(opts *PauseOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"pause"},
		container,
		opts,
		-1,
	)
}

type PluginOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage plugins.
func Plugin(opts *PluginOpts) (string, error) {
	return runCtrCmd(
		[]string{"plugin"},
		[]string{},
		opts,
		-1,
	)
}

type PortOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// List port mappings or a specific mapping for the container.
func Port(opts *PortOpts, container string, privatePortProto string) (string, error) {
	return runCtrCmd(
		[]string{"port"},
		[]string{container, privatePortProto},
		opts,
		-1,
	)
}

type PsOpts struct {
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
func Ps(opts *PsOpts) (string, error) {
	return runCtrCmd(
		[]string{"ps"},
		[]string{},
		opts,
		0,
	)
}

type PullOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Download all tagged images in the repository.
	AllTags bool

	// Skip image verification.
	DisableContentTrust bool

	// Print usage.
	Help bool

	// Set platform if server is multi-platform capable.
	Platform string

	// Suppress verbose output.
	Quiet bool
}

// Download an image from a registry.
func Pull(opts *PullOpts, nameTagDigest string) (string, error) {
	return runCtrCmd(
		[]string{"pull"},
		[]string{nameTagDigest},
		opts,
		0,
	)
}

type PushOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Push all tags of an image to the repository.
	AllTags bool

	// Skip image signing.
	DisableContentTrust bool

	// Print usage.
	Help bool

	// Suppress verbose output.
	Quiet bool
}

// Upload an image to a registry.
func Push(opts *PushOpts, nameTag string) (string, error) {
	return runCtrCmd(
		[]string{"push"},
		[]string{nameTag},
		opts,
		0,
	)
}

type RenameOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Rename a container.
func Rename(opts *RenameOpts, container string, newName string) (string, error) {
	return runCtrCmd(
		[]string{"rename"},
		[]string{container, newName},
		opts,
		-1,
	)
}

type RestartOpts struct {
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
func Restart(opts *RestartOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"restart"},
		container,
		opts,
		0,
	)
}

type RmOpts struct {
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
func Rm(opts *RmOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"rm"},
		container,
		opts,
		0,
	)
}

type RmiOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force removal of the image.
	Force bool

	// Print usage.
	Help bool

	// Do not delete untagged parents.
	NoPrune bool
}

// Remove one or more images.
func Rmi(opts *RmiOpts, image ...string) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"rmi"},
		image,
		opts,
		0,
	)
}

type RunOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Add a custom host-to-IP mapping (host:ip).
	AddHost string

	// Add an annotation to the container (passed through to the OCI runtime).
	Annotation string

	// Attach to STDIN, STDOUT or STDERR.
	Attach string

	// Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0).
	BlkioWeight string

	// Block IO weight (relative device weight).
	BlkioWeightDevice string

	// Add Linux capabilities.
	CapAdd string

	// Drop Linux capabilities.
	CapDrop string

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
	Device string

	// Add a rule to the cgroup allowed devices list.
	DeviceCgroupRule string

	// Limit read rate (bytes per second) from a device.
	DeviceReadBps string

	// Limit read rate (IO per second) from a device.
	DeviceReadIops string

	// Limit write rate (bytes per second) to a device.
	DeviceWriteBps string

	// Limit write rate (IO per second) to a device.
	DeviceWriteIops string

	// Skip image verification.
	DisableContentTrust bool

	// Set custom DNS servers.
	Dns string

	// Set DNS options.
	DnsOpt string

	// Set DNS options.
	DnsOption string

	// Set custom DNS search domains.
	DnsSearch string

	// Container NIS domain name.
	Domainname string

	// Overwrite the default ENTRYPOINT of the image.
	Entrypoint string

	// Set environment variables.
	Env string

	// Read in a file of environment variables.
	EnvFile string

	// Expose a port or a range of ports.
	Expose string

	// GPU devices to add to the container ('all' to pass all GPUs).
	Gpus string

	// Add additional groups to join.
	GroupAdd string

	// Command to run to check health.
	HealthCmd string

	// Time between running the check (ms|s|m|h) (default 0s).
	HealthInterval string

	// Consecutive failures needed to report unhealthy.
	HealthRetries *int

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
	Label string

	// Read in a line delimited file of labels.
	LabelFile string

	// Add link to another container.
	Link string

	// Container IPv4/IPv6 link-local addresses.
	LinkLocalIp string

	// Logging driver for the container.
	LogDriver string

	// Log driver options.
	LogOpt string

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
	NetAlias string

	// Connect a container to a network.
	Network string

	// Add network-scoped alias for the container.
	NetworkAlias string

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
	Publish string

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
	SecurityOpt string

	// Size of /dev/shm.
	ShmSize string

	// Proxy received signals to the process.
	SigProxy bool

	// Signal to stop the container.
	StopSignal string

	// Timeout (in seconds) to stop a container.
	StopTimeout *int

	// Storage driver options for the container.
	StorageOpt string

	// Sysctl options.
	Sysctl string

	// Mount a tmpfs directory.
	Tmpfs string

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
	Volume string

	// Optional volume driver for the container.
	VolumeDriver string

	// Mount volumes from the specified container(s).
	VolumesFrom string

	// Working directory inside the container.
	Workdir string
}

// Create and run a new container from an image.
func Run(opts *RunOpts, image string, command string, arg ...string) (string, error) {
	return runCtrCmd(
		[]string{"run"},
		append([]string{image, command}, arg...),
		opts,
		0,
	)
}

type SaveOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Write to a file, instead of STDOUT.
	Output string
}

// Save one or more images to a tar archive (streamed to STDOUT by default).
func Save(opts *SaveOpts, image ...string) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"save"},
		image,
		opts,
		0,
	)
}

type SearchOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Pretty-print search using a Go template.
	Format string

	// Print usage.
	Help bool

	// Max number of search results.
	Limit *int

	// Don't truncate output.
	NoTrunc bool
}

// Search Docker Hub for images.
func Search(opts *SearchOpts, term string) (string, error) {
	return runCtrCmd(
		[]string{"search"},
		[]string{term},
		opts,
		0,
	)
}

type SecretOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Swarm secrets.
func Secret(opts *SecretOpts) (string, error) {
	return runCtrCmd(
		[]string{"secret"},
		[]string{},
		opts,
		-1,
	)
}

type ServiceOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Swarm services.
func Service(opts *ServiceOpts) (string, error) {
	return runCtrCmd(
		[]string{"service"},
		[]string{},
		opts,
		-1,
	)
}

type StackOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Orchestrator to use (swarm|all).
	Orchestrator string
}

// Manage Swarm stacks.
func Stack(opts *StackOpts) (string, error) {
	return runCtrCmd(
		[]string{"stack"},
		[]string{},
		opts,
		0,
	)
}

type StartOpts struct {
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
func Start(opts *StartOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"start"},
		container,
		opts,
		0,
	)
}

type StatsOpts struct {
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
func Stats(opts *StatsOpts, container ...string) (string, error) {
	return runCtrCmd(
		[]string{"stats"},
		container,
		opts,
		0,
	)
}

type StopOpts struct {
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
func Stop(opts *StopOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"stop"},
		container,
		opts,
		0,
	)
}

type SwarmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Swarm.
func Swarm(opts *SwarmOpts) (string, error) {
	return runCtrCmd(
		[]string{"swarm"},
		[]string{},
		opts,
		-1,
	)
}

type SystemOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage Docker.
func System(opts *SystemOpts) (string, error) {
	return runCtrCmd(
		[]string{"system"},
		[]string{},
		opts,
		-1,
	)
}

type TagOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
func Tag(opts *TagOpts, sourceImageTag string, targetImageTag string) (string, error) {
	return runCtrCmd(
		[]string{"tag"},
		[]string{sourceImageTag, targetImageTag},
		opts,
		-1,
	)
}

type TopOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Display the running processes of a container.
func Top(opts *TopOpts, container string, psOptions string) (string, error) {
	return runCtrCmd(
		[]string{"top"},
		[]string{container, psOptions},
		opts,
		-1,
	)
}

type TrustOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage trust on Docker images.
func Trust(opts *TrustOpts) (string, error) {
	return runCtrCmd(
		[]string{"trust"},
		[]string{},
		opts,
		-1,
	)
}

type UnpauseOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Unpause all processes within one or more containers.
func Unpause(opts *UnpauseOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"unpause"},
		container,
		opts,
		-1,
	)
}

type UpdateOpts struct {
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
func Update(opts *UpdateOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"update"},
		container,
		opts,
		0,
	)
}

type VersionOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool
}

// Show the Docker version information.
func Version(opts *VersionOpts) (string, error) {
	return runCtrCmd(
		[]string{"version"},
		[]string{},
		opts,
		0,
	)
}

type VolumeOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage volumes.
func Volume(opts *VolumeOpts, command string) (string, error) {
	return runCtrCmd(
		[]string{"volume"},
		[]string{command},
		opts,
		-1,
	)
}

type WaitOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Block until one or more containers stop, then print their exit codes.
func Wait(opts *WaitOpts, container ...string) (string, error) {
	if len(container) == 0 {
		return "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"wait"},
		container,
		opts,
		-1,
	)
}

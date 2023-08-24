package ctrctl

type ContainerRunOpts struct {
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
	Detach string

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
	DisableContentTrust string

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
	HealthRetries string

	// Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s).
	HealthStartPeriod string

	// Maximum time to allow one check to run (ms|s|m|h) (default 0s).
	HealthTimeout string

	// Print usage.
	Help string

	// Container host name.
	Hostname string

	// Run an init inside the container that forwards signals and reaps processes.
	Init string

	// Keep STDIN open even if not attached.
	Interactive string

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
	NoHealthcheck string

	// Disable OOM Killer.
	OomKillDisable string

	// Tune host's OOM preferences (-1000 to 1000).
	OomScoreAdj string

	// PID namespace to use.
	Pid string

	// Tune container pids limit (set -1 for unlimited).
	PidsLimit string

	// Set platform if server is multi-platform capable.
	Platform string

	// Give extended privileges to this container.
	Privileged string

	// Publish a container's port(s) to the host.
	Publish string

	// Publish all exposed ports to random ports.
	PublishAll string

	// Pull image before running (`always`, `missing`, `never`).
	Pull string

	// Suppress the pull output.
	Quiet string

	// Mount the container's root filesystem as read only.
	ReadOnly string

	// Restart policy to apply when a container exits.
	Restart string

	// Automatically remove the container when it exits.
	Rm string

	// Runtime to use for this container.
	Runtime string

	// Security Options.
	SecurityOpt string

	// Size of /dev/shm.
	ShmSize string

	// Proxy received signals to the process.
	SigProxy string

	// Signal to stop the container.
	StopSignal string

	// Timeout (in seconds) to stop a container.
	StopTimeout string

	// Storage driver options for the container.
	StorageOpt string

	// Sysctl options.
	Sysctl string

	// Mount a tmpfs directory.
	Tmpfs string

	// Allocate a pseudo-TTY.
	Tty string

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
func ContainerRun(opts *ContainerRunOpts, image string, command string, arg ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "container", "run" },
		append([]string{ image,command }, arg...),
		opts,
		0,
	)
}

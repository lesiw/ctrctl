package ctrctl

type BuildOpts struct {
	// Add a custom host-to-IP mapping (`host:ip`).
	AddHost string

	// Set build-time variables.
	BuildArg string

	// Images to consider as cache sources.
	CacheFrom string

	// Optional parent cgroup for the container.
	CgroupParent string

	// Compress the build context using gzip.
	Compress string

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
	DisableContentTrust string

	// Name of the Dockerfile (Default is `PATH/Dockerfile`).
	File string

	// Always remove intermediate containers.
	ForceRm string

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
	NoCache string

	// Set platform if server is multi-platform capable.
	Platform string

	// Always attempt to pull a newer version of the image.
	Pull string

	// Suppress the build output and print image ID on success.
	Quiet string

	// Remove intermediate containers after a successful build.
	Rm string

	// Security options.
	SecurityOpt string

	// Size of `/dev/shm`.
	ShmSize string

	// Squash newly built layers into a single new layer.
	Squash string

	// Name and optionally a tag in the `name:tag` format.
	Tag string

	// Set the target build stage to build.
	Target string

	// Ulimit options.
	Ulimit string
}

// Build an image from a Dockerfile.
func Build(opts *BuildOpts, path string, url string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "build" },
		[]string{ path, url },
		opts,
		0,
	)
}

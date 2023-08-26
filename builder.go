package ctrctl

import "os/exec"

type BuilderBuildOpts struct {
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
func BuilderBuild(opts *BuilderBuildOpts, path string, url string) (string, error) {
	return runCtrCmd(
		[]string{"builder", "build"},
		[]string{path, url},
		opts,
		0,
	)
}

type BuilderPruneOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Remove all unused build cache, not just dangling ones.
	All bool

	// Provide filter values (e.g. `until=24h`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Print usage.
	Help bool

	// Amount of disk space to keep for cache.
	KeepStorage string
}

// Remove build cache.
func BuilderPrune(opts *BuilderPruneOpts) (string, error) {
	return runCtrCmd(
		[]string{"builder", "prune"},
		[]string{},
		opts,
		-1,
	)
}

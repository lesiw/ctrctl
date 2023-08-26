package ctrctl

import (
	"fmt"
	"os/exec"
)

type ImageBuildOpts struct {
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
func ImageBuild(opts *ImageBuildOpts, path string, url string) (string, error) {
	return runCtrCmd(
		[]string{"image", "build"},
		[]string{path, url},
		opts,
		0,
	)
}

type ImageHistoryOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:.
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
func ImageHistory(opts *ImageHistoryOpts, image string) (string, error) {
	return runCtrCmd(
		[]string{"image", "history"},
		[]string{image},
		opts,
		0,
	)
}

type ImageImportOpts struct {
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
func ImageImport(opts *ImageImportOpts, fileUrl string, RepositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"image", "import"},
		[]string{fileUrl, RepositoryTag},
		opts,
		0,
	)
}

type ImageInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool
}

// Display detailed information on one or more images.
func ImageInspect(opts *ImageInspectOpts, image ...string) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"image", "inspect"},
		image,
		opts,
		0,
	)
}

type ImageLoadOpts struct {
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
func ImageLoad(opts *ImageLoadOpts) (string, error) {
	return runCtrCmd(
		[]string{"image", "load"},
		[]string{},
		opts,
		0,
	)
}

type ImageLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Show all images (default hides intermediate images).
	All bool

	// Show digests.
	Digests bool

	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
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
func ImageLs(opts *ImageLsOpts, repositoryTag string) (string, error) {
	return runCtrCmd(
		[]string{"image", "ls"},
		[]string{repositoryTag},
		opts,
		0,
	)
}

type ImagePruneOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Remove all unused images, not just dangling ones.
	All bool

	// Provide filter values (e.g. `until=<timestamp>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Print usage.
	Help bool
}

// Remove unused images.
func ImagePrune(opts *ImagePruneOpts) (string, error) {
	return runCtrCmd(
		[]string{"image", "prune"},
		[]string{},
		opts,
		0,
	)
}

type ImagePullOpts struct {
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
func ImagePull(opts *ImagePullOpts, nameTagDigest string) (string, error) {
	return runCtrCmd(
		[]string{"image", "pull"},
		[]string{nameTagDigest},
		opts,
		0,
	)
}

type ImagePushOpts struct {
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
func ImagePush(opts *ImagePushOpts, nameTag string) (string, error) {
	return runCtrCmd(
		[]string{"image", "push"},
		[]string{nameTag},
		opts,
		0,
	)
}

type ImageRmOpts struct {
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
func ImageRm(opts *ImageRmOpts, image ...string) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"image", "rm"},
		image,
		opts,
		0,
	)
}

type ImageSaveOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Write to a file, instead of STDOUT.
	Output string
}

// Save one or more images to a tar archive (streamed to STDOUT by default).
func ImageSave(opts *ImageSaveOpts, image ...string) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{"image", "save"},
		image,
		opts,
		0,
	)
}

type ImageTagOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
func ImageTag(opts *ImageTagOpts, sourceImageTag string, targetImageTag string) (string, error) {
	return runCtrCmd(
		[]string{"image", "tag"},
		[]string{sourceImageTag, targetImageTag},
		opts,
		-1,
	)
}

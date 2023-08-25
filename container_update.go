package ctrctl

type ContainerUpdateOpts struct {
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
func ContainerUpdate(opts *ContainerUpdateOpts, container ...string) (
	stdout string, stderr string, err error) {
	if len(container) == 0 {
		return "", "", fmt.Errorf("container must have at least one value")
	}
	return runCtrCmd(
		[]string{"container", "update"},
		container,
		opts,
		0,
	)
}

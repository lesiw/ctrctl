package ctrctl

import "fmt"

type VolumeCreateOpts struct {
	// Cluster Volume availability (`active`, `pause`, `drain`).
	Availability string

	// Specify volume driver name.
	Driver string

	// Cluster Volume group (cluster volumes).
	Group string

	// Print usage.
	Help bool

	// Set metadata for a volume.
	Label string

	// Minimum size of the Cluster Volume in bytes.
	LimitBytes string

	// Specify volume name.
	Name string

	// Set driver specific options.
	Opt string

	// Maximum size of the Cluster Volume in bytes.
	RequiredBytes string

	// Cluster Volume access scope (`single`, `multi`).
	Scope string

	// Cluster Volume secrets.
	Secret string

	// Cluster Volume access sharing (`none`, `readonly`, `onewriter`, `all`).
	Sharing string

	// A topology that the Cluster Volume would be preferred in.
	TopologyPreferred string

	// A topology that the Cluster Volume must be accessible from.
	TopologyRequired string

	// Cluster Volume access type (`mount`, `block`).
	Type string
}

// Create a volume.
func VolumeCreate(opts *VolumeCreateOpts, volume string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "create"},
		[]string{volume},
		opts,
		0,
	)
}

type VolumeInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool
}

// Display detailed information on one or more volumes.
func VolumeInspect(opts *VolumeInspectOpts, volume ...string) (
	stdout string, stderr string, err error) {
	if len(volume) == 0 {
		return "", "", fmt.Errorf("volume must have at least one value")
	}
	return runCtrCmd(
		[]string{"volume", "inspect"},
		volume,
		opts,
		0,
	)
}

type VolumeLsOpts struct {
	// Display only cluster volumes, and use cluster volume list formatting.
	Cluster bool

	// Provide filter values (e.g. `dangling=true`).
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

	// Only display volume names.
	Quiet bool
}

// List volumes.
func VolumeLs(opts *VolumeLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "ls"},
		[]string{},
		opts,
		0,
	)
}

type VolumePruneOpts struct {
	// Remove all unused volumes, not just anonymous ones.
	All bool

	// Provide filter values (e.g. `label=<label>`).
	Filter string

	// Do not prompt for confirmation.
	Force bool

	// Print usage.
	Help bool
}

// Remove all unused local volumes.
func VolumePrune(opts *VolumePruneOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "prune"},
		[]string{},
		opts,
		0,
	)
}

type VolumeRmOpts struct {
	// Force the removal of one or more volumes.
	Force bool

	// Print usage.
	Help bool
}

// Remove one or more volumes.
func VolumeRm(opts *VolumeRmOpts, volume ...string) (
	stdout string, stderr string, err error) {
	if len(volume) == 0 {
		return "", "", fmt.Errorf("volume must have at least one value")
	}
	return runCtrCmd(
		[]string{"volume", "rm"},
		volume,
		opts,
		0,
	)
}

type VolumeUpdateOpts struct {
	// Cluster Volume availability (`active`, `pause`, `drain`).
	Availability string

	// Print usage.
	Help bool
}

// Update a volume (cluster volumes only).
func VolumeUpdate(opts *VolumeUpdateOpts, volume string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"volume", "update"},
		[]string{volume},
		opts,
		0,
	)
}

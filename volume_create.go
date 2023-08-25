package ctrctl

type VolumeCreateOpts struct {
	// Cluster Volume availability (`active`, `pause`, `drain`).
	Availability string

	// Specify volume driver name.
	Driver string

	// Cluster Volume group (cluster volumes).
	Group string

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

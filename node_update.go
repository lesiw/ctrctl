package ctrctl

type NodeUpdateOpts struct {
	// Availability of the node (`active`, `pause`, `drain`).
	Availability string

	// Add or update a node label (`key=value`).
	LabelAdd string

	// Remove a node label if exists.
	LabelRm string

	// Role of the node (`worker`, `manager`).
	Role string
}

// Update a node.
func NodeUpdate(opts *NodeUpdateOpts, node string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"node", "update"},
		[]string{node},
		opts,
		0,
	)
}

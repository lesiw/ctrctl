package ctrctl

import "fmt"

type NodeDemoteOpts struct {
}

// Demote one or more nodes from manager in the swarm.
func NodeDemote(opts *NodeDemoteOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "demote"},
		node,
		opts,
		-1,
	)
}

type NodeInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more nodes.
func NodeInspect(opts *NodeInspectOpts, selfNode string, node ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"node", "inspect"},
		append([]string{selfNode}, node...),
		opts,
		0,
	)
}

type NodeLsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Only display IDs.
	Quiet bool
}

// List nodes in the swarm.
func NodeLs(opts *NodeLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"node", "ls"},
		[]string{},
		opts,
		0,
	)
}

type NodePromoteOpts struct {
}

// Promote one or more nodes to manager in the swarm.
func NodePromote(opts *NodePromoteOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "promote"},
		node,
		opts,
		-1,
	)
}

type NodePsOpts struct {
	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Only display task IDs.
	Quiet bool
}

// List tasks running on one or more nodes, defaults to current node.
func NodePs(opts *NodePsOpts, node ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"node", "ps"},
		node,
		opts,
		0,
	)
}

type NodeRmOpts struct {
	// Force remove a node from the swarm.
	Force bool
}

// Remove one or more nodes from the swarm.
func NodeRm(opts *NodeRmOpts, node ...string) (
	stdout string, stderr string, err error) {
	if len(node) == 0 {
		return "", "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "rm"},
		node,
		opts,
		0,
	)
}

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

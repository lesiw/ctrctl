package ctrctl

import (
	"fmt"
	"os/exec"
)

type NodeDemoteOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Demote one or more nodes from manager in the swarm.
func NodeDemote(opts *NodeDemoteOpts, node ...string) (string, error) {
	if len(node) == 0 {
		return "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "demote"},
		node,
		opts,
		-1,
	)
}

type NodeInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Format output using a custom template:
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more nodes.
func NodeInspect(opts *NodeInspectOpts, selfNode string, node ...string) (string, error) {
	return runCtrCmd(
		[]string{"node", "inspect"},
		append([]string{selfNode}, node...),
		opts,
		0,
	)
}

type NodeLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

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

	// Only display IDs.
	Quiet bool
}

// List nodes in the swarm.
func NodeLs(opts *NodeLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"node", "ls"},
		[]string{},
		opts,
		0,
	)
}

type NodePromoteOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Promote one or more nodes to manager in the swarm.
func NodePromote(opts *NodePromoteOpts, node ...string) (string, error) {
	if len(node) == 0 {
		return "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "promote"},
		node,
		opts,
		-1,
	)
}

type NodePsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Filter output based on conditions provided.
	Filter string

	// Pretty-print tasks using a Go template.
	Format string

	// Print usage.
	Help bool

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Only display task IDs.
	Quiet bool
}

// List tasks running on one or more nodes, defaults to current node.
func NodePs(opts *NodePsOpts, node ...string) (string, error) {
	return runCtrCmd(
		[]string{"node", "ps"},
		node,
		opts,
		0,
	)
}

type NodeRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force remove a node from the swarm.
	Force bool

	// Print usage.
	Help bool
}

// Remove one or more nodes from the swarm.
func NodeRm(opts *NodeRmOpts, node ...string) (string, error) {
	if len(node) == 0 {
		return "", fmt.Errorf("node must have at least one value")
	}
	return runCtrCmd(
		[]string{"node", "rm"},
		node,
		opts,
		0,
	)
}

type NodeUpdateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Availability of the node (`active`, `pause`, `drain`).
	Availability string

	// Print usage.
	Help bool

	// Add or update a node label (`key=value`).
	LabelAdd string

	// Remove a node label if exists.
	LabelRm string

	// Role of the node (`worker`, `manager`).
	Role string
}

// Update a node.
func NodeUpdate(opts *NodeUpdateOpts, node string) (string, error) {
	return runCtrCmd(
		[]string{"node", "update"},
		[]string{node},
		opts,
		0,
	)
}

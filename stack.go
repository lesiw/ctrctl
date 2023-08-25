package ctrctl

import "fmt"

type StackConfigOpts struct {
	// Path to a Compose file, or `-` to read from stdin.
	ComposeFile string

	// Print usage.
	Help bool

	// Orchestrator to use (swarm|all).
	Orchestrator string

	// Skip interpolation and output only merged config.
	SkipInterpolation bool
}

// Outputs the final config file, after doing merges and interpolations.
func StackConfig(opts *StackConfigOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"stack", "config"},
		[]string{},
		opts,
		0,
	)
}

type StackDeployOpts struct {
	// Path to a Compose file, or `-` to read from stdin.
	ComposeFile string

	// Print usage.
	Help bool

	// Orchestrator to use (swarm|all).
	Orchestrator string

	// Prune services that are no longer referenced.
	Prune bool

	// Query the registry to resolve image digest and supported platforms (`always`, `changed`, `never`).
	ResolveImage string

	// Send registry authentication details to Swarm agents.
	WithRegistryAuth bool
}

// Deploy a new stack or update an existing stack.
func StackDeploy(opts *StackDeployOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"stack", "deploy"},
		[]string{stack},
		opts,
		0,
	)
}

type StackLsOpts struct {
	// Format output using a custom template:.
	// 'table':            Print output in table format with column headers (default).
	// 'table TEMPLATE':   Print output in table format using the given Go template.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print usage.
	Help bool

	// Orchestrator to use (swarm|all).
	Orchestrator string
}

// List stacks.
func StackLs(opts *StackLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"stack", "ls"},
		[]string{},
		opts,
		0,
	)
}

type StackPsOpts struct {
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

	// Do not map IDs to Names.
	NoResolve bool

	// Do not truncate output.
	NoTrunc bool

	// Orchestrator to use (swarm|all).
	Orchestrator string

	// Only display task IDs.
	Quiet bool
}

// List the tasks in the stack.
func StackPs(opts *StackPsOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"stack", "ps"},
		[]string{stack},
		opts,
		0,
	)
}

type StackRmOpts struct {
	// Print usage.
	Help bool

	// Orchestrator to use (swarm|all).
	Orchestrator string
}

// Remove one or more stacks.
func StackRm(opts *StackRmOpts, stack ...string) (
	stdout string, stderr string, err error) {
	if len(stack) == 0 {
		return "", "", fmt.Errorf("stack must have at least one value")
	}
	return runCtrCmd(
		[]string{"stack", "rm"},
		stack,
		opts,
		0,
	)
}

type StackServicesOpts struct {
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

	// Orchestrator to use (swarm|all).
	Orchestrator string

	// Only display IDs.
	Quiet bool
}

// List the services in the stack.
func StackServices(opts *StackServicesOpts, stack string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"stack", "services"},
		[]string{stack},
		opts,
		0,
	)
}

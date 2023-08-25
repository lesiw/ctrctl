package ctrctl

import "fmt"

type ConfigCreateOpts struct {
	// Config labels.
	Label string

	// Template driver.
	TemplateDriver string
}

// Create a config from a file or STDIN.
func ConfigCreate(opts *ConfigCreateOpts, config string, file string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"config", "create"},
		[]string{config, file},
		opts,
		0,
	)
}

type ConfigInspectOpts struct {
	// Format output using a custom template:.
	// 'json':             Print in JSON format.
	// 'TEMPLATE':         Print output using the given Go template.
	// Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates.
	Format string

	// Print the information in a human friendly format.
	Pretty bool
}

// Display detailed information on one or more configs.
func ConfigInspect(opts *ConfigInspectOpts, config ...string) (
	stdout string, stderr string, err error) {
	if len(config) == 0 {
		return "", "", fmt.Errorf("config must have at least one value")
	}
	return runCtrCmd(
		[]string{"config", "inspect"},
		config,
		opts,
		0,
	)
}

type ConfigLsOpts struct {
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

// List configs.
func ConfigLs(opts *ConfigLsOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{"config", "ls"},
		[]string{},
		opts,
		0,
	)
}

type ConfigRmOpts struct {
}

// Remove one or more configs.
func ConfigRm(opts *ConfigRmOpts, config ...string) (
	stdout string, stderr string, err error) {
	if len(config) == 0 {
		return "", "", fmt.Errorf("config must have at least one value")
	}
	return runCtrCmd(
		[]string{"config", "rm"},
		config,
		opts,
		-1,
	)
}

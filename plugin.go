package ctrctl

import (
	"fmt"
	"os/exec"
)

type PluginCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Compress the context using gzip.
	Compress bool

	// Print usage.
	Help bool
}

// Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory.
func PluginCreate(opts *PluginCreateOpts, plugin string, pluginDataDir string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "create"},
		[]string{plugin, pluginDataDir},
		opts,
		0,
	)
}

type PluginDisableOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force the disable of an active plugin.
	Force bool

	// Print usage.
	Help bool
}

// Disable a plugin.
func PluginDisable(opts *PluginDisableOpts, plugin string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "disable"},
		[]string{plugin},
		opts,
		0,
	)
}

type PluginEnableOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// HTTP client timeout (in seconds).
	Timeout *int
}

// Enable a plugin.
func PluginEnable(opts *PluginEnableOpts, plugin string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "enable"},
		[]string{plugin},
		opts,
		0,
	)
}

type PluginInspectOpts struct {
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

// Display detailed information on one or more plugins.
func PluginInspect(opts *PluginInspectOpts, plugin ...string) (string, error) {
	if len(plugin) == 0 {
		return "", fmt.Errorf("plugin must have at least one value")
	}
	return runCtrCmd(
		[]string{"plugin", "inspect"},
		plugin,
		opts,
		0,
	)
}

type PluginInstallOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Local name for plugin.
	Alias string

	// Do not enable the plugin on install.
	Disable bool

	// Skip image verification.
	DisableContentTrust bool

	// Grant all permissions necessary to run the plugin.
	GrantAllPermissions bool

	// Print usage.
	Help bool
}

// Install a plugin.
func PluginInstall(opts *PluginInstallOpts, plugin string, keyValue ...string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "install"},
		append([]string{plugin}, keyValue...),
		opts,
		0,
	)
}

type PluginLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Provide filter values (e.g. `enabled=true`).
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

	// Only display plugin IDs.
	Quiet bool
}

// List plugins.
func PluginLs(opts *PluginLsOpts) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "ls"},
		[]string{},
		opts,
		0,
	)
}

type PluginPushOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Skip image signing.
	DisableContentTrust bool

	// Print usage.
	Help bool
}

// Push a plugin to a registry.
func PluginPush(opts *PluginPushOpts, pluginTag string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "push"},
		[]string{pluginTag},
		opts,
		0,
	)
}

type PluginRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Force the removal of an active plugin.
	Force bool

	// Print usage.
	Help bool
}

// Remove one or more plugins.
func PluginRm(opts *PluginRmOpts, plugin ...string) (string, error) {
	if len(plugin) == 0 {
		return "", fmt.Errorf("plugin must have at least one value")
	}
	return runCtrCmd(
		[]string{"plugin", "rm"},
		plugin,
		opts,
		0,
	)
}

type PluginSetOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Change settings for a plugin.
func PluginSet(opts *PluginSetOpts, plugin string, keyValue ...string) (string, error) {
	if len(keyValue) == 0 {
		return "", fmt.Errorf("keyValue must have at least one value")
	}
	return runCtrCmd(
		[]string{"plugin", "set"},
		append([]string{plugin}, keyValue...),
		opts,
		-1,
	)
}

type PluginUpgradeOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Skip image verification.
	DisableContentTrust bool

	// Grant all permissions necessary to run the plugin.
	GrantAllPermissions bool

	// Print usage.
	Help bool

	// Do not check if specified remote plugin matches existing plugin image.
	SkipRemoteCheck bool
}

// Upgrade an existing plugin.
func PluginUpgrade(opts *PluginUpgradeOpts, plugin string, remote string) (string, error) {
	return runCtrCmd(
		[]string{"plugin", "upgrade"},
		[]string{plugin, remote},
		opts,
		0,
	)
}

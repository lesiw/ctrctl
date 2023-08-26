package ctrctl

import (
	"fmt"
	"os/exec"
)

type TrustInspectOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Print the information in a human friendly format.
	Pretty bool
}

// Return low-level information about keys and signatures.
func TrustInspect(opts *TrustInspectOpts, imageTag ...string) (string, error) {
	if len(imageTag) == 0 {
		return "", fmt.Errorf("imageTag must have at least one value")
	}
	return runCtrCmd(
		[]string{"trust", "inspect"},
		imageTag,
		opts,
		-1,
	)
}

type TrustKeyOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage keys for signing Docker images.
func TrustKey(opts *TrustKeyOpts) (string, error) {
	return runCtrCmd(
		[]string{"trust", "key"},
		[]string{},
		opts,
		-1,
	)
}

type TrustKeyGenerateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Directory to generate key in, defaults to current directory.
	Dir string

	// Print usage.
	Help bool
}

// Generate and load a signing key-pair.
func TrustKeyGenerate(opts *TrustKeyGenerateOpts, name string) (string, error) {
	return runCtrCmd(
		[]string{"trust", "key", "generate"},
		[]string{name},
		opts,
		-1,
	)
}

type TrustKeyLoadOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Name for the loaded key.
	Name string
}

// Load a private key file for signing.
func TrustKeyLoad(opts *TrustKeyLoadOpts, keyfile string) (string, error) {
	return runCtrCmd(
		[]string{"trust", "key", "load"},
		[]string{keyfile},
		opts,
		0,
	)
}

type TrustRevokeOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Do not prompt for confirmation.
	Yes bool
}

// Remove trust for an image.
func TrustRevoke(opts *TrustRevokeOpts, imageTag string) (string, error) {
	return runCtrCmd(
		[]string{"trust", "revoke"},
		[]string{imageTag},
		opts,
		0,
	)
}

type TrustSignOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Sign a locally tagged image.
	Local bool
}

// Sign an image.
func TrustSign(opts *TrustSignOpts, imageTag string) (string, error) {
	return runCtrCmd(
		[]string{"trust", "sign"},
		[]string{imageTag},
		opts,
		-1,
	)
}

type TrustSignerOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool
}

// Manage entities who can sign Docker images.
func TrustSigner(opts *TrustSignerOpts) (string, error) {
	return runCtrCmd(
		[]string{"trust", "signer"},
		[]string{},
		opts,
		-1,
	)
}

type TrustSignerAddOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Print usage.
	Help bool

	// Path to the signer's public key file.
	Key string
}

// Add a signer.
func TrustSignerAdd(opts *TrustSignerAddOpts, name string, repository ...string) (string, error) {
	if len(repository) == 0 {
		return "", fmt.Errorf("repository must have at least one value")
	}
	return runCtrCmd(
		[]string{"trust", "signer", "add"},
		append([]string{name}, repository...),
		opts,
		0,
	)
}

type TrustSignerRemoveOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Do not prompt for confirmation before removing the most recent signer.
	Force bool

	// Print usage.
	Help bool
}

// Remove a signer.
func TrustSignerRemove(opts *TrustSignerRemoveOpts, name string, repository ...string) (string, error) {
	if len(repository) == 0 {
		return "", fmt.Errorf("repository must have at least one value")
	}
	return runCtrCmd(
		[]string{"trust", "signer", "remove"},
		append([]string{name}, repository...),
		opts,
		0,
	)
}

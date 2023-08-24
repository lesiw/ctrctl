package ctrctl

type SwarmCaOpts struct {
	// Path to the PEM-formatted root CA certificate to use for the new cluster.
	CaCert string

	// Path to the PEM-formatted root CA key to use for the new cluster.
	CaKey string

	// Validity period for node certificates (ns|us|ms|s|m|h).
	CertExpiry string

	// Exit immediately instead of waiting for the root rotation to converge.
	Detach bool

	// Specifications of one or more certificate signing endpoints.
	ExternalCa string

	// Suppress progress output.
	Quiet bool

	// Rotate the swarm CA - if no certificate or key are provided, new ones will be generated.
	Rotate bool
}

// Display and rotate the root CA.
func SwarmCa(opts *SwarmCaOpts) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "swarm", "ca" },
		[]string{},
		opts,
		0,
	)
}

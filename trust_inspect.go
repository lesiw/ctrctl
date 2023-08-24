package ctrctl

type TrustInspectOpts struct {
	// Print the information in a human friendly format.
	Pretty bool
}

// Return low-level information about keys and signatures.
func TrustInspect(opts *TrustInspectOpts, imageTag string, extraImageTag ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "trust", "inspect" },
		append([]string{ imageTag }, extraImageTag...),
		opts,
		-1,
	)
}

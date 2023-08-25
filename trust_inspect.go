package ctrctl

type TrustInspectOpts struct {
	// Print the information in a human friendly format.
	Pretty bool
}

// Return low-level information about keys and signatures.
func TrustInspect(opts *TrustInspectOpts, imageTag ...string) (
	stdout string, stderr string, err error) {
	if len(imageTag) == 0 {
		return "", "", fmt.Errorf("imageTag must have at least one value")
	}
	return runCtrCmd(
		[]string{"trust", "inspect"},
		imageTag,
		opts,
		-1,
	)
}

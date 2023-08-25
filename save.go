package ctrctl

import "fmt"

type SaveOpts struct {
	// Write to a file, instead of STDOUT.
	Output string
}

// Save one or more images to a tar archive (streamed to STDOUT by default).
func Save(opts *SaveOpts, image ...string) (
	stdout string, stderr string, err error) {
	if len(image) == 0 {
		return "", "", fmt.Errorf("image must have at least one value")
	}
	return runCtrCmd(
		[]string{ "save" },
		image,
		opts,
		0,
	)
}

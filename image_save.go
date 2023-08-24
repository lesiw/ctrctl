package ctrctl

type ImageSaveOpts struct {
	// Write to a file, instead of STDOUT.
	Output string
}

// Save one or more images to a tar archive (streamed to STDOUT by default).
func ImageSave(opts *ImageSaveOpts, image string, extraImage ...string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "image", "save" },
		append([]string{ image }, extraImage...),
		opts,
		0,
	)
}

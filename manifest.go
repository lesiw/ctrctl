package ctrctl

type ManifestOpts struct {

}

// Manage Docker image manifests and manifest lists.
func Manifest(opts *ManifestOpts, command string) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		[]string{ "manifest" },
		[]string{ command },
		opts,
		-1,
	)
}

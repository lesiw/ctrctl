# ctrctl

[![Go Reference](https://pkg.go.dev/badge/github.com/lesiw/ctrctl.svg)](https://pkg.go.dev/github.com/lesiw/ctrctl)

Package `ctrctl` wraps container CLIs.

## Example

```go
package main

import "github.com/lesiw/ctrctl"

func main() {
    ctrctl.Cli = []string{"docker"} // or {"podman"}, or {"lima", "nerdctl"}...
    ctrctl.Verbose = true           // useful for debugging

    id, _, err := ctrctl.ContainerRun(
        &ctrctl.ContainerRunOpts{
            Detach: true,
            Tty:    true,
        },
        "alpine",
        "cat",
    )
    if err != nil {
        panic(err)
    }

    _, _, err = ctrctl.ContainerExec(nil, id, "echo", "Hello from alpine!")
    if err != nil {
        panic(err)
    }

    _, _, err = ctrctl.ContainerRm(&ctrctl.ContainerRmOpts{Force: true}, id)
    if err != nil {
        panic(err)
    }
}
```

## Motivation

The [Docker API](https://docs.docker.com/engine/api/sdk/) is not portable
between `docker` and alternate container engines like `containerd` and `podman`.
The [Docker CLI](https://docs.docker.com/engine/reference/commandline/cli/),
however, is a well-established interface that has been implemented by other
tools such as `nerdctl` and the `podman` CLI.

This means working with CLIs is the only reliable way to produce container
management code that works across multiple engines, but running commands is both
tedious and prone to typos. `ctrctl` provides a more idiomatic experience for
interacting with Docker-compatible CLIs.

`ctrctl`’s wrapper functions are generated from the
[Dockermentation](https://github.com/docker/docs). No container engine
implements Docker's entire interface, but generating `ctrctl` from Docker
ensures that all potential functionality is covered.

To switch between `docker` and other Docker-compatible CLIs, set `ctrctl.Cli` to
the appropriate value.

## Simple usage

All wrapper functions return a `string` containing stdout and an optional
`error`.

`error` may be of type `ctrctl.CliError`, which contains a `Stderr` field for
debugging purposes.

## Advanced usage

All wrapper functions’ options structs have a `Cmd` field. Set this to an
`&exec.Cmd` to override the default command behavior.

Note that setting `exec.Cmd.Stdout` or `exec.Cmd.Stderr` to an `io.Writer` will
disable automatic capture of those outputs. Bypassing capture allows the
underlying container CLI to work in interactive mode when attached to
`os.Stdout` and `os.Stderr`.

In this example, standard streams are overridden to expose the output of the
`ImagePull` to the end user, then drop them into an interactive shell in an
`alpine` container. Once they exit the shell, the container is removed.

```go
package main

import (
    "os"
    "os/exec"

    "github.com/lesiw/ctrctl"
)

func main() {
    _, err := ctrctl.ImagePull(
        &ctrctl.ImagePullOpts{
            Cmd: &exec.Cmd{
                Stdout: os.Stdout,
                Stderr: os.Stderr,
            },
        },
        "alpine:latest",
    )
    if err != nil {
        panic(err)
    }

    id, err := ctrctl.ContainerRun(
        &ctrctl.ContainerRunOpts{
            Detach: true,
            Tty:    true,
        },
        "alpine",
        "cat",
    )
    if err != nil {
        panic(err)
    }

    _, _ = ctrctl.Exec(
        &ctrctl.ExecOpts{
            Cmd: &exec.Cmd{
                Stdin:  os.Stdin,
                Stdout: os.Stdout,
                Stderr: os.Stderr,
            },
            Interactive: true,
            Tty:         true,
        },
        id,
        "/bin/sh",
    )

    _, err = ctrctl.ContainerRm(&ctrctl.ContainerRmOpts{Force: true}, id)
    if err != nil {
        panic(err)
    }
}
```

## Regenerating the interface

Install [`gofmt`](https://pkg.go.dev/cmd/gofmt) and
[`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports), then run `go
generate ./...`

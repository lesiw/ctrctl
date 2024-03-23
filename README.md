# ctrctl

[![Go Reference](https://pkg.go.dev/badge/lesiw.io/ctrctl.svg)](https://pkg.go.dev/lesiw.io/ctrctl)

Package `ctrctl` wraps container CLIs.

## Example

```go
package main

import "lesiw.io/ctrctl"

func main() {
    ctrctl.Cli = []string{"docker"} // or {"podman"}, or {"lima", "nerdctl"}...

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

    out, err := ctrctl.ContainerExec(nil, id, "echo", "Hello from alpine!")
    if err != nil {
        panic(err)
    }
    fmt.Println(out)

    _, err = ctrctl.ContainerRm(&ctrctl.ContainerRmOpts{Force: true}, id)
    if err != nil {
        panic(err)
    }
}
```

## Motivation

Most container engines ship alongside a CLI that provides an experience similar
to that of Docker on Linux. In the pursuit of this goal, they often hide several
layers of indirection through internal APIs and VM boundaries while maintaining
argument-for-argument compatibility with the [Docker
CLI](https://docs.docker.com/engine/reference/commandline/cli/).

Since the CLI is the primary interface that users and automation scripts
interact with, it’s also the most likely interface where bugs will be noticed
and, hopefully, fixed. Conversely, the primary consumer of engines’ internal
APIs are their own command lines and a handful of plugins, so issues and
documentation gaps in the API layer are less likely to be noticed and
prioritized.

For these reasons, Docker-compatible CLIs serve as excellent abstraction points
for projects that manage containers. `docker`, `nerdctl`, `podman`, and even
`kubectl` have more in common with one another than any of their internal APIs
or SDKs do. However, working with `exec.Command` is verbose and lacks in-editor
completion for container commands.

`ctrctl` fills this gap by providing CLI wrapper functions and option structs
automatically generated from the
[Dockermentation](https://github.com/docker/docs). While no container engine
implements Docker’s entire interface, generating `ctrctl` wrappers from Docker
ensures that all potential shared functionality will be covered.

To switch between `docker` and other Docker-compatible CLIs, just set
`ctrctl.Cli` to the appropriate value.

## Simple usage

All wrapper functions take an optional struct as the first argument. The format
of the option struct is always `CommandOpts`, where `Command` is the name of the
function being called.

Commands return a `string` containing stdout and an optional `error`. `error`
may be of type `ctrctl.CliError`, which contains a `Stderr` field for debugging
purposes.

Set `ctrctl.Verbose = true` to stream the exact commands being run, along with
their output, to standard out. The format is similar to using `set +x` in a
shell script.

## Advanced usage

All wrapper functions’ options structs have a `Cmd` field. Set this to an
`&exec.Cmd` to override the default command behavior.

Note that setting `Cmd.Stdout` or `Cmd.Stderr` to an `io.Writer` will disable
automatic capture of those outputs. Bypassing capture allows the underlying
container CLI to work in interactive mode when they are attached to `os.Stdout`
and `os.Stderr`, respectively.

In this example, standard streams are overridden to expose the output of the
`ImagePull` to the end user, then drop them into an interactive shell in an
`alpine` container. Once they exit the shell, the container is removed.

```go
package main

import (
    "os"
    "os/exec"

    "lesiw.io/ctrctl"
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

    _, _ = ctrctl.ContainerExec(
        &ctrctl.ContainerExecOpts{
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

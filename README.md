go-pkgs
=======
Package pkgs finds all packages in all of the `GOPATH` trees. It is
library functionality equivalent to `go list all` (see `go help
packages` for more info).

Much of the code was adapted from the cmd/go `matchPackages` func.

Docs: [godoc.org/github.com/sqs/go-pkgs](http://godoc.org/github.com/sqs/go-pkgs)

## Installation

	go get github.com/sqs/go-pkgs

## Example Usage

    import (
        "github.com/sqs/go-pkgs"
        "go/build
    )

    // List all packages in all GOPATH trees.
    pkglist, err := pkgs.FindAll("", build.Default, 0)

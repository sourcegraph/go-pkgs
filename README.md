go-pkgs
=======

[![xrefs](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-pkgs/badges/xrefs.png)](https://sourcegraph.com/github.com/sourcegraph/go-pkgs)
[![funcs](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-pkgs/badges/funcs.png)](https://sourcegraph.com/github.com/sourcegraph/go-pkgs)
[![top func](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-pkgs/badges/top-func.png)](https://sourcegraph.com/github.com/sourcegraph/go-pkgs)
[![library users](https://sourcegraph.com/api/repos/github.com/sourcegraph/go-pkgs/badges/library-users.png)](https://sourcegraph.com/github.com/sourcegraph/go-pkgs)

Package pkgs finds all packages in all of the `GOPATH` trees. It is
library functionality equivalent to `go list all` (see `go help
packages` for more info).

Much of the code was adapted from the cmd/go `matchPackages` func.

Docs: [go-pkgs on Sourcegraph](https://sourcegraph.com/github.com/sourcegraph/go-pkgs)

## Installation

	go get github.com/sourcegraph/go-pkgs

## Example Usage

    import (
        "github.com/sourcegraph/go-pkgs"
        "go/build
    )

    // List all packages in all GOPATH trees.
    pkglist, err := pkgs.FindAll("", build.Default, 0)

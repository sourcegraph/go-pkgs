// Package pkgs finds all packages in all of the GOPATH trees. It is
// library functionality equivalent to `go list all` (see `go help
// packages` for more info).
//
// Much of the code was adapted from the cmd/go matchPackages func.
package pkgs

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

type FindMode uint

const (
	IncludeStdlib FindMode = 1 << iota
)

// FindAll returns a list of all packages in all of the GOPATH trees
// in the given build context. If prefix is non-empty, only packages
// whose import paths begin with prefix are returned.
func FindAll(prefix string, buildContext build.Context, mode FindMode) (pkgs []*build.Package, err error) {
	have := map[string]bool{
		"builtin": true, // ignore pseudo-package that exists only for documentation
	}
	if !buildContext.CgoEnabled {
		have["runtime/cgo"] = true // ignore during walk
	}

	// TODO(sqs): find cmd packages as well

	var gorootSrcPkg = filepath.Join(buildContext.GOROOT, "src/pkg")

	for _, src := range buildContext.SrcDirs() {
		if src == gorootSrcPkg && mode&IncludeStdlib == 0 {
			continue // skip stdlib
		}
		src = filepath.Clean(src) + string(filepath.Separator)
		start := filepath.Join(src, prefix)
		filepath.Walk(start, func(path string, fi os.FileInfo, err error) error {
			if err != nil || !fi.IsDir() || path == src {
				return nil
			}

			// Avoid .foo, _foo, and testdata directory trees.
			_, elem := filepath.Split(path)
			if strings.HasPrefix(elem, ".") || strings.HasPrefix(elem, "_") || elem == "testdata" {
				return filepath.SkipDir
			}

			name := filepath.ToSlash(path[len(start):])
			if have[name] {
				return nil
			}
			have[name] = true

			pkg, err := buildContext.ImportDir(path, 0)
			if err != nil && strings.Contains(err.Error(), "no Go source files") {
				return nil
			}
			pkgs = append(pkgs, pkg)
			return nil
		})
	}
	return pkgs, nil
}

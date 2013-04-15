package pkgs

import (
	"go/build"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestFindAll(t *testing.T) {
	expected := []*build.Package{
		{ImportPath: "a"},
		{ImportPath: "b"},
		{ImportPath: "b/c"},
	}

	buildContext := build.Default
	buildContext.GOPATH, _ = filepath.Abs("testdata/")
	actual, err := FindAll(buildContext, 0)
	if err != nil {
		t.Fatalf("FindAll failed: %v", err)
	}
	packageListsImportPathsAreEqual(t, actual, expected)
}

func packageListsImportPathsAreEqual(t *testing.T, actual []*build.Package, expected []*build.Package) {
	apaths := packageListImportPaths(actual)
	epaths := packageListImportPaths(expected)
	sort.Strings(apaths)
	sort.Strings(epaths)
	if !reflect.DeepEqual(apaths, epaths) {
		t.Fatalf("package lists differ in import paths\nexpected: %s\ngot:      %s", epaths, apaths)
	}
}

func packageListImportPaths(ps []*build.Package) []string {
	ips := make([]string, len(ps))
	for i, p := range ps {
		ips[i] = p.ImportPath
	}
	return ips
}

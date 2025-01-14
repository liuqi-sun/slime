package testfiles

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

var filesources []FileSource

// FileSource implements one way of retrieving test file content.  For
// example, one file source could read from the original source code
// file tree, another from bindata compiled into a test executable.
type FileSource interface {
	// ReadTestFile retrieves the content of a file that gets maintained
	// alongside a test's source code. Files are identified by the
	// relative path inside the repository containing the tests, for
	// example "cluster/gce/upgrade.sh" inside kubernetes/kubernetes.
	//
	// When the file is not found, a nil slice is returned. An error is
	// returned for all fatal errors.
	ReadTestFile(filePath string) ([]byte, error)

	// DescribeFiles returns a multi-line description of which
	// files are available via this source. It is meant to be
	// used as part of the error message when a file cannot be
	// found.
	DescribeFiles() string
}

func AddFileSource(filesource FileSource) {
	filesources = append(filesources, filesource)
}

// Read tries to retrieve the desired file content from
// one of the registered file sources.
func Read(filePath string) ([]byte, error) {
	if len(filesources) == 0 {
		return nil, fmt.Errorf("no file sources registered (yet?), cannot retrieve test file %s", filePath)
	}
	for _, filesource := range filesources {
		data, err := filesource.ReadTestFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("fatal error retrieving test file %s: %s", filePath, err)
		}
		if data != nil {
			return data, nil
		}
	}
	// Here we try to generate an error that points test authors
	// or users in the right direction for resolving the problem.
	error := fmt.Sprintf("Test file %q was not found.\n", filePath)
	for _, filesource := range filesources {
		error += filesource.DescribeFiles()
		error += "\n"
	}
	return nil, errors.New(error)
}

// RootFileSource looks for files relative to a root directory.
type RootFileSource struct {
	Root string
}

// ReadTestFile looks for the file relative to the configured
// root directory. If the path is already absolute, for example
// in a test that has its own method of determining where
// files are, then the path will be used directly.
func (r RootFileSource) ReadTestFile(filePath string) ([]byte, error) {
	var fullPath string
	if path.IsAbs(filePath) {
		fullPath = filePath
	} else {
		fullPath = filepath.Join(r.Root, filePath)
	}
	data, err := ioutil.ReadFile(fullPath)
	if os.IsNotExist(err) {
		// Not an error (yet), some other provider may have the file.
		return nil, nil
	}
	return data, err
}

// DescribeFiles explains that it looks for files inside a certain
// root directory.
func (r RootFileSource) DescribeFiles() string {
	description := fmt.Sprintf("Test files are expected in %q", r.Root)
	if !path.IsAbs(r.Root) {
		// The default in test_context.go is the relative path
		// ../../, which doesn't really help locating the
		// actual location. Therefore we add also the absolute
		// path if necessary.
		abs, err := filepath.Abs(r.Root)
		if err == nil {
			description += fmt.Sprintf(" = %q", abs)
		}
	}
	description += "."
	return description
}

// BindataFileSource handles files stored in a package generated with bindata.
type BindataFileSource struct {
	Asset      func(string) ([]byte, error)
	AssetNames func() []string
}

// ReadTestFile looks for an asset with the given path.
func (b BindataFileSource) ReadTestFile(filePath string) ([]byte, error) {
	fileBytes, err := b.Asset(filePath)
	if err != nil {
		// It would be nice to have a better way to detect
		// "not found" errors :-/
		if strings.HasSuffix(err.Error(), "not found") {
			return nil, nil
		}
	}
	return fileBytes, nil
}

// DescribeFiles explains about gobindata and then lists all available files.
func (b BindataFileSource) DescribeFiles() string {
	var lines []string
	lines = append(lines, "The following files are built into the test executable via gobindata. For questions on maintaining gobindata, contact the sig-testing group.")
	assets := b.AssetNames()
	sort.Strings(assets)
	lines = append(lines, assets...)
	description := strings.Join(lines, "\n   ")
	return description
}
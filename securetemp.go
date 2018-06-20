package securetemp

import (
	"io/ioutil"
	"os"
)

// TODO(leon): It's {K,M,G}iB, not {K,M,G}B (with an "i")
const (
	// SizeKB represents one Kibibyte
	SizeKB = 1 << (10 * iota)
	// SizeMB represents one Mebibyte (1024 KiB)
	SizeMB
	// SizeGB represents one Gibibyte (1024 MiB)
	SizeGB
)

const (
	// DefaultSize specifies the default size for a RAM disk in MB
	DefaultSize = 4 * SizeMB

	// globalPrefix will be used if we need a prefix for
	// something (e.g. files & folders)
	globalPrefix = "securetemp"
)

// TempDir creates a new RAM disk with size 'size' (in bytes)
// and returns the path to it.
// Use this function only if you intend to create multiple
// files inside your RAM disk, else prefer to use 'TempFile'.
func TempDir(size int) (string, func(), error) {
	path, cleanupFunc, err := createRAMDisk(size)
	if err != nil {
		return "", nil, err
	}
	return path, cleanupFunc, nil
}

// TempFile creates a new RAM disk with size 'size' (in bytes),
// creates a temp file in it and returns a pointer to that file.
// Use this function only if you intend to create a single file
// inside your RAM disk, else prefer to use 'TempDir'.
func TempFile(size int) (*os.File, func(), error) {
	path, cleanupFunc, err := TempDir(size)
	if err != nil {
		return nil, nil, err
	}
	doCleanup := true
	defer func() {
		if doCleanup {
			cleanupFunc()
		}
	}()

	file, err := ioutil.TempFile(path, globalPrefix)
	if err != nil {
		return nil, nil, err
	}

	doCleanup = false
	return file, func() { _ = file.Close(); cleanupFunc() }, nil
}

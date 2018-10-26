package securetemp

import (
	"os"
	"testing"
)

const (
	size = DefaultSize
)

func TestTempDir(t *testing.T) {
	tmpDir, cleanupFunc, err := TempDir(size)
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	if _, err := os.Stat(tmpDir); err != nil {
		t.Fatalf("something is wrong with the temp dir: %v", err)
	}

	cleanupFunc()
	if _, err := os.Stat(tmpDir); err == nil {
		t.Fatalf("temp dir %q should no longer exist after cleanup, but does", tmpDir)
	}
}

func TestTempFile(t *testing.T) {
	tmpFile, cleanupFunc, err := TempFile(size)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := tmpFile.Stat(); err != nil {
		t.Fatalf("something is wrong with the temp file: %v", err)
	}

	cleanupFunc()
	if _, err := tmpFile.Stat(); err == nil {
		t.Fatalf("temp dir %q should no longer exist after cleanup, but does", tmpFile.Name())
	}
}

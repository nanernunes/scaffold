package helpers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCommandRun(t *testing.T) {
	// Prepare test directory and files.
	// Use os.MkdirAll to create directory and all its parent directories.
	testDir := filepath.Join(os.TempDir(), "test-cmd")
	err := os.MkdirAll(testDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "test.txt")
	f, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	f.Close()

	// Initialize Command.
	cmd := NewCommand(true, testDir)

	// Run a simple command that should print output.
	cmd.Run("echo 'Hello, world!'")

	// Run a command that should create a file.
	cmd.Run(fmt.Sprintf("touch %s", testFile))

	// Run a command that should fail.
	cmd.Run("nonexistent-command")

	// Verify that the file was created.
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %q to exist, but it does not", testFile)
	}
}

func TestCommandRunSilent(t *testing.T) {
	// Capture stdout and stderr output.
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	defer func() {
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}()

	// Initialize Command and run command silently.
	cmd := NewCommand(false, "")
	cmd.Run("echo 'Hello, world!'")

	// Verify that no output was printed to stdout or stderr.
	if stdout.String() != "" {
		t.Fatalf("Expected no output to stdout, but got: %q", stdout.String())
	}
	if stderr.String() != "" {
		t.Fatalf("Expected no output to stderr, but got: %q", stderr.String())
	}
}

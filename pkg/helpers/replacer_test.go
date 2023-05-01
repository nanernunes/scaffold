package helpers

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFileContains(t *testing.T) {
	// Create a temporary file
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write some content to the file
	content := "hello world"
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}

	// Check that the file contains the content
	if !FileContains(tmpfile.Name(), content) {
		t.Errorf("Expected file to contain '%s'", content)
	}
}

func TestAppendInFile(t *testing.T) {
	// Create a temporary file
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write some content to the file
	content := "one\ntwo\nend\n"
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}

	// Append a line to the file
	newLine := "new\n"
	AppendInFile(tmpfile.Name(), "two", "end", newLine)

	// Read the contents of the file
	bytes, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	fileContent := string(bytes)

	// Check that the new line was appended to the file
	expectedContent := "one\ntwo\nnew\nend\n"
	if strings.TrimSpace(fileContent) != strings.TrimSpace(expectedContent) {
		t.Errorf("Expected file content to be '%s', got '%s'", expectedContent, fileContent)
	}
}

package version

import (
	"testing"
)

// TestVersionGet tests the Get function
func TestVersionGet(t *testing.T) {
	// Get the version
	version := Get()

	// Check that the version is not empty
	if version == "" {
		t.Error("Expected version to not be empty")
	}

	// Check that the version matches the constant
	if version != VERSION {
		t.Errorf("Expected version to be %s, got %s", VERSION, version)
	}
}

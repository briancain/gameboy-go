package version

import (
	"testing"
)

// TestVersionGet tests the Get function
func TestVersionGet(t *testing.T) {
	// Get the version
	version := GetVersion()

	// Check that the version is not empty
	if version == "" {
		t.Error("Expected version to not be empty")
	}

	// Check that the version matches the variable
	if version != Version {
		t.Errorf("Expected version to be %s, got %s", Version, version)
	}
}

// TestGetRef tests the GetRef function
func TestGetRef(t *testing.T) {
	// Get the ref
	ref := GetRef()

	// Check that the ref matches the variable
	if ref != Ref {
		t.Errorf("Expected ref to be %s, got %s", Ref, ref)
	}
}

// TestGetRelease tests the GetRelease function
func TestGetRelease(t *testing.T) {
	// Save original values
	origVersion := Version
	origRef := Ref

	// Test case 1: No ref
	Version = "v1.0.0"
	Ref = ""
	release := GetRelease()
	if release != "v1.0.0" {
		t.Errorf("Expected release to be v1.0.0, got %s", release)
	}

	// Test case 2: With ref
	Version = "v1.0.0"
	Ref = "abc123"
	release = GetRelease()
	if release != "v1.0.0+abc123" {
		t.Errorf("Expected release to be v1.0.0+abc123, got %s", release)
	}

	// Restore original values
	Version = origVersion
	Ref = origRef
}

// TestGet tests the Get function
func TestGet(t *testing.T) {
	// Save original values
	origVersion := Version
	origRef := Ref

	// Test case 1: No ref
	Version = "v1.0.0"
	Ref = ""
	formatted := Get()
	if formatted != "GameBoy-Go v1.0.0" {
		t.Errorf("Expected formatted version to be 'GameBoy-Go v1.0.0', got '%s'", formatted)
	}

	// Test case 2: With ref
	Version = "v1.0.0"
	Ref = "abc123"
	formatted = Get()
	if formatted != "GameBoy-Go v1.0.0 (abc123)" {
		t.Errorf("Expected formatted version to be 'GameBoy-Go v1.0.0 (abc123)', got '%s'", formatted)
	}

	// Restore original values
	Version = origVersion
	Ref = origRef
}

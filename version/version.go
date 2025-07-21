package version

var (
	// Version is the current version of the application
	Version = "v0.0.1"
	// Ref is the git ref of the application, if not an official release
	Ref = ""
)

// GetVersion returns the version of the application
func GetVersion() string {
	return Version
}

// GetRef returns the git ref of the application
func GetRef() string {
	return Ref
}

// GetRelease returns the version with the git ref appended if available
func GetRelease() string {
	version := Version
	if Ref != "" {
		version += "+" + Ref
	}

	return version
}

// Get returns a formatted version string for display
func Get() string {
	if Ref == "" {
		return "GameBoy-Go " + Version
	}
	return "GameBoy-Go " + Version + " (" + Ref + ")"
}

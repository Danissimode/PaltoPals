// version.go: Build-time versioning info for Paltopals (set via -ldflags)
package internal

// Version, Commit, Date are injected at build time.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

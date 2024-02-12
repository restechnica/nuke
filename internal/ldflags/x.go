package ldflags

// -X key=value ldflags definitions, injected at build time.
var (
	// Version The release version of the nuke binary.
	Version string = "dev"
)

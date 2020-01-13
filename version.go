package goroota

var (
	// GitCommit is the git commit that was compiled. This will be filled in by
	// the compiler.
	GitCommit string
	// GitDescribe is the git description that was compiled. This will be filled
	// in by the compiler.
	GitDescribe string
)

// Version is the semantic version number being executed.
const Version = "1.0.0"

// VersionPrerelease identifies a pre-release marker for the version. If ""
// (empty string) then this is a final release. Otherwise, this a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = ""

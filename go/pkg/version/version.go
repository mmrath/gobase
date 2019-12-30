package version

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	// Version is baked by go build -ldflags "-X version.Version=$VERSION"
	Version string
	// GitCommit is baked by go build -ldflags "-X version.GitCommit=$GIT_COMMIT"
	GitCommit string
	// BuildTime is baked by go build -ldflags "-X 'version.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'"
	BuildTime string
)

// PrintVersionOrContinue will print git commit and exit with os.Exit(0) if CLI v flag is present
func PrintVersion() {
	if Version != "" || GitCommit != "" || BuildTime != "" {
		fmt.Printf("version: %s (%s) | %s", Version, GitCommit, BuildTime)
		log.Info().
			Str("version", Version).
			Str("GitCommit", GitCommit).
			Str("BuildTime", BuildTime).Msg("build info")
	}
	versionFlag := flag.Bool("v", false, "Print the current version and exit")

	flag.Parse()

	switch {
	case *versionFlag:
		fmt.Printf("version: %s (%s) | %s", Version, GitCommit, BuildTime)
		os.Exit(0)
	}
}

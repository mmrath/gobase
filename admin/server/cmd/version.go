package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version is set by the build scripts.
var Version = "NA"

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`Version: %s
Go Version: %s
Go OS/ARCH: %s %s
`, Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
}

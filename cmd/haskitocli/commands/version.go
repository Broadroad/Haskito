package commands

import (
	"runtime"

	"github.com/haskito/version"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Haskitocli",
	Run: func(cmd *cobra.Command, args []string) {
		jww.FEEDBACK.Printf("Haskitocli v%s %s/%s\n", version.Version, runtime.GOOS, runtime.GOARCH)
	},
}

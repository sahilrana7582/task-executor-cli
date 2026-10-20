package cmd

import (
	"os"

	"github.com/sahil27582/task-worker-cli/internal/ui"
	"github.com/spf13/cobra"
)

// rootCmd is the base command. Subcommands add themselves to it
// from their own files using init().
var rootCmd = &cobra.Command{
	Use:   "tw",
	Short: "TaskWorker - minimal task executor framework",
	Long:  ui.Banner(),
}

// Execute runs the CLI and exits with a non-zero code on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

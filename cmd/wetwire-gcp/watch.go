package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newWatchCmd creates the watch subcommand
func newWatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch [path]",
		Short: "Watch for changes and rebuild manifests",
		Long: `Watch Go files for changes and automatically rebuild GCP manifests.

Examples:
  wetwire-gcp watch
  wetwire-gcp watch ./gcp`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			// TODO: Implement watch
			fmt.Fprintf(cmd.OutOrStdout(), "GCP watch for %s (not yet implemented)\n", path)
			return nil
		},
	}

	return cmd
}

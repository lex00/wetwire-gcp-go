package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newDiffCmd creates the diff subcommand
func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff [path]",
		Short: "Show differences between Go code and generated manifests",
		Long: `Show what would change in the generated YAML manifests.

Examples:
  wetwire-gcp diff
  wetwire-gcp diff ./gcp`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			// TODO: Implement diff
			fmt.Fprintf(cmd.OutOrStdout(), "GCP diff for %s (not yet implemented)\n", path)
			return nil
		},
	}

	return cmd
}

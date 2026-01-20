package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newDesignCmd creates the design subcommand
func newDesignCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "design [description]",
		Short: "AI-assisted GCP infrastructure design",
		Long: `Design GCP infrastructure interactively with AI assistance.

Examples:
  wetwire-gcp design "Create a Cloud SQL instance"
  wetwire-gcp design "Set up a GKE cluster with Config Connector"`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement AI-assisted design
			fmt.Fprintln(cmd.OutOrStdout(), "GCP design mode (not yet implemented)")
			return nil
		},
	}

	return cmd
}

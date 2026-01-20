package main

import (
	"context"

	coredomain "github.com/lex00/wetwire-core-go/domain"
	gcpdomain "github.com/lex00/wetwire-gcp-go/domain"
	"github.com/spf13/cobra"
)

// newMCPCmd creates the mcp subcommand for Model Context Protocol server
func newMCPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Start MCP server for Claude Code integration",
		Long: `Start a Model Context Protocol (MCP) server for Claude Code integration.

The MCP server provides tools for:
  - wetwire_build: Generate GCP manifests
  - wetwire_lint: Check and fix code
  - wetwire_validate: Validate against Config Connector schemas

Configure in Claude Code:
  {
    "mcpServers": {
      "wetwire-gcp": {
        "command": "wetwire-gcp",
        "args": ["mcp"]
      }
    }
  }`,
		RunE: runMCPServer,
	}

	return cmd
}

// runMCPServer starts the MCP server on stdio transport.
func runMCPServer(cmd *cobra.Command, args []string) error {
	// Create GCP domain instance
	gcpDomain := &gcpdomain.GCPDomain{}

	// Build MCP server using auto-generation from domain
	server := coredomain.BuildMCPServer(gcpDomain)

	// Start stdio server
	return server.Start(context.Background())
}

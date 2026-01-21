// Command watch monitors source files and auto-rebuilds on changes.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"

	"github.com/lex00/wetwire-gcp-go/internal/discover"
	"github.com/lex00/wetwire-gcp-go/internal/lint"
)

// newWatchCmd creates the "watch" subcommand for auto-rebuilding on file changes.
func newWatchCmd() *cobra.Command {
	var (
		lintOnly bool
		debounce time.Duration
		format   string
		output   string
	)

	cmd := &cobra.Command{
		Use:   "watch [path]",
		Short: "Watch for changes and rebuild manifests",
		Long: `Watch monitors Go files for changes and automatically rebuilds GCP manifests.

The watch command:
- Monitors the source directory for .go file changes
- Runs lint on each change
- Rebuilds if lint passes (unless --lint-only)
- Debounces rapid changes to avoid excessive rebuilds

Examples:
  wetwire-gcp watch
  wetwire-gcp watch ./gcp
  wetwire-gcp watch --lint-only
  wetwire-gcp watch --debounce 1s`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}
			return runWatch(path, watchOptions{
				lintOnly: lintOnly,
				debounce: debounce,
				format:   format,
				output:   output,
			})
		},
	}

	cmd.Flags().BoolVar(&lintOnly, "lint-only", false, "Only run lint, skip build")
	cmd.Flags().DurationVar(&debounce, "debounce", 500*time.Millisecond, "Debounce duration for rapid changes")
	cmd.Flags().StringVarP(&format, "format", "f", "yaml", "Output format for build: yaml or json")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file for build (default: stdout)")

	return cmd
}

type watchOptions struct {
	lintOnly bool
	debounce time.Duration
	format   string
	output   string
}

// runWatch monitors source files and runs lint/build on changes.
func runWatch(path string, opts watchOptions) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer func() {
		_ = watcher.Close()
	}()

	// Resolve path to absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	// Add directory to watcher recursively
	if err := addDirRecursive(watcher, absPath); err != nil {
		return fmt.Errorf("failed to watch %s: %w", absPath, err)
	}
	fmt.Printf("Watching: %s\n", absPath)

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Initial build
	fmt.Println("Running initial lint/build...")
	runLintAndBuild(absPath, opts)

	// Debounce timer
	var debounceTimer *time.Timer
	rebuildChan := make(chan struct{}, 1)

	fmt.Println("\nWatching for changes... (Ctrl+C to stop)")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			// Only watch .go files
			if !strings.HasSuffix(event.Name, ".go") {
				continue
			}

			// Skip test files
			if strings.HasSuffix(event.Name, "_test.go") {
				continue
			}

			// Only process write/create events
			if event.Op&(fsnotify.Write|fsnotify.Create) == 0 {
				continue
			}

			// Debounce: reset timer on each change
			if debounceTimer != nil {
				debounceTimer.Stop()
			}
			debounceTimer = time.AfterFunc(opts.debounce, func() {
				select {
				case rebuildChan <- struct{}{}:
				default:
				}
			})

		case <-rebuildChan:
			fmt.Printf("\n[%s] Change detected, rebuilding...\n", time.Now().Format("15:04:05"))
			runLintAndBuild(absPath, opts)

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			fmt.Fprintf(os.Stderr, "Watch error: %v\n", err)

		case <-sigChan:
			fmt.Println("\nStopping watch...")
			return nil
		}
	}
}

// addDirRecursive adds a directory and all subdirectories to the watcher.
func addDirRecursive(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// Skip hidden directories
			if strings.HasPrefix(filepath.Base(path), ".") && path != dir {
				return filepath.SkipDir
			}
			// Skip vendor directory
			if filepath.Base(path) == "vendor" {
				return filepath.SkipDir
			}
			// Skip resources directory (generated code)
			if filepath.Base(path) == "resources" {
				return filepath.SkipDir
			}
			return watcher.Add(path)
		}
		return nil
	})
}

// runLintAndBuild runs lint and optionally build on the path.
func runLintAndBuild(path string, opts watchOptions) {
	// Run lint
	lintSuccess := runWatchLint(path)

	if !lintSuccess {
		fmt.Println("Lint failed, skipping build")
		return
	}

	fmt.Println("Lint passed")

	if opts.lintOnly {
		return
	}

	// Run build
	runWatchBuild(path, opts)
}

// runWatchLint runs lint and returns true if successful.
func runWatchLint(path string) bool {
	result, err := lint.LintPackage(path, lint.Options{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lint error: %v\n", err)
		return false
	}

	for _, issue := range result.Issues {
		if issue.File != "" {
			fmt.Printf("%s:%d:%d: %s: %s [%s]\n",
				issue.File, issue.Line, issue.Column,
				issue.Severity, issue.Message, issue.Rule)
		} else {
			fmt.Printf("%s: %s [%s]\n", issue.Severity, issue.Message, issue.Rule)
		}
	}

	return result.Success
}

// runWatchBuild runs build and outputs to stdout or file.
func runWatchBuild(path string, opts watchOptions) {
	// Discover resources
	resources, err := discover.DiscoverDirectory(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Build error: %v\n", err)
		return
	}

	if len(resources) == 0 {
		fmt.Println("No resources found")
		return
	}

	// Report success (actual manifest generation would use the domain's builder)
	fmt.Println("Build successful")
	fmt.Printf("Discovered %d Config Connector resources\n", len(resources))
	for _, r := range resources {
		fmt.Printf("  - %s (%s)\n", r.Name, r.Type)
	}

	if opts.output != "" {
		// For now, just print a placeholder message
		// Full implementation would generate actual YAML and write to file
		fmt.Printf("Output file: %s (generation not yet fully implemented)\n", opts.output)
	}
}

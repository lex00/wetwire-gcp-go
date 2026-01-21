// Command test runs automated persona-based testing.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lex00/wetwire-core-go/agent/personas"
	"github.com/lex00/wetwire-gcp-go/internal/kiro"
	"github.com/spf13/cobra"
)

// newTestCmd creates the "test" subcommand for automated persona-based testing.
func newTestCmd() *cobra.Command {
	var outputDir string
	var personaName string
	var scenario string
	var stream bool
	var allPersonas bool

	cmd := &cobra.Command{
		Use:   "test [prompt]",
		Short: "Run automated persona-based testing",
		Long: `Run automated testing with AI personas to evaluate code generation quality.

Available personas:
  - beginner: New to GCP, asks many clarifying questions
  - intermediate: Familiar with GCP basics, asks targeted questions
  - expert: Deep GCP knowledge, asks advanced questions

The test runs the Kiro agent in non-interactive mode and evaluates:
- Whether valid Go code was generated
- Whether the code passes lint checks
- Whether the code builds successfully

Example:
    wetwire-gcp test --persona beginner "Create a Cloud Storage bucket"
    wetwire-gcp test --all-personas "Create a GKE cluster"
    wetwire-gcp test "Create a VPC network"`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			prompt := args[0]
			if allPersonas {
				return runTestAllPersonas(prompt, outputDir, scenario, stream)
			}
			return runTest(prompt, outputDir, personaName, scenario, stream)
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for generated files")
	cmd.Flags().StringVarP(&personaName, "persona", "p", "intermediate", "Persona to use (beginner, intermediate, expert)")
	cmd.Flags().StringVarP(&scenario, "scenario", "S", "default", "Scenario name for tracking")
	cmd.Flags().BoolVarP(&stream, "stream", "s", false, "Stream AI responses")
	cmd.Flags().BoolVar(&allPersonas, "all-personas", false, "Run test with all personas")

	return cmd
}

// runTest executes a single persona test.
func runTest(prompt, outputDir, personaName, scenario string, stream bool) error {
	// Check if tests are disabled
	if os.Getenv("SKIP_KIRO_TESTS") == "1" {
		fmt.Println("Skipping Kiro test (SKIP_KIRO_TESTS=1)")
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nInterrupted, cleaning up...")
		cancel()
	}()

	// Validate persona name
	if _, err := personas.Get(personaName); err != nil {
		return fmt.Errorf("invalid persona: %w", err)
	}

	// Create test runner
	runner := kiro.NewTestRunner(outputDir)

	// Set up streaming if enabled
	if stream {
		runner.StreamHandler = func(text string) {
			fmt.Print(text)
		}
	}

	// Ensure test environment is ready
	if err := runner.EnsureTestEnvironment(); err != nil {
		return fmt.Errorf("preparing test environment: %w", err)
	}

	fmt.Printf("Running test with persona '%s' and scenario '%s'\n", personaName, scenario)
	fmt.Printf("Prompt: %s\n\n", prompt)

	// Get the persona object
	persona, err := personas.Get(personaName)
	if err != nil {
		return fmt.Errorf("invalid persona: %w", err)
	}

	// Run the test with persona
	result, err := runner.RunWithPersona(ctx, prompt, persona)
	if err != nil {
		return fmt.Errorf("test failed: %w", err)
	}

	// Print summary
	fmt.Println("\n--- Test Summary ---")
	fmt.Printf("Persona: %s\n", personaName)
	fmt.Printf("Scenario: %s\n", scenario)
	fmt.Printf("Duration: %s\n", result.Duration)
	fmt.Printf("Success: %v\n", result.Success)
	fmt.Printf("Lint passed: %v\n", result.LintPassed)
	fmt.Printf("Build passed: %v\n", result.BuildPassed)
	fmt.Printf("Files created: %d\n", len(result.FilesCreated))
	for _, f := range result.FilesCreated {
		fmt.Printf("  - %s\n", f)
	}
	if len(result.ErrorMessages) > 0 {
		fmt.Println("Errors:")
		for _, e := range result.ErrorMessages {
			fmt.Printf("  - %s\n", e)
		}
	}

	if !result.Success {
		return fmt.Errorf("test failed")
	}

	return nil
}

// runTestAllPersonas runs the test with all available personas sequentially.
func runTestAllPersonas(prompt, outputDir, scenario string, stream bool) error {
	personaNames := personas.Names()
	var failed []string

	fmt.Printf("Running tests with all %d personas\n\n", len(personaNames))

	for _, personaName := range personaNames {
		// Create persona-specific output directory
		personaOutputDir := fmt.Sprintf("%s/%s", outputDir, personaName)

		fmt.Printf("=== Running persona: %s ===\n", personaName)

		err := runTest(prompt, personaOutputDir, personaName, scenario, stream)
		if err != nil {
			fmt.Printf("Persona %s: FAILED - %v\n\n", personaName, err)
			failed = append(failed, personaName)
		} else {
			fmt.Printf("Persona %s: PASSED\n\n", personaName)
		}
	}

	// Print summary
	fmt.Println("\n=== All Personas Summary ===")
	fmt.Printf("Total: %d\n", len(personaNames))
	fmt.Printf("Passed: %d\n", len(personaNames)-len(failed))
	fmt.Printf("Failed: %d\n", len(failed))
	if len(failed) > 0 {
		fmt.Printf("Failed personas: %v\n", failed)
		return fmt.Errorf("%d personas failed", len(failed))
	}

	return nil
}

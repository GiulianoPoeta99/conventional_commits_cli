// Package main is the entry point of the Conventional Commits Assistant application.
package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	commit "github.com/GiulianoPoeta99/conventional_commits_cli/internal"
	t "github.com/GiulianoPoeta99/conventional_commits_cli/internal/types"
	ui "github.com/GiulianoPoeta99/conventional_commits_cli/internal/ui"
)

// main is the entry point for the application.
// It collects user inputs to build a commit message following Conventional Commits standards,
// then formats and executes the commit.
func Run() {
	// Print welcome message for the assistant.
	fmt.Println("🚀 Conventional Commits Assistant")

	// Initialize an empty commit configuration.
	config := t.CommitConfig{}

	var err error

	// Prompt user to select the commit type.
	config.Type, err = ui.SelectCommitType()
	if err != nil {
		fmt.Printf("Error selecting commit type: %v\n", err)
		os.Exit(1)
	}

	// Ask the user to provide an optional scope for the commit.
	config.Scope, err = ui.OptionalInput("Add a scope for this change. (optional, press Enter to omit)")
	if err != nil {
		fmt.Printf("Error entering scope: %v\n", err)
		os.Exit(1)
	}

	// Confirm if the user wants to include an emoji with the commit.
	useEmoji, err := ui.ConfirmSelect("Do you want to include an emoji?")
	if err != nil {
		fmt.Printf("Error selecting emoji option: %v\n", err)
		os.Exit(1)
	}

	// If emoji is desired, prompt for emoji selection with suggestions based on commit type.
	if useEmoji {
		config.Emoji, err = ui.SelectEmojiWithSuggestions(config.Type)
		if err != nil {
			fmt.Printf("Error selecting emoji: %v\n", err)
			os.Exit(1)
		}
	}

	// Request user input for the commit description with validation.
	config.Description, err = ui.InputWithValidation(
		"Commit description",
		"",
		func(input string) error {
			if len(input) < 3 {
				return errors.New("description must have at least 3 characters")
			}
			return nil
		},
	)
	if err != nil {
		fmt.Printf("Error entering description: %v\n", err)
		os.Exit(1)
	}

	// Ask for an optional commit body.
	config.Body, err = ui.OptionalInput("Commit body (optional, press Enter to omit)")
	if err != nil {
		fmt.Printf("Error entering body: %v\n", err)
		os.Exit(1)
	}

	// Confirm if the change is breaking.
	config.Breaking, err = ui.ConfirmSelect("Is this a breaking change?")
	if err != nil {
		fmt.Printf("Error selecting breaking change: %v\n", err)
		os.Exit(1)
	}

	// If the change is breaking, request an optional explanation.
	if config.Breaking {
		config.BreakingReason, err = ui.OptionalInput("Describe why this is a breaking change (optional, press Enter to use the default message)")
		if err != nil {
			fmt.Printf("Error entering breaking change reason: %v\n", err)
			os.Exit(1)
		}
	}

	// Confirm whether the user wants to add reviewers.
	addReviewers, err := ui.ConfirmSelect("Do you want to add reviewers?")
	if err != nil {
		fmt.Printf("Error asking about reviewers: %v\n", err)
		os.Exit(1)
	}

	// Collect the list of reviewers if confirmed.
	if addReviewers {
		for {
			reviewer, err := ui.InputWithValidation(
				"Enter reviewer (e.g., 'John Smith')",
				"",
				func(input string) error {
					if len(input) < 1 {
						return errors.New("reviewer name cannot be empty")
					}
					return nil
				},
			)
			if err != nil {
				fmt.Printf("Error entering reviewer: %v\n", err)
				os.Exit(1)
			}

			config.Reviewers = append(config.Reviewers, reviewer)

			addMore, err := ui.ConfirmSelect("Do you want to add another reviewer?")
			if err != nil {
				fmt.Printf("Error asking about more reviewers: %v\n", err)
				os.Exit(1)
			}

			// Stop asking if no more reviewers are to be added.
			if !addMore {
				break
			}
		}
	}

	// Confirm whether the user wants to reference issues.
	refIssues, err := ui.ConfirmSelect("Do you want to reference issues?")
	if err != nil {
		fmt.Printf("Error asking about issue references: %v\n", err)
		os.Exit(1)
	}

	// Collect issue references if confirmed.
	if refIssues {
		for {
			issue, err := ui.InputWithValidation(
				"Enter issue reference (e.g., '#123')",
				"#",
				func(input string) error {
					if !strings.HasPrefix(input, "#") {
						return errors.New("issue reference must start with #")
					}
					if len(input) < 2 {
						return errors.New("issue reference cannot be empty")
					}
					return nil
				},
			)
			if err != nil {
				fmt.Printf("Error entering issue reference: %v\n", err)
				os.Exit(1)
			}

			config.ReferenceIssues = append(config.ReferenceIssues, issue)

			addMore, err := ui.ConfirmSelect("Do you want to reference another issue?")
			if err != nil {
				fmt.Printf("Error asking about more issues: %v\n", err)
				os.Exit(1)
			}

			// Stop asking if no more issue references are required.
			if !addMore {
				break
			}
		}
	}

	// Format the final commit message using the provided configuration.
	commitMessage := commit.FormatCommitMessage(config)

	// Confirm the commit and proceed with execution.
	err = commit.ConfirmAndCommit(commitMessage)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Notify the user that the commit was created successfully.
	fmt.Println("✅ Commit successfully created")
}

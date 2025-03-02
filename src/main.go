package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	t "github.com/GiulianPoeta99/conventional_commits_cli/src/types"
	ui "github.com/GiulianPoeta99/conventional_commits_cli/src/ui"
)

func main() {
	fmt.Println("🚀 Conventional Commits Assistant")

	config := t.CommitConfig{}

	var err error

	config.Type, err = ui.SelectCommitType()
	if err != nil {
		fmt.Printf("Error selecting commit type: %v\n", err)
		os.Exit(1)
	}

	config.Scope, err = ui.OptionalInput("Add a scope for this change. (optional, press Enter to omit)")
	if err != nil {
		fmt.Printf("Error entering scope: %v\n", err)
		os.Exit(1)
	}

	useEmoji, err := ui.ConfirmSelect("Do you want to include an emoji?")
	if err != nil {
		fmt.Printf("Error selecting emoji option: %v\n", err)
		os.Exit(1)
	}

	if useEmoji {
		config.Emoji, err = ui.SelectEmojiWithSuggestions(config.Type)
		if err != nil {
			fmt.Printf("Error selecting emoji: %v\n", err)
			os.Exit(1)
		}
	}

	config.Description, err = ui.InputWithValidation(
		"Commit description",
		"",
		func(input string) error {
			if len(input) < 3 {
				return errors.New("Description must have at least 3 characters")
			}
			return nil
		},
	)
	if err != nil {
		fmt.Printf("Error entering description: %v\n", err)
		os.Exit(1)
	}

	config.Body, err = ui.OptionalInput("Commit body (optional, press Enter to omit)")
	if err != nil {
		fmt.Printf("Error entering body: %v\n", err)
		os.Exit(1)
	}

	config.Breaking, err = ui.ConfirmSelect("Is this a breaking change?")
	if err != nil {
		fmt.Printf("Error selecting breaking change: %v\n", err)
		os.Exit(1)
	}

	if config.Breaking {
		config.BreakingReason, err = ui.OptionalInput("Describe why this is a breaking change (optional, press Enter to use the default message)")
		if err != nil {
			fmt.Printf("Error entering breaking change reason: %v\n", err)
			os.Exit(1)
		}
	}

	addReviewers, err := ui.ConfirmSelect("Do you want to add reviewers?")
	if err != nil {
		fmt.Printf("Error asking about reviewers: %v\n", err)
		os.Exit(1)
	}

	if addReviewers {
		for {
			reviewer, err := ui.InputWithValidation(
				"Enter reviewer (e.g., 'John Smith')",
				"",
				func(input string) error {
					if len(input) < 1 {
						return errors.New("Reviewer name cannot be empty")
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

			if !addMore {
				break
			}
		}
	}

	refIssues, err := ui.ConfirmSelect("Do you want to reference issues?")
	if err != nil {
		fmt.Printf("Error asking about issue references: %v\n", err)
		os.Exit(1)
	}

	if refIssues {
		for {
			issue, err := ui.InputWithValidation(
				"Enter issue reference (e.g., '#123')",
				"#",
				func(input string) error {
					if !strings.HasPrefix(input, "#") {
						return errors.New("Issue reference must start with #")
					}
					if len(input) < 2 {
						return errors.New("Issue reference cannot be empty")
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

			if !addMore {
				break
			}
		}
	}

	commitMessage := formatCommitMessage(config)

	err = confirmAndCommit(commitMessage)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Commit successfully created")
}

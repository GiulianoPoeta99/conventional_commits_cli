// Package internal provides functions to format commit messages and execute commits based on user input.
package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	t "github.com/GiulianoPoeta99/conventional_commits_cli/internal/types"
	ui "github.com/GiulianoPoeta99/conventional_commits_cli/internal/ui"
)

// FormatCommitMessage formats the commit message according to the provided configuration.
// It constructs the message by combining type, scope, emoji, description, body, breaking changes,
// reviewers, and referenced issues.
func FormatCommitMessage(config t.CommitConfig) string {
	message := config.Type.Code

	// Append scope if provided.
	if config.Scope != "" {
		message += "(" + config.Scope + ")"
	}

	// Append an exclamation mark if it's a breaking change.
	if config.Breaking {
		message += "!"
	}

	message += ": "

	// Include the emoji code if available.
	if config.Emoji.Code != "" {
		message += ":" + config.Emoji.Code + ": "
	}

	message += config.Description

	// Append the commit body if provided.
	if config.Body != "" {
		message += "\n\n" + config.Body
	}

	// Append breaking change note and reason if applicable.
	if config.Breaking {
		if !strings.HasSuffix(message, "\n\n") {
			message += "\n\n"
		}

		message += "BREAKING CHANGE: "
		if config.BreakingReason != "" {
			message += config.BreakingReason
		} else {
			message += "This commit introduces changes incompatible with previous versions"
		}
	}

	// Append reviewers information.
	if len(config.Reviewers) > 0 {
		if !strings.HasSuffix(message, "\n\n") {
			message += "\n\n"
		}

		for _, reviewer := range config.Reviewers {
			message += "Reviewed-by: " + reviewer + "\n"
		}
		message = strings.TrimSuffix(message, "\n")
	}

	// Append referenced issues.
	if len(config.ReferenceIssues) > 0 {
		if !strings.HasSuffix(message, "\n\n") && !strings.HasSuffix(message, "\n") {
			message += "\n\n"
		} else if strings.HasSuffix(message, "\n") {
			message += "\n"
		} else {
			message += "\n\n"
		}

		for _, issue := range config.ReferenceIssues {
			message += "Refs: " + issue + "\n"
		}
		message = strings.TrimSuffix(message, "\n")
	}

	return message
}

// executeCommit executes the commit using git commands.
// First, it checks if there are staged changes and then commits with the provided message.
func executeCommit(message string) error {
	// Check for staged changes.
	cmd := exec.Command("git", "diff", "--staged", "--quiet")
	err := cmd.Run()

	if err != nil {
		// If there are staged changes, perform the commit.
		cmd = exec.Command("git", "commit", "-m", message)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		return errors.New("no staged changes to commit. Use 'git add' first")
	}
}

// ConfirmAndCommit prints the commit message for confirmation and then executes the commit if confirmed.
// It returns an error if the commit is cancelled or if an error occurs during the commit process.
func ConfirmAndCommit(message string) error {
	fmt.Println("\n============= Commit message =============")
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
	fmt.Println("==========================================")

	confirm, err := ui.ConfirmSelect("Confirm commit?")
	if err != nil {
		fmt.Printf("Error in the confirmation of the commit: %v\n", err)
		os.Exit(1)
	}

	if confirm {
		return executeCommit(message)
	}

	return errors.New("commit canceled by user")
}

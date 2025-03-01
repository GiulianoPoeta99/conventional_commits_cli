package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	t "conventional_commits_cli/src/types"
	ui "conventional_commits_cli/src/ui"
)

func formatCommitMessage(config t.CommitConfig) string {
	message := config.Type.Code

	if config.Scope != "" {
		message += "(" + config.Scope + ")"
	}

	if config.Breaking {
		message += "!"
	}

	message += ": "

	if config.Emoji.Code != "" {
		message += ":" + config.Emoji.Code + ": "
	}

	message += config.Description

	if config.Body != "" {
		message += "\n\n" + config.Body
	}

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

	if len(config.Reviewers) > 0 {
		if !strings.HasSuffix(message, "\n\n") {
			message += "\n\n"
		}

		for _, reviewer := range config.Reviewers {
			message += "Reviewed-by: " + reviewer + "\n"
		}
		message = strings.TrimSuffix(message, "\n")
	}

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

func executeCommit(message string) error {
	cmd := exec.Command("git", "diff", "--staged", "--quiet")
	err := cmd.Run()

	if err != nil {
		cmd = exec.Command("git", "commit", "-m", message)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		return errors.New("No staged changes to commit. Use 'git add' first")
	}
}

func confirmAndCommit(message string) error {
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

	return errors.New("Commit canceled by user")
}

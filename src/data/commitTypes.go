package data

import (
	t "github.com/GiulianPoeta99/conventional_commits_cli/src/types"
)

func GetCommitTypes() []t.CommitType {
	return []t.CommitType{
		{
			Code:        "feat",
			Description: "A new feature",
		},
		{
			Code:        "fix",
			Description: "A bug fix",
		},
		{
			Code:        "docs",
			Description: "Documentation only changes",
		},
		{
			Code:        "style",
			Description: `Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)`,
		},
		{
			Code:        "refactor",
			Description: "A code change that neither fixes a bug nor adds a feature",
		},
		{
			Code:        "perf",
			Description: "A code change that improves performance",
		},
		{
			Code:        "test",
			Description: "Adding missing tests or correcting existing tests",
		},
		{
			Code:        "build",
			Description: "Changes that affect the build system or external dependencies (examples scopes: gulp, broccoli, npm)",
		},
		{
			Code:        "ci",
			Description: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)",
		},
		{
			Code:        "chore",
			Description: "Other changes that don't modify src or test files",
		},
		{
			Code:        "revert",
			Description: "Reverts a previous commit",
		},
	}
}

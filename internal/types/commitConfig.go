// Package types defines the structures used for constructing commit configuration data.
package types

// CommitConfig holds all information required to format a commit message.
type CommitConfig struct {
	Type            CommitType
	Scope           string
	Emoji           Emoji
	Description     string
	Body            string
	Breaking        bool
	BreakingReason  string
	Reviewers       []string
	ReferenceIssues []string
}

package types

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

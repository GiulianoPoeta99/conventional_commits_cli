// Package ui provides user interface helpers using promptui for collecting inputs.
package ui

import (
	"fmt"
	"strings"

	d "github.com/GiulianoPoeta99/conventional_commits_cli/internal/data"
	t "github.com/GiulianoPoeta99/conventional_commits_cli/internal/types"

	"github.com/manifoldco/promptui"
)

// InputWithValidation displays a prompt with the given label and default value,
// validating the input using the provided function.
func InputWithValidation(
	label string,
	defaultValue string,
	validate func(input string) error,
) (string, error) {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   defaultValue,
		AllowEdit: true,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// SelectCommitType prompts the user to select a commit type from a list of available types.
// It returns the selected CommitType.
func SelectCommitType() (t.CommitType, error) {
	commitTypes := d.GetCommitTypes()
	items := []string{}

	// Format commit types into displayable strings.
	for _, t := range commitTypes {
		items = append(
			items,
			fmt.Sprintf("%s -> %s", strings.ToUpper(t.Code), t.Description),
		)
	}

	prompt := promptui.Select{
		Label: "Select the type of change that you're committing",
		Items: items,
		Size:  len(commitTypes),
	}

	index, _, err := prompt.Run()
	if err != nil {
		return t.CommitType{}, err
	}
	return commitTypes[index], nil
}

// SuggestEmojis returns a list of recommended emojis based on the provided commit type.
func SuggestEmojis(commitType t.CommitType) []t.Emoji {
	emojis := d.GetEmojis()
	suggestions := []t.Emoji{}

	// Map commit types to suggested emoji codes.
	typeToEmojis := map[string][]string{
		"feat":     {"sparkles", "rocket", "tada"},
		"fix":      {"bug", "ambulance", "adhesive_bandage", "goal_net"},
		"docs":     {"memo", "bulb", "pencil2"},
		"style":    {"art", "lipstick"},
		"refactor": {"recycle", "hammer", "truck"},
		"perf":     {"zap", "chart_with_upwards_trend"},
		"test":     {"white_check_mark", "test_tube"},
		"build":    {"package", "construction_worker"},
		"ci":       {"green_heart", "construction"},
		"chore":    {"wrench", "bricks"},
		"revert":   {"rewind", "coffin"},
	}

	// Filter emojis based on suggested codes.
	if emojiCodes, ok := typeToEmojis[commitType.Code]; ok {
		for _, code := range emojiCodes {
			for _, emoji := range emojis {
				if emoji.Code == code {
					suggestions = append(suggestions, emoji)
				}
			}
		}
	}

	return suggestions
}

// SelectEmojiWithSuggestions allows the user to select an emoji.
// It provides recommendations based on the commit type, placing them at the top.
func SelectEmojiWithSuggestions(commitType t.CommitType) (t.Emoji, error) {
	suggestions := SuggestEmojis(commitType)
	allEmojis := d.GetEmojis()

	// Merge the suggestions and the rest of the emojis.
	displayEmojis := append([]t.Emoji{}, suggestions...)

	for _, emoji := range allEmojis {
		found := false
		for _, suggested := range suggestions {
			if emoji.Code == suggested.Code {
				found = true
				break
			}
		}

		if !found {
			displayEmojis = append(displayEmojis, emoji)
		}
	}

	items := []string{}

	// Format the list of emojis for display; add a prefix for recommended ones.
	for i, e := range displayEmojis {
		prefix := ""
		if i < len(suggestions) {
			prefix = "🔍 "
		}
		items = append(
			items,
			fmt.Sprintf(
				"%s%s (:%s:) -> %s",
				prefix, e.Symbol, e.Code, e.Description,
			),
		)
	}

	// Create a prompt with search capability.
	prompt := promptui.Select{
		Label:        "Select an emoji (🔍 = Recommendation)",
		Items:        items,
		Size:         10,
		CursorPos:    0,
		HideSelected: false,
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(items[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return t.Emoji{}, err
	}
	return displayEmojis[index], nil
}

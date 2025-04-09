package ui

import (
	"fmt"
	"strings"

	d "github.com/GiulianPoeta99/conventional_commits_cli/internal/data"
	t "github.com/GiulianPoeta99/conventional_commits_cli/internal/types"

	"github.com/manifoldco/promptui"
)

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

func SelectCommitType() (t.CommitType, error) {
	commitTypes := d.GetCommitTypes()
	items := []string{}

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

func SuggestEmojis(commitType t.CommitType) []t.Emoji {
	emojis := d.GetEmojis()
	suggestions := []t.Emoji{}

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

func SelectEmojiWithSuggestions(commitType t.CommitType) (t.Emoji, error) {
	suggestions := SuggestEmojis(commitType)
	allEmojis := d.GetEmojis()

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

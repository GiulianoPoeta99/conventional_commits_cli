package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

// CommitType represents a conventional commit type with its description
type CommitType struct {
	Code        string
	Description string
}

// Emoji represents a commit emoji with its code and description
type Emoji struct {
	Code        string
	Description string
}

// CommitConfig holds the user's selections for a commit
type CommitConfig struct {
	Type        CommitType
	Scope       string
	Emoji       Emoji
	Description string
	Body        string
	Breaking    bool
}

// GetCommitTypes returns all available conventional commit types
func GetCommitTypes() []CommitType {
	return []CommitType{
		{"feat", "A new feature"},
		{"fix", "A bug fix"},
		{"docs", "Documentation only changes"},
		{"style", "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
		{"refactor", "A code change that neither fixes a bug nor adds a feature"},
		{"perf", "A code change that improves performance"},
		{"test", "Adding missing tests or correcting existing tests"},
		{"build", "Changes that affect the build system or external dependencies (examples scopes: gulp, broccoli, npm)"},
		{"ci", "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"},
		{"chore", "Other changes that don't modify src or test files"},
		{"revert", "Reverts a previous commit"},
	}
}

// GetEmojis returns all available emojis for commits
func GetEmojis() []Emoji {
	return []Emoji{
		{"art", "Improve structure / format of the code"},
		{"zap", "Improve performance"},
		{"fire", "Remove code or files"},
		{"bug", "Fix a bug"},
		{"ambulance", "Critical hotfix"},
		{"sparkles", "Introduce new features"},
		{"memo", "Add or update documentation"},
		{"rocket", "Deploy stuff"},
		{"lipstick", "Add or update the UI and style files"},
		{"tada", "Begin a project"},
		{"white_check_mark", "Add, update, or pass test"},
		{"lock", "Fix security issues"},
		{"closed_lock_with_key", "Add or update secrets"},
		{"bookmark", "Release / version tags"},
		{"rotating_light", "Fix compiler / linter warnings"},
		{"construction", "Work in progress"},
		{"green_heart", "Fix CI build"},
		{"arrow_down", "Downgrade dependencies"},
		{"arrow_up", "Upgrade dependencies"},
		{"pushpin", "Pin dependencies to specific versions"},
		{"construction_worker", "Add or update CI build system"},
		{"chart_with_upwards_trend", "Add or update analytics or track code"},
		{"recycle", "Refactor code"},
		{"heavy_plus_sign", "Add a dependency"},
		{"heavy_minus_sign", "Remove a dependency"},
		{"wrench", "Add or update configuration files"},
		{"hammer", "Add or update development scripts"},
		{"globe_with_meridians", "Internationalization and localization"},
		{"pencil2", "Fix typos"},
		{"poop", "Write bad code that needs to be improved"},
		{"rewind", "Revert changes"},
		{"twisted_rightwards_arrows", "Merge branches"},
		{"package", "Add or update compiled files or packages"},
		{"alien", "Update code due to external API changes"},
		{"truck", "Move or rename resources (e.g.: files, paths, routes)"},
		{"page_facing_up", "Add or update license"},
		{"boom", "Introduce breaking changes"},
		{"bento", "Add or update assets"},
		{"wheelchair", "Improve accessibility"},
		{"bulb", "Add or update comments in source code"},
		{"beers", "Write code drunkenly"},
		{"speech_balloon", "Add or update text and literals"},
		{"card_file_box", "Perform database related changes"},
		{"loud_sound", "Add or update logs"},
		{"mute", "Remove logs"},
		{"busts_in_silhouette", "Add or update contributor(s)"},
		{"children_crossing", "Improve user experience / usability"},
		{"building_construction", "Make architectural changes"},
		{"iphone", "Work on responsive design"},
		{"clown_face", "Mock things"},
		{"egg", "Add or update an easter egg"},
		{"see_no_evil", "Add or update a .gitignore file"},
		{"camera_flash", "Add or update snapshots"},
		{"alembic", "Perform experiments"},
		{"mag", "Improve SEO"},
		{"label", "Add or update types"},
		{"seedling", "Add or update seed files"},
		{"triangular_flag_on_post", "Add, update, or remove feature flags"},
		{"goal_net", "Catch errors"},
		{"dizzy", "Add or update animations and transitions"},
		{"wastebasket", "Deprecate code that needs to be cleaned up"},
		{"passport_control", "Work on code related to authorization, roles, and permissions"},
		{"adhesive_bandage", "Simple fix for a non-critical issue"},
		{"monocle_face", "Data exploration / inspection"},
		{"coffin", "Remove dead code"},
		{"test_tube", "Add a failing test"},
		{"necktie", "Add or update business logic"},
		{"stethoscope", "Add or update health check"},
		{"bricks", "Infrastructure related changes"},
		{"technologist", "Improve developer experience"},
		{"money_with_wings", "Add sponsorships or money related infrastructure"},
		{"thread", "Add or update code related to multithreading or concurrency"},
		{"safety_vest", "Add or update code related to validation"},
	}
}

// Function to select the commit type
func selectCommitType() (CommitType, error) {
	commitTypes := GetCommitTypes()
	items := []string{}

	for _, t := range commitTypes {
		items = append(items, fmt.Sprintf("%s -> %s", t.Code, t.Description))
	}

	prompt := promptui.Select{
		Label: "Select the type of change that you're committing",
		Items: items,
		Size:  len(commitTypes), // Show all available types
	}

	index, _, err := prompt.Run()
	if err != nil {
		return CommitType{}, err
	}
	return commitTypes[index], nil
}

// Function to input the commit scope
func inputScope() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Add a scope for this change (OPTIONAL) press Enter to omit",
		Default:   "",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// Function to select an emoji
func selectEmoji() (Emoji, error) {
	emojis := GetEmojis()
	items := []string{}

	for _, e := range emojis {
		items = append(items, fmt.Sprintf(":%s: -> %s", e.Code, e.Description))
	}

	prompt := promptui.Select{
		Label:        "Select an emoji",
		Items:        items,
		Size:         10,
		CursorPos:    0,
		HideSelected: false,
		// Add a searcher function to filter emojis
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(items[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return Emoji{}, err
	}
	return emojis[index], nil
}

// Function to input the commit description
func inputDescription() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Commit description",
		Default:   "",
		AllowEdit: true,
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Description must have at least 3 characters")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// Function to input the commit body
func inputBody() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Commit body (optional, press Enter to omit)",
		Default:   "",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// Function to ask if it's a breaking change
func askBreakingChange() (bool, error) {
	prompt := promptui.Select{
		Label: "Is this a breaking change?",
		Items: []string{"No", "Yes"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return index == 1, nil
}

// Function to format the commit message
func formatCommitMessage(config CommitConfig) string {
	// Initialize the commit message with the type
	message := config.Type.Code

	// Add scope if provided
	if config.Scope != "" {
		message += "(" + config.Scope + ")"
	}

	// Add breaking change indicator if needed
	if config.Breaking {
		message += "!"
	}

	// Add description
	message += ": "

	// Add emoji if selected
	if config.Emoji.Code != "" {
		message += ":" + config.Emoji.Code + ": "
	}

	message += config.Description

	// Add body if provided
	if config.Body != "" {
		message += "\n\n" + config.Body
	}

	// Add BREAKING CHANGE footer if needed
	if config.Breaking {
		if config.Body != "" {
			message += "\n"
		} else {
			message += "\n\n"
		}
		message += "BREAKING CHANGE: This commit introduces changes incompatible with previous versions"
	}

	return message
}

// Function to execute the commit
func executeCommit(message string) error {
	// First check if there are staged changes
	cmd := exec.Command("git", "diff", "--staged", "--quiet")
	err := cmd.Run()

	if err != nil {
		// If there are staged changes (the previous command returns an error)
		cmd = exec.Command("git", "commit", "-m", message)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		return errors.New("No staged changes to commit. Use 'git add' first")
	}
}

// Function to show the commit message and confirm
func confirmAndCommit(message string) error {
	fmt.Println("\n=== Commit message ===")
	fmt.Println(message)
	fmt.Println("========================\n")

	prompt := promptui.Select{
		Label: "Confirm commit?",
		Items: []string{"Yes", "No"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return err
	}

	if index == 0 {
		return executeCommit(message)
	}

	return errors.New("Commit canceled by user")
}

// Function to generate emoji suggestions based on commit type
func suggestEmojis(commitType CommitType) []Emoji {
	emojis := GetEmojis()
	suggestions := []Emoji{}

	// Mapping commit types to relevant emoji codes
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

// Function to select emoji with suggestions
func selectEmojiWithSuggestions(commitType CommitType) (Emoji, error) {
	suggestions := suggestEmojis(commitType)
	allEmojis := GetEmojis()

	// Create a slice with suggestions first
	displayEmojis := append([]Emoji{}, suggestions...)

	// Add the rest of emojis
	for _, emoji := range allEmojis {
		// Check if it's already in the suggestions
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

	// Create items to display
	items := []string{}

	for i, e := range displayEmojis {
		prefix := ""
		if i < len(suggestions) {
			prefix = "🔍 "
		}
		items = append(items, fmt.Sprintf("%s:%s: -> %s", prefix, e.Code, e.Description))
	}

	prompt := promptui.Select{
		Label:        "Select an emoji",
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
		return Emoji{}, err
	}
	return displayEmojis[index], nil
}

func main() {
	fmt.Println("🚀 Conventional Commits Assistant")

	// Collect commit information
	config := CommitConfig{}

	var err error

	// Select commit type
	config.Type, err = selectCommitType()
	if err != nil {
		fmt.Printf("Error selecting commit type: %v\n", err)
		os.Exit(1)
	}

	// Ask for scope (optional)
	config.Scope, err = inputScope()
	if err != nil {
		fmt.Printf("Error entering scope: %v\n", err)
		os.Exit(1)
	}

	// Ask if it's a breaking change
	config.Breaking, err = askBreakingChange()
	if err != nil {
		fmt.Printf("Error selecting breaking change: %v\n", err)
		os.Exit(1)
	}

	// Select emoji (optional)
	emojiPrompt := promptui.Select{
		Label: "Do you want to include an emoji?",
		Items: []string{"Yes", "No"},
	}

	emojiIndex, _, err := emojiPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting emoji option: %v\n", err)
		os.Exit(1)
	}

	if emojiIndex == 0 {
		// Use the improved function that shows suggestions based on commit type
		config.Emoji, err = selectEmojiWithSuggestions(config.Type)
		if err != nil {
			fmt.Printf("Error selecting emoji: %v\n", err)
			os.Exit(1)
		}
	}

	// Enter description
	config.Description, err = inputDescription()
	if err != nil {
		fmt.Printf("Error entering description: %v\n", err)
		os.Exit(1)
	}

	// Enter body (optional)
	config.Body, err = inputBody()
	if err != nil {
		fmt.Printf("Error entering body: %v\n", err)
		os.Exit(1)
	}

	// Format commit message
	commitMessage := formatCommitMessage(config)

	// Confirm and execute commit
	err = confirmAndCommit(commitMessage)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Commit successfully created")
}

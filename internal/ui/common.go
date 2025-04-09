// Package ui provides user interface helpers using promptui for collecting inputs.
package ui

import "github.com/manifoldco/promptui"

// ConfirmSelect displays a selection prompt asking for a confirmation (Yes/No).
// It returns true if "Yes" is selected.
func ConfirmSelect(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"No", "Yes"},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return index == 1, nil
}

// OptionalInput displays a prompt that allows the user to input an optional value.
// It returns the entered value or an empty string if omitted.
func OptionalInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   "",
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

package ui

import "github.com/manifoldco/promptui"

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

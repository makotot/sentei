package tui

import (
	"github.com/charmbracelet/huh"
)

func Form(branches []string) ([]string, error) {
	selected := []string{}

	var options []huh.Option[string]

	for _, branch := range branches {
		options = append(options, huh.NewOption(branch, branch))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select branches want to delete").
				Options(options...).
				Value(&selected),
		),
	)

	err := form.Run()

	return selected, err
}

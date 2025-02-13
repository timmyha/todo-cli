package ui

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

func SelectTask(tasks []string) string {
	var selectedTask string

	options := []huh.Option[string]{}
	for _, task := range tasks {
		options = append(options, huh.NewOption(task, task))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a task to complete").
				Options(options...).
				Value(&selectedTask),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Selection error:", err)
		return ""
	}

	return selectedTask
}


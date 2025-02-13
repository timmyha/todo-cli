package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func SelectTask(tasks []string) string {
	var selectedTask string

	form := huh.NewForm(
		huh.NewSelect[string]().
			Title("Select a task to complete").
			Options(huh.NewOptionsFromSlice(tasks)...).
			Value(&selectedTask),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Selection error:", err)
		return ""
	}

	return selectedTask
}


package cmd

import (
	"fmt"
	"github.com/timmyha/todo-cli/internal/storage"
	"github.com/timmyha/todo-cli/internal/ui"
)

func CompleteTask(day string) {
	tasks, err := storage.GetTasks(day)
	if err != nil || len(tasks) == 0 {
		fmt.Println("No tasks found for", day)
		return
	}

	selectedTask := ui.SelectTask(tasks)
	if selectedTask == "" {
		fmt.Println("No task selected.")
		return
	}

	err = storage.MarkTaskCompleted(day, selectedTask)
	if err != nil {
		fmt.Println("Error completing task:", err)
	} else {
		fmt.Println("Task marked as completed.")
	}
}


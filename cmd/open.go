package cmd

import (
	"os"
	"os/exec"
)

func OpenEditor(arg string) {
	switch arg {
	case "completed":
		cmd := exec.Command("nvim", "completed.md")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	default:
		cmd := exec.Command("nvim", "todo.md")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
}

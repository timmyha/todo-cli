package cmd

import (
	"os"
	"os/exec"
)

func OpenEditor() {
	cmd := exec.Command("nvim", "todo.md")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}


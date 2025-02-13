package storage

import (
	"bufio"
	"os"
	"strings"
	"time"
)

const (
	todoFile      = "todo.md"
	completedFile = "completed.md"
)

func AppendTask(day, task string) error {
	file, err := os.OpenFile(todoFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	sectionFound := false

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if strings.EqualFold(line, "# "+day) {
			sectionFound = true
		}
	}

	if !sectionFound {
		lines = append(lines, "\n# "+day)
	}

	lines = append(lines, "- [ ] "+task)

	return os.WriteFile(todoFile, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func GetTasks(day string) ([]string, error) {
	file, err := os.Open(todoFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	inSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.EqualFold(line, "# "+day) {
			inSection = true
			continue
		}
		if inSection && strings.HasPrefix(line, "# ") {
			break
		}
		if inSection && strings.HasPrefix(line, "- [ ] ") {
			tasks = append(tasks, line)
		}
	}
	return tasks, nil
}

func MarkTaskCompleted(day, task string) error {
	file, err := os.Open(todoFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var newLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != task {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile(todoFile, []byte(strings.Join(newLines, "\n")+"\n"), 0644)
	if err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02")
	completedTask := task + " (Completed on " + now + ")"

	file, err = os.OpenFile(completedFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(completedTask + "\n")
	return err
}

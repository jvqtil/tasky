package main

import (
	"fmt"
	"os"
	"encoding/json"
	"strings"
	"regexp"
)

const tasksFile = "tasky.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("cli tasks manager")
		return
	}

	do := os.Args[1]

	switch do {
		case "add", "put", "touch", "new": 
		task := strings.Join(os.Args[2:], " ")
		addTask(task)
		case "list", "ls":
		listTasks()
		case "rem", "rm", "remove", "del", "delete": 
		task := strings.Join(os.Args[2:], " ")
		remTask(task)
	}
}

func addTask(task string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error loading tasks")
	}

	if task == "" {
		return
	}

	tasks = append(tasks, task)

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("error saving tasks")
		return
	}

	fmt.Println("task added")
}

func remTask(task string) {
	task = strings.ReplaceAll(task, "*", ".*")
	task = strings.ReplaceAll(task, "?", ".")

	reg, err := regexp.Compile("^" + task + "$")
	if err != nil {
		fmt.Println("error", err)
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error loading tasks")
	}

	var updTasks []string
	for _, t := range tasks {
		if !reg.MatchString(t) {
			updTasks = append(updTasks, t)
		}
	}

	if len(updTasks) == len(tasks) {
		fmt.Println("No such task found")
		return
	}

	err = saveTasks(updTasks)
	if err != nil {
		fmt.Println("error saving tasks")
		return
	}

	fmt.Println("task removed")
}

func loadTasks() ([]string, error) {
	var tasks []string

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error loading tasks")
		return
	}

	if len(tasks) == 0 {
		fmt.Println("no tasks")
		return
	}

	fmt.Println("tasks list")
	fmt.Println()
	for _, task := range tasks {
		fmt.Println(task)
	}
}

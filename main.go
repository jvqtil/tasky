package main

import (
	"fmt"
	"os"
	"encoding/json"
	"strings"
	"regexp"
)

const taskyFile = "tasky.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("tasky is a cli task manager!")
		return
	}

	help := "commands: \n add \n remove \n list \nsee aliases to this commands on https://github.com/jvqtil/tasky/"
	do := os.Args[1]

	switch do {
		case "add", "put", "touch", "new": 
		task := strings.Join(os.Args[2:], " ")
		addTask(task)
		case "list", "ls":
		listTasks()
		case "rem", "rm", "remove", "delete", "del": 
		task := strings.Join(os.Args[2:], " ")
		remTask(task)
		case "help", "man":
		fmt.Println(help)
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
		fmt.Println("no such task found")
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

	data, err := os.ReadFile(taskyFile)
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
	return os.WriteFile(taskyFile, data, 0644)
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

	fmt.Println("tasks list -", taskyFile)
	fmt.Println()
	for _, task := range tasks {
		fmt.Println(task)
	}
}

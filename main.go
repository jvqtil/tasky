package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const taskyFile = "tasky.json"
const archiveFile = "taskyarchive.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("tasky is a cli task manager!\n")
		listTasks()
		return
	}

	help := "commands: \n add \n remove \n list \n archive \nsee aliases to this commands on https://github.com/jvqtil/tasky/"
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
	case "archive", `hide`, `move`, `mv`:
		if len(os.Args) > 2 {
			task := strings.Join(os.Args[2:], " ")
			archiveTask(task)
		} else {
			listArchive()
		}
	case "help", "man":
		fmt.Println(help)
	}
}

func addTask(task string) {
	tasks, err := loadTasks(taskyFile)
	if err != nil {
		fmt.Println("error loading tasks")
		return
	}

	if task == "" {
		return
	}

	tasks = append(tasks, task)

	err = saveTasks(taskyFile, tasks)
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
		return
	}

	tasks, err := loadTasks(taskyFile)
	if err != nil {
		fmt.Println("error loading tasks")
		return
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

	err = saveTasks(taskyFile, updTasks)
	if err != nil {
		fmt.Println("error saving tasks")
		return
	}

	fmt.Println("task removed")
}

func archiveTask(task string) {
	task = strings.ReplaceAll(task, "*", ".*")
	task = strings.ReplaceAll(task, "?", ".")

	reg, err := regexp.Compile("^" + task + "$")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	tasks, err := loadTasks(taskyFile)
	if err != nil {
		fmt.Println("error loading tasks")
		return
	}

	archive, err := loadTasks(archiveFile)
	if err != nil {
		fmt.Println("error loading archive")
		return
	}

	var updTasks []string
	var foundInTasks bool
	for _, t := range tasks {
		if reg.MatchString(t) {
			foundInTasks = true
			archive = append(archive, t)
		} else {
			updTasks = append(updTasks, t)
		}
	}

	if foundInTasks {
		err = saveTasks(taskyFile, updTasks)
		if err != nil {
			fmt.Println("error saving tasks")
			return
		}

		err = saveTasks(archiveFile, archive)
		if err != nil {
			fmt.Println("error saving archive")
			return
		}

		fmt.Println("task archived")
		return
	}

	var updArchive []string
	var foundInArchive bool
	for _, t := range archive {
		if reg.MatchString(t) {
			foundInArchive = true
			tasks = append(tasks, t)
		} else {
			updArchive = append(updArchive, t)
		}
	}

	if foundInArchive {
		err = saveTasks(taskyFile, tasks)
		if err != nil {
			fmt.Println("error saving tasks")
			return
		}

		err = saveTasks(archiveFile, updArchive)
		if err != nil {
			fmt.Println("error saving archive")
			return
		}

		fmt.Println("task unarchived")
		return
	}

	fmt.Println("no such task found")
}

func listTasks() {
	tasks, err := loadTasks(taskyFile)
	if err != nil {
		fmt.Println("error loading tasks")
		return
	}

	count := len(tasks)
	if count == 0 {
		fmt.Println("no tasks")
		return
	}

	fmt.Printf("tasks list (%d) - %s\n\n", count, taskyFile)
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func listArchive() {
	archive, err := loadTasks(archiveFile)
	if err != nil {
		fmt.Println("error loading archive")
		return
	}

	count := len(archive)
	if count == 0 {
		fmt.Println("no archived tasks")
		return
	}

	fmt.Printf("archived tasks list (%d) - %s\n\n", count, archiveFile)
	for _, task := range archive {
		fmt.Println(task)
	}
}

func loadTasks(fileName string) ([]string, error) {
	var tasks []string

	data, err := os.ReadFile(fileName)
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

func saveTasks(fileName string, tasks []string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

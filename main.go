package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const taskyFile = "tasky.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("tasky is a cli task manager!\n")
		err := listTasks()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	help := "commands: \n add \n list \n done \nsee aliases to this commands on https://github.com/jvqtil/tasky"
	do := os.Args[1]

	switch strings.ToLower(do) {
	case "add", "put", "touch", "new", "make":
		task := strings.Join(os.Args[2:], " ")
		err := addTask(task)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "list", "ls":
		err := listTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "done", "did":
		task := strings.Join(os.Args[2:], " ")
		err := doneTask(task)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "help", "man":
		fmt.Println(help)
	default:
		fmt.Println("command not found, see tasky help\n", help)
	}
}

func addTask(thisTask string) error {
	if thisTask == "" {
		return fmt.Errorf("cant save empty task!")
	}

	tasks, err := loadTasks(taskyFile)
	if err != nil {
		return err
	}

	tasks = append(tasks, thisTask)

	err = saveTasks(taskyFile, tasks)
	if err != nil {
		return err
	}

	fmt.Println("task added!")
	return nil
}

func listTasks() error {
	tasks, err := loadTasks(taskyFile)
	if err != nil {
		return err
	}

	count := len(tasks)
	if count == 0 {
		return fmt.Errorf("no tasks found!")
	}

	fmt.Printf("tasks list (%d) - %s\n\n", count, taskyFile)
	for _, task := range tasks {
		fmt.Println(task)
	}

	return nil
}

func doneTask(task string) error {
	if task == "all" {
		task = "*"
	}

	task = strings.ReplaceAll(task, "*", ".*")
	task = strings.ReplaceAll(task, "?", ".")

	reg, err := regexp.Compile("^" + task + "$")
	if err != nil {
		return err
	}

	tasks, err := loadTasks(taskyFile)
	if err != nil {
		return err
	}

	var updTasks []string
	for _, t := range tasks {
		if !reg.MatchString(t) {
			updTasks = append(updTasks, t)
		}
	}

	if len(updTasks) == len(tasks) {
		return fmt.Errorf("no such task found!")
	}

	err = saveTasks(taskyFile, updTasks)
	if err != nil {
		return err
	}

	fmt.Println("task done!")
	return nil
}

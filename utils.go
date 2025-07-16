package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTasks(fileName string) ([]string, error) {
	var tasks []string

	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, fmt.Errorf("error reading tasky file! %v", err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding tasky file! %v", err)
	}

	return tasks, nil
}

func saveTasks(fileName string, tasks []string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding tasky file! %v", err)
	}

	return os.WriteFile(fileName, data, 0644)
}

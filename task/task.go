package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type TaskId = int

type Task struct {
	Description string `json:"description"`
	ID          TaskId `json:"id"`
	Done        bool   `json:"done"`
}

func LoadTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(file) == 0 {
		fmt.Println("File is empty. Returning empty task list.")
		return []Task{}, nil
	}

	var tasks []Task
	if err = json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(filename string, tasks []Task) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func AddTask(description string, filename string) (TaskId, error) {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return 0, err
	}
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := Task{Description: description, ID: id, Done: false}
	tasks = append(tasks, task)

	err = SaveTasks(filename, tasks)
	return task.ID, err
}

func UpdateTaskStatus(filename string, id int) error {
	tasks, err := LoadTasks(filename)

	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %v not found", id)
	}

	return SaveTasks(filename, tasks)
}

func DeleteTask(filename string, id int) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	updatedTasks := make([]Task, 0, len(tasks))
	found := false
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("task not found")
	}

	return SaveTasks(filename, updatedTasks)
}

func ReadTasks(filename string) []Task {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return []Task{}
	}
	return tasks
}
